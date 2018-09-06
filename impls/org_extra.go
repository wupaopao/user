package impls

import (
	"time"

	"basis/sms/proxy/sms"
	"business/agency/proxy/agency"
	"business/user/cidl"
	"business/user/common/cache"
	"business/user/common/com"
	"business/user/common/db"
	"common/api"

	cache2 "github.com/mz-eco/mz/cache"
	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
	"github.com/mz-eco/mz/utils"
)

func init() {
	AddUserOrgUserLoginHandler()
	AddOrgUserLoginMobileVerifyHandler()

	AddUserOrgUserLogoutHandler()
}

type UserOrgUserLoginImpl struct {
	cidl.ApiUserOrgUserLogin
}

func AddUserOrgUserLoginHandler() {
	AddHandler(
		cidl.META_USER_ORG_USER_LOGIN,
		func() http.ApiHandler {
			return &UserOrgUserLoginImpl{
				ApiUserOrgUserLogin: cidl.MakeApiUserOrgUserLogin(),
			}
		},
	)
}

func (m *UserOrgUserLoginImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	mobile := m.Ask.Mobile
	userCache := cache.NewUserCache()
	cacheKey := userCache.KeyMobileVerifyCityLogin(mobile)
	value, err := userCache.Cache.Get(cacheKey)
	if err != nil && err != cache2.Nil {
		ctx.Errorf(api.ErrCacheReadFailed, "get mobile verify code cache failed. %s", err)
		return
	}

	// TODO 记得删除假验证码
	if "520398" != m.Ask.VerifyCode && (err == cache2.Nil || value != m.Ask.VerifyCode) {
		ctx.Errorf(api.ErrSMSVerifyCodeNotMatch, "verify code does not match.")
		return
	}

	// 设置授权信息
	dbUser := db.NewMallUser()
	user, err := dbUser.GetUserByMobile(mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return
	}

	// 用户不存在
	if err == conn.ErrNoRows || (!user.IsOrgStaff && !user.IsOrgManager) {
		ctx.Errorf(cidl.ErrOrgLoginUserNotExist, "user was not exist")
		return
	}

	// 获取组织成员信息
	ackStaff, err := agency.NewProxy("agency-service").InnerAgencyStaffInfoByUserID(user.UserID)
	if err != nil {
		ctx.Errorf(api.ErrProxyFailed, "get agency staff info failed. %s", err)
		return
	}

	//该账号被禁用
	if ackStaff.Staff.IsDisable {
		ctx.Errorf(cidl.ErrLoginUserForbidden, "user was forbidden")
		return
	}

	auth := &com.AuthCity{
		UserId: user.UserID,
		User:   user,
		Staff: &cidl.AuthCityStaff{
			OrganizationId:    ackStaff.Staff.OrganizationId,
			OrganizationName:  ackStaff.Staff.OrganizationName,
			GroupBuyingMode:   uint32(ackStaff.Organization.GroupBuyingMode),
			RoleId:            ackStaff.Staff.RoleId,
			RoleName:          ackStaff.Staff.RoleName,
			RoleAuthorization: ackStaff.StaffRole.RoleAuthorization,
		},
	}

	token, err := auth.NewToken()
	if err != nil {
		ctx.Errorf(cidl.ErrOrgGenerateTokenFailed, "new auth token failed. %s", err)
		return
	}

	m.Ack.User = user
	m.Ack.UserId = user.UserID
	m.Ack.Token = token
	m.Ack.Staff = auth.Staff

	ctx.Json(m.Ack)

	// 删除手机验证码
	_, err = userCache.Cache.Delete(cacheKey)
	if err != nil {
		log.Warnf("delete verify code cache failed. %s", err)
		return
	}

}

type OrgUserLoginMobileVerifyImpl struct {
	cidl.ApiOrgUserLoginMobileVerify
}

func AddOrgUserLoginMobileVerifyHandler() {
	AddHandler(
		cidl.META_ORG_USER_LOGIN_MOBILE_VERIFY,
		func() http.ApiHandler {
			return &OrgUserLoginMobileVerifyImpl{
				ApiOrgUserLoginMobileVerify: cidl.MakeApiOrgUserLoginMobileVerify(),
			}
		},
	)
}

func (m *OrgUserLoginMobileVerifyImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	mobile := m.Query.Mobile
	userCache := cache.NewUserCache()
	success, err := userCache.Cache.SetNX(userCache.KeyMobileVerifyCityLoginInterval(mobile), 1, 55*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set city login verify code interval failed. %s", err)
		return
	}

	if !success {
		ctx.Errorf(api.ErrSendSMSIntervalLimit, "time of sending verify code has not arrived.")
		return
	}

	verifyCode := utils.VerifyCode6()
	err = userCache.Cache.Set(userCache.KeyMobileVerifyCityLogin(mobile), verifyCode, 10*time.Minute+5*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set city login verify code failed. %s", err)
		return
	}

	_, err = sms.NewProxy("sms-service").SmsSendTemplateActionVerify(&sms.AskSmsSendTemplateActionVerify{
		Action: "城市合伙人端登陆",
		Mobile: m.Query.Mobile,
		Code:   verifyCode,
		Time:   "10分钟",
	})
	if err != nil {
		ctx.Errorf(api.ErrSendSMSFailed, "send sms verify code failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 退出登陆
type UserOrgUserLogoutImpl struct {
	cidl.ApiUserOrgUserLogout
}

func AddUserOrgUserLogoutHandler() {
	AddHandler(
		cidl.META_USER_ORG_USER_LOGOUT,
		func() http.ApiHandler {
			return &UserOrgUserLogoutImpl{
				ApiUserOrgUserLogout: cidl.MakeApiUserOrgUserLogout(),
			}
		},
	)
}

func (m *UserOrgUserLogoutImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	token := ctx.Session.Token
	err = com.DeleteAuthCityByToken(token)
	if err != nil {
		ctx.Errorf(api.ErrDeleteTokenFailed, "delete auth city token failed. %s", err)
		return
	}

	ctx.Succeed()
}

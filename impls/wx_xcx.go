package impls

import (
	"time"

	"basis/sms/proxy/sms"
	mq2 "business/auth/common/mq"
	"business/user/cidl"
	"business/user/common/cache"
	"business/user/common/com"
	"business/user/common/db"
	"business/user/common/mq"
	"business/user/common/wx"
	"common/api"

	cache2 "github.com/mz-eco/mz/cache"
	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
	"github.com/mz-eco/mz/utils"
)

func init() {
	AddWxXcxUserAuthHandler()
	AddWxXcxUserAuthFakeHandler()
	AddWxXcxUserFakeUserListHandler()
	AddWxXcxUserEditBasicByUserIDHandler()
	AddWxXcxUserEditNicknameByUserIDHandler()
	AddWxXcxUserMobileVerifyHandler()
	AddWxXcxUserEditMobileByUserIDHandler()
	AddUserWxXcxUserWxBindMobileVerifyHandler()
	AddUserWxXcxUserWxBindMobileHandler()
	AddWxXcxUserCheckBindMobileByMobileHandler()
	AddWxXcxUserUnbindWxHandler()
}

// 微信小程序授权登陆
type WxXcxUserAuthImpl struct {
	cidl.ApiWxXcxUserAuth
}

func AddWxXcxUserAuthHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_AUTH,
		func() http.ApiHandler {
			return &WxXcxUserAuthImpl{
				ApiWxXcxUserAuth: cidl.MakeApiWxXcxUserAuth(),
			}
		},
	)
}

func (m *WxXcxUserAuthImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	wxXcx := wx.GetWxXcx()
	sessionKey, err := wxXcx.GetSessionKey(m.Ask.Code)
	if err != nil {
		ctx.Errorf(api.ErrWxXcxGetSessionKeyFailed, "get session key failed. %s", err)
		return
	}

	userInfo, err := wxXcx.DecryptUserInfoEncryptedData(sessionKey.SessionKey, m.Ask.EncryptedData, m.Ask.InitVector)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxDecryptUserInfoFailed, "decrypt user info failed. %s", err)
		return
	}

	// 未关注公众号进入小程序时，微信不返回unionId, 但是UnoinId在UserInfo加密信息里
	if sessionKey.UnionId == "" {
		sessionKey.UnionId = userInfo.UnionId
	}

	unionId := sessionKey.UnionId
	if unionId == "" {
		ctx.Errorf(cidl.ErrWxXcxEmptyUnionId, "get empty union id.", err)
		return
	}

	dbUser := db.NewMallUser()
	isVisitor := false
	user, err := dbUser.GetUserByUnionId(unionId)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user failed. %s", err)
		return

	}else if err == conn.ErrNoRows {
		user = nil
		isVisitor = true
	}

	if !isVisitor {
		if (user.IsOrgManager && user.IsDisableOrgManager) || (user.IsOrgStaff && user.IsDisableOrgStaff) || (user.IsCmtManager && user.IsDisableCmtManger) {
			ctx.Errorf(cidl.ErrLoginUserForbidden,"account is forbidden.")
			return
		}
	
	}

	auth := &com.AuthWxXcx{
		IsVisitor:  isVisitor,
		SessionKey: sessionKey,
		User:       user,
		WxUserInfo: userInfo,
	}

	if isVisitor == true {
		auth.UserId = sessionKey.UnionId

	} else {
		auth.UserId = user.UserID
		err = auth.SetUserInfo(user)
		if err != nil {
			ctx.Errorf(cidl.ErrGenerateAuthWxXcxTokenFailed, "set auth user info failed. %s", err)
			return
		}
	}

	token, err := auth.NewToken()
	if err != nil {
		ctx.Errorf(cidl.ErrGenerateAuthWxXcxTokenFailed, "new wx xcx auth token failed. %s", err)
		return
	}

	m.Ack.IsVisitor = isVisitor
	m.Ack.Token = token
	m.Ack.UserId = auth.UserId
	m.Ack.User = user

	m.Ack.Organization = auth.Organization
	m.Ack.CommunityManager = auth.CommunityManager

	ctx.Json(m.Ack)
}

// TODO 微信假授权接口，需要删除
type WxXcxUserAuthFakeImpl struct {
	cidl.ApiWxXcxUserAuthFake
}

func AddWxXcxUserAuthFakeHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_AUTH_FAKE,
		func() http.ApiHandler {
			return &WxXcxUserAuthFakeImpl{
				ApiWxXcxUserAuthFake: cidl.MakeApiWxXcxUserAuthFake(),
			}
		},
	)
}

func (m *WxXcxUserAuthFakeImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	wxXcx := wx.GetWxXcx()
	sessionKey, err := wxXcx.GetSessionKey(m.Ask.Code)
	if err != nil {
		ctx.Errorf(api.ErrWxXcxGetSessionKeyFailed, "get session key failed. %s", err)
		return
	}

	userInfo, err := wxXcx.DecryptUserInfoEncryptedData(sessionKey.SessionKey, m.Ask.EncryptedData, m.Ask.InitVector)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxDecryptUserInfoFailed, "decrypt user info failed. %s", err)
		return
	}

	// 未关注公众号进入小程序时，微信不返回unionId, 但是UnoinId在UserInfo加密信息里
	if sessionKey.UnionId == "" {
		sessionKey.UnionId = userInfo.UnionId
	}

	// TODO 假union id
	sessionKey.UnionId = m.Ask.FakeUnionId
	userInfo.UnionId = m.Ask.FakeUnionId
	userInfo.Nickname = "假" + userInfo.Nickname

	unionId := sessionKey.UnionId
	if unionId == "" {
		ctx.Errorf(cidl.ErrWxXcxEmptyUnionId, "get empty union id.", err)
		return
	}

	dbUser := db.NewMallUser()
	isVisitor := false
	user, err := dbUser.GetUserByUnionId(unionId)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user failed. %s", err)
		return

	} else if err == conn.ErrNoRows {
		user = nil
		isVisitor = true
	}

	auth := &com.AuthWxXcx{
		IsVisitor:  isVisitor,
		SessionKey: sessionKey,
		User:       user,
		WxUserInfo: userInfo,
	}

	if isVisitor == true {
		auth.UserId = sessionKey.UnionId
	} else {
		auth.UserId = user.UserID
		err = auth.SetUserInfo(user)
		if err != nil {
			ctx.Errorf(cidl.ErrGenerateAuthWxXcxTokenFailed, "set auth user info failed. %s", err)
			return
		}
	}

	token, err := auth.NewToken()
	if err != nil {
		ctx.Errorf(cidl.ErrGenerateAuthWxXcxTokenFailed, "new wx xcx auth token failed. %s", err)
		return
	}

	m.Ack.IsVisitor = isVisitor
	m.Ack.Token = token
	m.Ack.UserId = auth.UserId
	m.Ack.User = user

	m.Ack.Organization = auth.Organization
	m.Ack.CommunityManager = auth.CommunityManager

	ctx.Json(m.Ack)
}

// TODO 微信假用户，待删除
type WxXcxUserFakeUserListImpl struct {
	cidl.ApiWxXcxUserFakeUserList
}

func AddWxXcxUserFakeUserListHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_FAKE_USER_LIST,
		func() http.ApiHandler {
			return &WxXcxUserFakeUserListImpl{
				ApiWxXcxUserFakeUserList: cidl.MakeApiWxXcxUserFakeUserList(),
			}
		},
	)
}

func (m *WxXcxUserFakeUserListImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	strSql := `
		SELECT * FROM usr_fake_user
	`
	dbUser := db.NewMallUser()
	rows, err := dbUser.DB.Query(strSql)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get fake user failed. %s", err)
		return
	}

	for rows.Next() {
		fakeUser := cidl.NewFakeUser()
		err = rows.StructScan(fakeUser)
		if err != nil {
			ctx.Errorf(api.ErrDbQueryFailed, "get fake user failed. %s", err)
			return
		}
		m.Ack.Users = append(m.Ack.Users, fakeUser)
	}

	ctx.Json(m.Ack)
}

// 修改基本信息
type WxXcxUserEditBasicByUserIDImpl struct {
	cidl.ApiWxXcxUserEditBasicByUserID
}

func AddWxXcxUserEditBasicByUserIDHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_EDIT_BASIC_BY_USER_ID,
		func() http.ApiHandler {
			return &WxXcxUserEditBasicByUserIDImpl{
				ApiWxXcxUserEditBasicByUserID: cidl.MakeApiWxXcxUserEditBasicByUserID(),
			}
		},
	)
}

func (m *WxXcxUserEditBasicByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	uid := ctx.Session.Uid
	dbUser := db.NewMallUser()
	_, err = dbUser.UpdateUserBasic(uid, m.Ask.Nickname, m.Ask.Mobile)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update user basic failed. %s", err)
		return
	}

	// 广播更改消息
	topic, err := mq.GetTopicServiceUserService()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-user-service failed. %s", err)
		return
	}

	err = topic.ModifyUserInfo(&mq.ModifyUserInfoMessage{
		UserId: uid,
		Values: map[string]interface{}{
			"nickname": m.Ask.Nickname,
			"mobile":   m.Ask.Mobile,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish service-user-service failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 修改昵称
type WxXcxUserEditNicknameByUserIDImpl struct {
	cidl.ApiWxXcxUserEditNicknameByUserID
}

func AddWxXcxUserEditNicknameByUserIDHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_EDIT_NICKNAME_BY_USER_ID,
		func() http.ApiHandler {
			return &WxXcxUserEditNicknameByUserIDImpl{
				ApiWxXcxUserEditNicknameByUserID: cidl.MakeApiWxXcxUserEditNicknameByUserID(),
			}
		},
	)
}

func (m *WxXcxUserEditNicknameByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	uid := ctx.Session.Uid
	dbUser := db.NewMallUser()
	strSql := `
		UPDATE usr_user SET nickname=? WHERE uid=?
	`
	_, err = dbUser.DB.Exec(strSql, m.Ask.Nickname, uid)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update user nickname failed. %s", err)
		return
	}

	// 广播更改消息
	topic, err := mq.GetTopicServiceUserService()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-user-service failed. %s", err)
		return
	}

	err = topic.ModifyUserInfo(&mq.ModifyUserInfoMessage{
		UserId: uid,
		Values: map[string]interface{}{
			"nickname": m.Ask.Nickname,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish service-user-service failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 获取手机验证码
type WxXcxUserMobileVerifyImpl struct {
	cidl.ApiWxXcxUserMobileVerify
}

func AddWxXcxUserMobileVerifyHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_MOBILE_VERIFY,
		func() http.ApiHandler {
			return &WxXcxUserMobileVerifyImpl{
				ApiWxXcxUserMobileVerify: cidl.MakeApiWxXcxUserMobileVerify(),
			}
		},
	)
}

func (m *WxXcxUserMobileVerifyImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	mobile := m.Query.Mobile
	userCache := cache.NewUserCache()
	success, err := userCache.Cache.SetNX(userCache.KeyMobileVerifyWxXcxModifyMobileInterval(mobile), 1, 55*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set wx_xcx modify mobile verify code cache failed. %s", err)
		return
	}

	if !success {
		ctx.Errorf(api.ErrSendSMSIntervalLimit, "time of sending wx_xcx verify code has not arrived. %s", err)
		return
	}

	verifyCode := utils.VerifyCode6()
	err = userCache.Cache.Set(userCache.KeyMobileVerifyWxXcxModifyMobile(mobile), verifyCode, 10*time.Minute+5*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set mobile verify code cache failed. %s", err)
		return
	}

	_, err = sms.NewProxy("sms-service").SmsSendTemplateActionVerify(&sms.AskSmsSendTemplateActionVerify{
		Action: "团长小程序修改手机",
		Mobile: mobile,
		Time:   "10分钟",
		Code:   verifyCode,
	})
	if err != nil {
		ctx.Errorf(api.ErrSendSMSFailed, "send sms verify code failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 修改手机
type WxXcxUserEditMobileByUserIDImpl struct {
	cidl.ApiWxXcxUserEditMobileByUserID
}

func AddWxXcxUserEditMobileByUserIDHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_EDIT_MOBILE_BY_USER_ID,
		func() http.ApiHandler {
			return &WxXcxUserEditMobileByUserIDImpl{
				ApiWxXcxUserEditMobileByUserID: cidl.MakeApiWxXcxUserEditMobileByUserID(),
			}
		},
	)
}

func (m *WxXcxUserEditMobileByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	// 手机验证码
	mobile := m.Ask.Mobile
	userCache := cache.NewUserCache()
	cacheKey := userCache.KeyMobileVerifyWxXcxModifyMobile(mobile)
	value, err := userCache.Cache.Get(cacheKey)
	if err != nil && err != cache2.Nil {
		ctx.Errorf(api.ErrCacheReadFailed, "get mobile verify code cache failed. %s", err)
		return
	}

	// TODO 记得删掉手机验证码
	if "123456" != m.Ask.VerifyCode && (err == cache2.Nil || value != m.Ask.VerifyCode) {
		ctx.Errorf(api.ErrSMSVerifyCodeNotMatch, "verify code does not match.")
		return
	}

	uid := ctx.Session.Uid
	dbUser := db.NewMallUser()
	strSql := `
		UPDATE
			usr_user
		SET
			mobile=?
		WHERE
			uid=?
			AND mobile<>?
	`
	result, err := dbUser.DB.Exec(strSql, m.Ask.Mobile, uid, m.Ask.Mobile)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update user mobile failed. %s", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get row affected failed. %s", err)
		return
	}

	if affected == 0 {
		ctx.Errorf(cidl.ErrWxXcxMobileNotChange, "update user mobile failed. %s", err)
		return
	}

	// 广播更改消息
	topic, err := mq.GetTopicServiceUserService()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-user-service failed. %s", err)
		return
	}

	err = topic.ModifyUserInfo(&mq.ModifyUserInfoMessage{
		UserId: uid,
		Values: map[string]interface{}{
			"mobile": m.Ask.Mobile,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish service-user-service failed. %s", err)
		return
	}

	ctx.Succeed()

	// 删除手机验证码
	_, err = userCache.Cache.Delete(cacheKey)
	if err != nil {
		log.Warnf("delete verify code cache failed. %s", err)
		return
	}
}

// 微信绑定手机号手机验证码
type UserWxXcxUserWxBindMobileVerifyImpl struct {
	cidl.ApiUserWxXcxUserWxBindMobileVerify
}

func AddUserWxXcxUserWxBindMobileVerifyHandler() {
	AddHandler(
		cidl.META_USER_WX_XCX_USER_WX_BIND_MOBILE_VERIFY,
		func() http.ApiHandler {
			return &UserWxXcxUserWxBindMobileVerifyImpl{
				ApiUserWxXcxUserWxBindMobileVerify: cidl.MakeApiUserWxXcxUserWxBindMobileVerify(),
			}
		},
	)
}

func (m *UserWxXcxUserWxBindMobileVerifyImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	mobile := m.Query.Mobile
	userCache := cache.NewUserCache()
	success, err := userCache.Cache.SetNX(userCache.KeyMobileVerifyWxXcxBindMobileInterval(mobile), 1, 55*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set wx_xcx bind mobile verify code cache failed. %s", err)
		return
	}

	if !success {
		ctx.Errorf(api.ErrSendSMSIntervalLimit, "time of sending wx_xcx verify code has not arrived. %s", err)
		return
	}

	verifyCode := utils.VerifyCode6()
	err = userCache.Cache.Set(userCache.KeyMobileVerifyWxXcxBindMobile(mobile), verifyCode, 10*time.Minute+5*time.Second)
	if err != nil {
		ctx.Errorf(api.ErrCacheWriteFailed, "set mobile verify code cache failed. %s", err)
		return
	}

	_, err = sms.NewProxy("sms-service").SmsSendTemplateActionVerify(&sms.AskSmsSendTemplateActionVerify{
		Action: "团长小程序绑定手机",
		Mobile: mobile,
		Time:   "10分钟",
		Code:   verifyCode,
	})
	if err != nil {
		ctx.Errorf(api.ErrSendSMSFailed, "send sms verify code failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 微信绑定手机号
type UserWxXcxUserWxBindMobileImpl struct {
	cidl.ApiUserWxXcxUserWxBindMobile
}

func AddUserWxXcxUserWxBindMobileHandler() {
	AddHandler(
		cidl.META_USER_WX_XCX_USER_WX_BIND_MOBILE,
		func() http.ApiHandler {
			return &UserWxXcxUserWxBindMobileImpl{
				ApiUserWxXcxUserWxBindMobile: cidl.MakeApiUserWxXcxUserWxBindMobile(),
			}
		},
	)
}

func (m *UserWxXcxUserWxBindMobileImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	mobile := m.Ask.Mobile
	userCache := cache.NewUserCache()
	cacheKey := userCache.KeyMobileVerifyWxXcxBindMobile(mobile)
	value, err := userCache.Cache.Get(cacheKey)
	if err != nil && err != cache2.Nil {
		ctx.Errorf(api.ErrCacheReadFailed, "get mobile verify code cache failed. %s", err)
		return
	}

	// TODO 记得删掉假验证码
	if "123456" != m.Ask.VerifyCode && (err == cache2.Nil || value != m.Ask.VerifyCode) {
		ctx.Errorf(api.ErrSMSVerifyCodeNotMatch, "verify code does not match.")
		return
	}

	userId := ctx.Session.Uid
	token := ctx.Session.Token
	auth, err := com.NewAuthWxXcxByToken(userId, token)
	if err != nil {
		ctx.Errorf(api.ErrIllegalToken, "new auth wx xcx by token failed. %s", err)
		return
	}

	if auth.IsVisitor == false {
		ctx.Errorf(api.ErrServer, "only visitor can bind mobile")
		return
	}

	wxUnionId := auth.SessionKey.UnionId
	dbUser := db.NewMallUser()
	user, err := dbUser.GetUserByMobile(mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed.")
		return
	} else if err == conn.ErrNoRows {
		ctx.Errorf(cidl.ErrWxXcxMobileNotExist, "mobile not exist.")
		return
	}

	if user.WxUnionId != "" {
		ctx.Errorf(cidl.ErrWxXcxMobileHasBindWxUnionId, "mobile was bind by other weixin.")
		return
	}

	strSql := `
		UPDATE
			usr_user
		SET
			wx_union_id=?
		WHERE
			mobile=? AND (wx_union_id IS NULL OR wx_union_id="")
	`
	result, err := dbUser.DB.Exec(strSql, wxUnionId, mobile)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update user mobile failed. %s", err)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "get affected rows count failed. %s", err)
		return
	}

	if rowsAffected == 0 {
		ctx.Errorf(api.ErrDBUpdateFailed, "0 user mobile was updated. %s", err)
		return
	}

	user, err = dbUser.GetUserByUnionId(wxUnionId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by union id failed. %s", err)
		return
	}

	// 更新信息
	auth.IsVisitor = false
	err = auth.SetUserInfo(user)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxUpdateTokenFailed, "set auth user info failed. %s", err)
		return
	}

	err = auth.SetToken(token)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxUpdateTokenFailed, "update token failed. %s", err)
		return
	}

	ctx.Succeed()

	// 删除手机验证码
	_, err = userCache.Cache.Delete(cacheKey)
	if err != nil {
		log.Warnf("delete verify code cache failed. %s", err)
	}

	err = mq2.BroadcastInvalidateAuthWxXcxToken(token)
	if err != nil {
		log.Warnf("broadcast invalidate auth wx_xcx token failed. %s", err)
	}

}

type WxXcxUserCheckBindMobileByMobileImpl struct {
	cidl.ApiWxXcxUserCheckBindMobileByMobile
}

func AddWxXcxUserCheckBindMobileByMobileHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_CHECK_BIND_MOBILE_BY_MOBILE,
		func() http.ApiHandler {
			return &WxXcxUserCheckBindMobileByMobileImpl{
				ApiWxXcxUserCheckBindMobileByMobile: cidl.MakeApiWxXcxUserCheckBindMobileByMobile(),
			}
		},
	)
}

func (m *WxXcxUserCheckBindMobileByMobileImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbGroupBuying := db.NewMallUser()
	user, err := dbGroupBuying.GetUserByMobile(m.Params.Mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return

	} else if err == conn.ErrNoRows {
		m.Ack.CanBeBound = true

	} else {
		m.Ack.UserExist = true
		m.Ack.UserName = user.Name

		if user.IsOrgManager || user.IsOrgStaff || user.IsCmtManager {
			m.Ack.CanBeBound = false
		} else {
			m.Ack.CanBeBound = true
		}
	}

	ctx.Json(m.Ack)
}

// 用户解绑微信
type WxXcxUserUnbindWxImpl struct {
	cidl.ApiWxXcxUserUnbindWx
}

func AddWxXcxUserUnbindWxHandler() {
	AddHandler(
		cidl.META_WX_XCX_USER_UNBIND_WX,
		func() http.ApiHandler {
			return &WxXcxUserUnbindWxImpl{
				ApiWxXcxUserUnbindWx: cidl.MakeApiWxXcxUserUnbindWx(),
			}
		},
	)
}

func (m *WxXcxUserUnbindWxImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := ctx.Session.Uid
	token := ctx.Session.Token
	auth, err := com.NewAuthWxXcxByToken(userId, token)
	if err != nil {
		ctx.Errorf(api.ErrIllegalToken, "new auth wx xcx by token failed. %s", err)
		return
	}

	if auth.IsVisitor == true {
		ctx.Errorf(api.ErrServer, "visitor can not unbind wx")
		return
	}

	unionId := auth.WxUserInfo.UnionId

	dbUser := db.NewMallUser()
	strSql := `
		UPDATE
			usr_user
		SET
			wx_union_id=NULL
        WHERE
			wx_union_id=? AND uid=?
	`
	_, err = dbUser.DB.Exec(strSql, unionId, userId)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "unbind user wx failed. %s", err)
		return
	}

	// 更新信息
	auth.IsVisitor = true
	auth.UserId = unionId
	auth.User = nil
	auth.Organization = nil
	auth.CommunityManager = nil
	auth.OrganizationManager = nil
	err = auth.SetToken(token)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxUpdateTokenFailed, "update token failed. %s", err)
		return
	}

	ctx.Succeed()

}

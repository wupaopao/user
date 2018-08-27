package impls

import (
	"common/api"
	"fmt"
	"net/url"
	"time"

	"business/user/cidl"
	"business/user/common/cache"
	"business/user/common/com"
	"business/user/common/wx"

	"github.com/mz-eco/mz/errors"
	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/utils"
)

func init() {
	AddAdminWxWorkLoginWxWorkAuthHandler()
	AddAdminWxWorkLoginWxWorkAuthDataHandler()
	AddAdminWxWorkLoginWxWorkCallbackHandler()
	AddAdminWxWorkLoginLogoutHandler()
	AddAdminWxWorkUserBasicInfoHandler()
}

const KEY_COOKIE_WX_WORK_AUTH = "wx_work_session_id"

// 企业微信授权页面跳转
type AdminWxWorkLoginWxWorkAuthImpl struct {
	cidl.ApiAdminWxWorkLoginWxWorkAuth
}

func AddAdminWxWorkLoginWxWorkAuthHandler() {
	AddHandler(
		cidl.META_ADMIN_WX_WORK_LOGIN_WX_WORK_AUTH,
		func() http.ApiHandler {
			return &AdminWxWorkLoginWxWorkAuthImpl{
				ApiAdminWxWorkLoginWxWorkAuth: cidl.MakeApiAdminWxWorkLoginWxWorkAuth(),
			}
		},
	)
}

func (m *AdminWxWorkLoginWxWorkAuthImpl) Handler(ctx *http.Context) {
	targetRedirectUri := m.Query.RedirectUri
	wxWork := wx.GetWxWork()
	authInfo := GetWxWorkAuthInfo(ctx, targetRedirectUri)
	authUrl := wxWork.GetAuthUrl(authInfo.RedirectUri, authInfo.State)
	ctx.Redirect(authUrl)
}

type WxWorkAuthInfo struct {
	AppId       string `json:"app_id"`
	AgentId     int32  `json:"agent_id"`
	RedirectUri string `json:"redirect_uri"`
	State       string `json:"state"`
}

func GetWxWorkAuthInfo(ctx *http.Context, targetRedirectUri string) (authInfo *WxWorkAuthInfo) {
	wxWork := wx.GetWxWork()
	authRedirectUri := fmt.Sprintf(
		"%s?redirect_uri=%s",
		wxWork.RedirectUri,
		url.PathEscape(targetRedirectUri),
	)

	authUrlState := utils.UniqueID()
	sessionId := utils.UniqueID()
	ctx.Engine.SetCookie(KEY_COOKIE_WX_WORK_AUTH, sessionId, 86400, "/", ctx.Engine.Request.Host, false, false)
	cache.NewUserCache().SetWxWorkAuthSession(sessionId, []byte(authUrlState), 86400*time.Second)

	authInfo = &WxWorkAuthInfo{
		AppId:       wxWork.CorpId,
		AgentId:     wxWork.AgentId,
		RedirectUri: authRedirectUri,
		State:       authUrlState,
	}

	return
}

// 获取企业微信授权数据
type AdminWxWorkLoginWxWorkAuthDataImpl struct {
	cidl.ApiAdminWxWorkLoginWxWorkAuthData
}

func AddAdminWxWorkLoginWxWorkAuthDataHandler() {
	AddHandler(
		cidl.META_ADMIN_WX_WORK_LOGIN_WX_WORK_AUTH_DATA,
		func() http.ApiHandler {
			return &AdminWxWorkLoginWxWorkAuthDataImpl{
				ApiAdminWxWorkLoginWxWorkAuthData: cidl.MakeApiAdminWxWorkLoginWxWorkAuthData(),
			}
		},
	)
}

func (m *AdminWxWorkLoginWxWorkAuthDataImpl) Handler(ctx *http.Context) {
	targetRedirectUri := m.Query.RedirectUri
	authInfo := GetWxWorkAuthInfo(ctx, targetRedirectUri)
	m.Ack.AppId = authInfo.AppId
	m.Ack.AgentId = authInfo.AgentId
	m.Ack.RedirectUri = authInfo.RedirectUri
	m.Ack.State = authInfo.State
	ctx.Json(m.Ack)
}

// 企业微信授权回调
type AdminWxWorkLoginWxWorkCallbackImpl struct {
	cidl.ApiAdminWxWorkLoginWxWorkCallback
}

func AddAdminWxWorkLoginWxWorkCallbackHandler() {
	AddHandler(
		cidl.META_ADMIN_WX_WORK_LOGIN_WX_WORK_CALLBACK,
		func() http.ApiHandler {
			return &AdminWxWorkLoginWxWorkCallbackImpl{
				ApiAdminWxWorkLoginWxWorkCallback: cidl.MakeApiAdminWxWorkLoginWxWorkCallback(),
			}
		},
	)
}

func (m *AdminWxWorkLoginWxWorkCallbackImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	if m.Query.Code == "" { // 用户取消授权或者未授权
		ctx.Error(api.ErrWxWorkUnauthorized, errors.New("用户取消授权或者未授权"))
		return
	}

	// 获取session_id
	sessionId, err := ctx.Engine.Cookie(KEY_COOKIE_WX_WORK_AUTH)
	if err != nil {
		err = errors.New("get wx work auth session id cookie failed. %s", err)
		ctx.Error(api.ErrInvalidParams, err)
		return
	}

	adminCache := cache.NewUserCache()
	authState, err := adminCache.GetWxWorkAuthSession(sessionId)
	if err != nil {
		err = errors.New("get wx work auth session failed. %s", err)
		ctx.Error(api.ErrServer, err)
		return
	}

	if authState != m.Query.State {
		err = errors.New("illegal state")
		ctx.Error(api.ErrWrongParams, err)
		return
	}

	strRedirectUri, err := url.PathUnescape(m.Query.RedirectUri)
	if err != nil {
		err = errors.New("unescape redirect uri failed. %s", err)
		ctx.Error(api.ErrServer, err)
		return
	}

	redirectUriHandler, err := utils.NewUrlHandlerByRawUrl(strRedirectUri)
	if err != nil {
		ctx.Errorf(api.ErrServer, "new redirect uri handler failed. %s", err)
		return
	}

	wxWork := wx.GetWxWork()
	userInfo, err := wxWork.GetUserInfo(m.Query.Code)
	if err != nil {
		ctx.Errorf(api.ErrServer, "get wx work user info failed. %s", err)
		return
	}

	// 删除cookie
	ctx.Engine.SetCookie(KEY_COOKIE_WX_WORK_AUTH, sessionId, -1, "/", ctx.Engine.Request.Host, false, false)

	// 删除授权session state
	err = adminCache.DeleteWxWorkAuthSession(sessionId)
	if err != nil {
		ctx.Errorf(api.ErrServer, "delete wx work auth session failed. %s", err)
		return
	}

	// 获取微信用户信息
	wxUser, err := wxWork.GetUser(userInfo.UserId)
	if err != nil {
		ctx.Errorf(api.ErrServer, "get wx work user failed. %s", err)
		return
	}

	// 记录到系统授权用户信息里
	auth := com.NewAuthAdmin(wxUser)
	token, err := auth.NewToken()
	if err != nil {
		ctx.Errorf(api.ErrServer, "new auth token failed. %s", err)
		return
	}

	redirectUriHandler.SetQuery("token", token)
	redirectUriHandler.SetQuery("uid", wxUser.UserID)

	ctx.Redirect(redirectUriHandler.String())

}

// 登出
type AdminWxWorkLoginLogoutImpl struct {
	cidl.ApiAdminWxWorkLoginLogout
}

func AddAdminWxWorkLoginLogoutHandler() {
	AddHandler(
		cidl.META_ADMIN_WX_WORK_LOGIN_LOGOUT,
		func() http.ApiHandler {
			return &AdminWxWorkLoginLogoutImpl{
				ApiAdminWxWorkLoginLogout: cidl.MakeApiAdminWxWorkLoginLogout(),
			}
		},
	)
}

func (m *AdminWxWorkLoginLogoutImpl) Handler(ctx *http.Context) {
	err := com.DeleteAuthAdminByToken(ctx.Session.Token)
	if err != nil {
		ctx.Errorf(api.ErrServer, "delete auth by token failed. %s", err)
		return
	}
	ctx.Succeed()
}

// 获取用户基本信息
type AdminWxWorkUserBasicInfoImpl struct {
	cidl.ApiAdminWxWorkUserBasicInfo
}

func AddAdminWxWorkUserBasicInfoHandler() {
	AddHandler(
		cidl.META_ADMIN_WX_WORK_USER_BASIC_INFO,
		func() http.ApiHandler {
			return &AdminWxWorkUserBasicInfoImpl{
				ApiAdminWxWorkUserBasicInfo: cidl.MakeApiAdminWxWorkUserBasicInfo(),
			}
		},
	)
}

func (m *AdminWxWorkUserBasicInfoImpl) Handler(ctx *http.Context) {
	var err error
	token := ctx.Session.Token
	uid := ctx.Session.Uid
	auth, err := com.NewAuthAdminByToken(uid, token)
	if err != nil || auth == nil {
		ctx.Errorf(api.ErrIllegalToken, "new auth by token failed. %s", err)
		return
	}

	m.Ack.UserID = auth.User.UserID
	m.Ack.Name = auth.User.Name
	m.Ack.Avatar = auth.User.Avatar

	ctx.Json(m.Ack)
}

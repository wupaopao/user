package impls

import (
	"business/user/cidl"
	"business/user/common/com"

	"github.com/mz-eco/mz/http"
)

func init() {
	AddInnerUserAuthTokenInfoAdminHandler()
	AddInnerUserAuthTokenInfoOrgHandler()
	AddInnerUserAuthTokenInfoWxXcxHandler()
}

type InnerUserAuthTokenInfoAdminImpl struct {
	cidl.ApiInnerUserAuthTokenInfoAdmin
}

func AddInnerUserAuthTokenInfoAdminHandler() {
	AddHandler(
		cidl.META_INNER_USER_AUTH_TOKEN_INFO_ADMIN,
		func() http.ApiHandler {
			return &InnerUserAuthTokenInfoAdminImpl{
				ApiInnerUserAuthTokenInfoAdmin: cidl.MakeApiInnerUserAuthTokenInfoAdmin(),
			}
		},
	)
}

func (m *InnerUserAuthTokenInfoAdminImpl) Handler(ctx *http.Context) {
	var err error
	uid := m.Ask.UserID
	token := m.Ask.Token
	auth, err := com.NewAuthAdminByToken(uid, token)
	if err != nil {
		ctx.Errorf(cidl.ErrAdminIllegalToken, "verify admin token failed. %s", err)
		return
	}
	m.Ack = (*cidl.AuthAdmin)(auth)
	ctx.Json(m.Ack)
}

type InnerUserAuthTokenInfoOrgImpl struct {
	cidl.ApiInnerUserAuthTokenInfoOrg
}

func AddInnerUserAuthTokenInfoOrgHandler() {
	AddHandler(
		cidl.META_INNER_USER_AUTH_TOKEN_INFO_ORG,
		func() http.ApiHandler {
			return &InnerUserAuthTokenInfoOrgImpl{
				ApiInnerUserAuthTokenInfoOrg: cidl.MakeApiInnerUserAuthTokenInfoOrg(),
			}
		},
	)
}

func (m *InnerUserAuthTokenInfoOrgImpl) Handler(ctx *http.Context) {
	var err error
	uid := m.Ask.UserID
	token := m.Ask.Token
	auth, err := com.NewAuthCityByToken(uid, token)
	if err != nil {
		ctx.Errorf(cidl.ErrOrgIllegalToken, "verify city token failed. %s", err)
		return
	}

	m.Ack = (*cidl.AuthCity)(auth)

	ctx.Json(m.Ack)
}

type InnerUserAuthTokenInfoWxXcxImpl struct {
	cidl.ApiInnerUserAuthTokenInfoWxXcx
}

func AddInnerUserAuthTokenInfoWxXcxHandler() {
	AddHandler(
		cidl.META_INNER_USER_AUTH_TOKEN_INFO_WX_XCX,
		func() http.ApiHandler {
			return &InnerUserAuthTokenInfoWxXcxImpl{
				ApiInnerUserAuthTokenInfoWxXcx: cidl.MakeApiInnerUserAuthTokenInfoWxXcx(),
			}
		},
	)
}

func (m *InnerUserAuthTokenInfoWxXcxImpl) Handler(ctx *http.Context) {
	var err error
	uid := m.Ask.UserID
	token := m.Ask.Token
	auth, err := com.NewAuthWxXcxByToken(uid, token)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxIllegalToken, "verify wx xcx token failed. %s", err)
		return
	}

	m.Ack = (*cidl.AuthWxXcx)(auth)

	ctx.Json(m.Ack)
}

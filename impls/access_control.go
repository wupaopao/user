package impls

import (
	"business/auth"

	"github.com/mz-eco/mz/http"
)

var AccessControlHandlers []http.AccessControlHandler

func AddAccessControlHandler(handler http.AccessControlHandler) {
	AccessControlHandlers = append(AccessControlHandlers, handler)
}

func init() {

	// 白名单
	whiteListAuthHandler := auth.NewWhiteListAuthHandler([]string{
		"/api/v1/inner/",
	})

	// 微信小程序端访问控制
	wxXcxAuthHandler := auth.NewWxXcxAuthHandler([]string{
		"/api/v1/user/wx_xcx/",
	}, []string{
		"/api/v1/user/wx_xcx/user/auth",
		"/api/v1/user/wx_xcx/user/auth_fake",      // TODO 测试用，待删除
		"/api/v1/user/wx_xcx/user/fake_user_list", // TODO 测试用，待删除
	})

	// 城市合伙人端访问控制
	orgAuthHandler := auth.NewOrgAuthHandler([]string{
		"/api/v1/user/org/",
	}, []string{
		"/api/v1/user/org/user/login",
		"/api/v1/user/org/user/login_mobile_verify",
	})

	// 运营端访问控制
	adminAuthHandler := auth.NewAdminAuthHandler([]string{
		"/api/v1/user/admin/",
	}, []string{
		"/api/v1/user/admin/wx_work/login/wx_work_auth",
		"/api/v1/user/admin/wx_work/login/wx_work_auth_data",
		"/api/v1/user/admin/wx_work/login/wx_work_callback",
	})

	chanHandlers := &auth.ChainAuthHandler{
		AuthHandlers: []*auth.AuthHandler{
			whiteListAuthHandler,
			wxXcxAuthHandler,
			orgAuthHandler,
			adminAuthHandler,
		},
	}

	AddAccessControlHandler(chanHandlers.AccessControlHandler)
}

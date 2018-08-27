package impls

import (
	"fmt"
	"net/url"
	"time"

	"common/api"
	"common/file"

	"business/user/cidl"
	"business/user/common/db"

	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/utils"
)

func init() {
	AddAdminUserInfoByUserIdHandler()
	AddAdminUserInfoByMobileHandler()
	AddAdminUserOrganizationManagerListHandler()

	AddAdminUserCommunityManagerListHandler()

	AddAdminUserOrganizationStaffListHandler()

	AddAdminUserIDCardPicTokenHandler()
	AddAdminUserAccessIDCardPicHandler()
	AddAdminUserCheckBindMobileByMobileHandler()
}

// 获取用户信息
type AdminUserInfoByUserIdImpl struct {
	cidl.ApiAdminUserInfoByUserId
}

func AddAdminUserInfoByUserIdHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_INFO_BY_USER_ID,
		func() http.ApiHandler {
			return &AdminUserInfoByUserIdImpl{
				ApiAdminUserInfoByUserId: cidl.MakeApiAdminUserInfoByUserId(),
			}
		},
	)
}

func (m *AdminUserInfoByUserIdImpl) Handler(ctx *http.Context) {
	userId := m.Params.UserID
	user, err := db.NewMallUser().GetUser(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get user from db failed. %s", err)
		return
	}
	m.Ack = user
	ctx.Json(m.Ack)
}

// 通过手机号获取用户信息
type AdminUserInfoByMobileImpl struct {
	cidl.ApiAdminUserInfoByMobile
}

func AddAdminUserInfoByMobileHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_INFO_BY_MOBILE,
		func() http.ApiHandler {
			return &AdminUserInfoByMobileImpl{
				ApiAdminUserInfoByMobile: cidl.MakeApiAdminUserInfoByMobile(),
			}
		},
	)
}

func (m *AdminUserInfoByMobileImpl) Handler(ctx *http.Context) {
	user, err := db.NewMallUser().GetUserByMobile(m.Query.Mobile)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get user from db failed. %s", err)
		return
	}
	m.Ack = user
	ctx.Json(m.Ack)
}

// 城市合伙人列表
type AdminUserOrganizationManagerListImpl struct {
	cidl.ApiAdminUserOrganizationManagerList
}

func AddAdminUserOrganizationManagerListHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_ORGANIZATION_MANAGER_LIST,
		func() http.ApiHandler {
			return &AdminUserOrganizationManagerListImpl{
				ApiAdminUserOrganizationManagerList: cidl.MakeApiAdminUserOrganizationManagerList(),
			}
		},
	)
}

func (m *AdminUserOrganizationManagerListImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	ack := m.Ack
	dbUser := db.NewMallUser()
	if m.Query.Search == "" {
		ack.Count, err = dbUser.OrgManagerCount()
	} else {
		ack.Count, err = dbUser.OrgManagerSearchCount(m.Query.Search)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get org manager count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(ack)
		return
	}

	if m.Query.Search == "" {
		ack.List, err = dbUser.OrgManagerList(m.Query.Page, m.Query.PageSize, false)
	} else {
		ack.List, err = dbUser.OrgManagerSearchList(m.Query.Page, m.Query.PageSize, m.Query.Search, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get org manager list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

// 社区合伙人列表
type AdminUserCommunityManagerListImpl struct {
	cidl.ApiAdminUserCommunityManagerList
}

func AddAdminUserCommunityManagerListHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_COMMUNITY_MANAGER_LIST,
		func() http.ApiHandler {
			return &AdminUserCommunityManagerListImpl{
				ApiAdminUserCommunityManagerList: cidl.MakeApiAdminUserCommunityManagerList(),
			}
		},
	)
}

func (m *AdminUserCommunityManagerListImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	ack := m.Ack
	dbUser := db.NewMallUser()

	if m.Query.Search == "" {
		ack.Count, err = dbUser.CmtManagerCount()
	} else {
		ack.Count, err = dbUser.CmtManagerSearchCount(m.Query.Search)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get cmt manager count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(ack)
		return
	}

	if m.Query.Search == "" {
		ack.List, err = dbUser.CmtManagerList(m.Query.Page, m.Query.PageSize, false)
	} else {
		ack.List, err = dbUser.CmtManagerSearchList(m.Query.Page, m.Query.PageSize, m.Query.Search, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get cmt manager list. %s", err)
		return
	}

	ctx.Json(ack)
}

// 组织成员列表
type AdminUserOrganizationStaffListImpl struct {
	cidl.ApiAdminUserOrganizationStaffList
}

func AddAdminUserOrganizationStaffListHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_ORGANIZATION_STAFF_LIST,
		func() http.ApiHandler {
			return &AdminUserOrganizationStaffListImpl{
				ApiAdminUserOrganizationStaffList: cidl.MakeApiAdminUserOrganizationStaffList(),
			}
		},
	)
}

func (m *AdminUserOrganizationStaffListImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	ack := m.Ack
	dbUser := db.NewMallUser()

	if m.Query.Search == "" {
		ack.Count, err = dbUser.OrgStaffCount()
	} else {
		ack.Count, err = dbUser.OrgStaffSearchCount(m.Query.Search)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get org staff count failed. %s", err)
		return
	}

	if ack.Count == 0 {
		ctx.Json(ack)
		return
	}

	if m.Query.Search == "" {
		ack.List, err = dbUser.OrgStaffList(m.Query.Page, m.Query.PageSize, false)
	} else {
		ack.List, err = dbUser.OrgStaffSearchList(m.Query.Page, m.Query.PageSize, m.Query.Search, false)
	}

	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get org staff list failed. %s", err)
		return
	}

	ctx.Json(ack)
}

// 身份证上传token
type AdminUserIDCardPicTokenImpl struct {
	cidl.ApiAdminUserIDCardPicToken
}

func AddAdminUserIDCardPicTokenHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_ID_CARD_PIC_TOKEN,
		func() http.ApiHandler {
			return &AdminUserIDCardPicTokenImpl{
				ApiAdminUserIDCardPicToken: cidl.MakeApiAdminUserIDCardPicToken(),
			}
		},
	)
}

func (m *AdminUserIDCardPicTokenImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	today, err := utils.DayStartTime(time.Now())
	if err != nil {
		ctx.Errorf(api.ErrServer, "get day start time failed. %s", err)
		return
	}

	qiniu, err := file.GetQiniuPrivateBucket()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get qiniu public bucket failed. %s", err)
		return
	}

	prefix := fmt.Sprintf("billimall/admin/id_card_pic/%d/", today.Unix())
	for _, fileName := range m.Ask.FileNames {
		if fileName == "" {
			ctx.Errorf(api.ErrWrongParams, "empty pic file name. %s", err)
			return
		}

		token, key, err := qiniu.GenerateUploadToken(fileName, prefix)
		if err != nil {
			return
		}

		storeUrl := qiniu.StoreUrl(key)
		accessUrl := qiniu.PrivateAccessUrl(storeUrl, 3600)
		m.Ack.Tokens = append(m.Ack.Tokens, &cidl.AckPicToken{
			OriginalFileName: fileName,
			Token:            token,
			Key:              key,
			StoreUrl:         storeUrl,
			AccessUrl:        accessUrl,
		})

	}

	ctx.Json(m.Ack)
}

// 访问身份证
type AdminUserAccessIDCardPicImpl struct {
	cidl.ApiAdminUserAccessIDCardPic
}

func AddAdminUserAccessIDCardPicHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_ACCESS_ID_CARD_PIC,
		func() http.ApiHandler {
			return &AdminUserAccessIDCardPicImpl{
				ApiAdminUserAccessIDCardPic: cidl.MakeApiAdminUserAccessIDCardPic(),
			}
		},
	)
}

func (m *AdminUserAccessIDCardPicImpl) Handler(ctx *http.Context) {
	var err error
	qiniu, err := file.GetQiniuPrivateBucket()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get qiniu private bucket failed. %s", err)
		return
	}
	uri, err := url.PathUnescape(m.Query.Uri)
	if err != nil {
		ctx.Errorf(api.ErrServer, "url path unescape failed. %s", err)
		return
	}

	uri = qiniu.PrivateAccessUrl(uri, 3600)
	ctx.Redirect(uri)
}

// 检查手机是否可以被绑定
type AdminUserCheckBindMobileByMobileImpl struct {
	cidl.ApiAdminUserCheckBindMobileByMobile
}

func AddAdminUserCheckBindMobileByMobileHandler() {
	AddHandler(
		cidl.META_ADMIN_USER_CHECK_BIND_MOBILE_BY_MOBILE,
		func() http.ApiHandler {
			return &AdminUserCheckBindMobileByMobileImpl{
				ApiAdminUserCheckBindMobileByMobile: cidl.MakeApiAdminUserCheckBindMobileByMobile(),
			}
		},
	)
}

func (m *AdminUserCheckBindMobileByMobileImpl) Handler(ctx *http.Context) {
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

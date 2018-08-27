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
	AddOrgUserInfoByUserIdHandler()
	AddOrgUserOrganizationManagerListHandler()
	AddOrgUserCommunityManagerListHandler()
	AddOrgUserOrganizationStaffListHandler()
	AddOrgUserIDCardPicTokenHandler()
	AddOrgUserAccessIDCardPicHandler()
	AddOrgUserCheckBindMobileByMobileHandler()
}

// 获取用户信息
type OrgUserInfoByUserIdImpl struct {
	cidl.ApiOrgUserInfoByUserId
}

func AddOrgUserInfoByUserIdHandler() {
	AddHandler(
		cidl.META_ORG_USER_INFO_BY_USER_ID,
		func() http.ApiHandler {
			return &OrgUserInfoByUserIdImpl{
				ApiOrgUserInfoByUserId: cidl.MakeApiOrgUserInfoByUserId(),
			}
		},
	)
}

func (m *OrgUserInfoByUserIdImpl) Handler(ctx *http.Context) {
	userId := m.Params.UserID
	user, err := db.NewMallUser().GetUser(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "获取用户信息失败. %s", err)
		return
	}
	m.Ack = user
	ctx.Json(m.Ack)
}

// 城市合伙人列表
type OrgUserOrganizationManagerListImpl struct {
	cidl.ApiOrgUserOrganizationManagerList
}

func AddOrgUserOrganizationManagerListHandler() {
	AddHandler(
		cidl.META_ORG_USER_ORGANIZATION_MANAGER_LIST,
		func() http.ApiHandler {
			return &OrgUserOrganizationManagerListImpl{
				ApiOrgUserOrganizationManagerList: cidl.MakeApiOrgUserOrganizationManagerList(),
			}
		},
	)
}

func (m *OrgUserOrganizationManagerListImpl) Handler(ctx *http.Context) {
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
type OrgUserCommunityManagerListImpl struct {
	cidl.ApiOrgUserCommunityManagerList
}

func AddOrgUserCommunityManagerListHandler() {
	AddHandler(
		cidl.META_ORG_USER_COMMUNITY_MANAGER_LIST,
		func() http.ApiHandler {
			return &OrgUserCommunityManagerListImpl{
				ApiOrgUserCommunityManagerList: cidl.MakeApiOrgUserCommunityManagerList(),
			}
		},
	)
}

func (m *OrgUserCommunityManagerListImpl) Handler(ctx *http.Context) {
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
type OrgUserOrganizationStaffListImpl struct {
	cidl.ApiOrgUserOrganizationStaffList
}

func AddOrgUserOrganizationStaffListHandler() {
	AddHandler(
		cidl.META_ORG_USER_ORGANIZATION_STAFF_LIST,
		func() http.ApiHandler {
			return &OrgUserOrganizationStaffListImpl{
				ApiOrgUserOrganizationStaffList: cidl.MakeApiOrgUserOrganizationStaffList(),
			}
		},
	)
}

func (m *OrgUserOrganizationStaffListImpl) Handler(ctx *http.Context) {
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

// 身份证Token
type OrgUserIDCardPicTokenImpl struct {
	cidl.ApiOrgUserIDCardPicToken
}

func AddOrgUserIDCardPicTokenHandler() {
	AddHandler(
		cidl.META_ORG_USER_ID_CARD_PIC_TOKEN,
		func() http.ApiHandler {
			return &OrgUserIDCardPicTokenImpl{
				ApiOrgUserIDCardPicToken: cidl.MakeApiOrgUserIDCardPicToken(),
			}
		},
	)
}

func (m *OrgUserIDCardPicTokenImpl) Handler(ctx *http.Context) {
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

	prefix := fmt.Sprintf("billimall/org/id_card_pic/%d/", today.Unix())
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

// 访问身份证图片
type OrgUserAccessIDCardPicImpl struct {
	cidl.ApiOrgUserAccessIDCardPic
}

func AddOrgUserAccessIDCardPicHandler() {
	AddHandler(
		cidl.META_ORG_USER_ACCESS_ID_CARD_PIC,
		func() http.ApiHandler {
			return &OrgUserAccessIDCardPicImpl{
				ApiOrgUserAccessIDCardPic: cidl.MakeApiOrgUserAccessIDCardPic(),
			}
		},
	)
}

func (m *OrgUserAccessIDCardPicImpl) Handler(ctx *http.Context) {
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

type OrgUserCheckBindMobileByMobileImpl struct {
	cidl.ApiOrgUserCheckBindMobileByMobile
}

func AddOrgUserCheckBindMobileByMobileHandler() {
	AddHandler(
		cidl.META_ORG_USER_CHECK_BIND_MOBILE_BY_MOBILE,
		func() http.ApiHandler {
			return &OrgUserCheckBindMobileByMobileImpl{
				ApiOrgUserCheckBindMobileByMobile: cidl.MakeApiOrgUserCheckBindMobileByMobile(),
			}
		},
	)
}

func (m *OrgUserCheckBindMobileByMobileImpl) Handler(ctx *http.Context) {
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

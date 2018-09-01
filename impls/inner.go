package impls

import (
	"fmt"

	mq2 "business/auth/common/mq"
	"business/user/cidl"
	"business/user/common/com"
	"business/user/common/db"
	"business/user/common/mq"
	"common/api"

	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/http"
	"github.com/mz-eco/mz/log"
)

func init() {
	AddInnerUserInfoByUserIDHandler()
	AddInnerUserUserInfoByMobileHandler()
	AddInnerUserCommunityManagerEditByUserIDHandler()
	AddInnerUserUserOrgManagerAddOrUpdateHandler()
	AddInnerUserUserOrgManagerUpdateHandler()
	AddInnerUserUserOrgManagerUnbindHandler()
	AddInnerUserUserOrgStaffAddOrUpdateHandler()
	AddInnerUserUserOrgStaffUpdateByUserIDHandler()
	AddInnerUserUserCmtManagerUpdateByUserIDHandler()
	AddInnerUserUserCmtManagerChangeHandler()
	AddInnerUserUserCmtManagerGetOrAddHandler()
	AddInnerUserUserCmtManagerUnbindHandler()
	AddInnerUserWxXcxCmtManagerGetOrAddHandler()
	AddInnerUserWxXcxRefreshTokenHandler()
	AddInnerUserSetIsDisableByUserIDHandler()
}

// 获取用户信息
type InnerUserInfoByUserIDImpl struct {
	cidl.ApiInnerUserInfoByUserID
}

func AddInnerUserInfoByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_USER_INFO_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerUserInfoByUserIDImpl{
				ApiInnerUserInfoByUserID: cidl.MakeApiInnerUserInfoByUserID(),
			}
		},
	)
}

func (m *InnerUserInfoByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userId := m.Params.UserID
	dbUser := db.NewMallUser()
	user, err := dbUser.GetUser(userId)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get user from db failed. %s", err)
		return
	}

	m.Ack = user
	ctx.Json(m.Ack)
}

// 通过手机号获取用户
type InnerUserUserInfoByMobileImpl struct {
	cidl.ApiInnerUserUserInfoByMobile
}

func AddInnerUserUserInfoByMobileHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_INFO_BY_MOBILE,
		func() http.ApiHandler {
			return &InnerUserUserInfoByMobileImpl{
				ApiInnerUserUserInfoByMobile: cidl.MakeApiInnerUserUserInfoByMobile(),
			}
		},
	)
}

func (m *InnerUserUserInfoByMobileImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbUser := db.NewMallUser()
	user, err := dbUser.GetUserByMobile(m.Params.Mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return

	} else if err == conn.ErrNoRows {
		m.Ack.Exist = false
		m.Ack.User = nil
		ctx.Json(m.Ack)
		return

	}

	m.Ack.Exist = true
	m.Ack.User = user

	ctx.Json(m.Ack)
}

// 设置用户为社区合伙人

type InnerUserCommunityManagerEditByUserIDImpl struct {
	cidl.ApiInnerUserCommunityManagerEditByUserID
}

func AddInnerUserCommunityManagerEditByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_USER_COMMUNITY_MANAGER_EDIT_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerUserCommunityManagerEditByUserIDImpl{
				ApiInnerUserCommunityManagerEditByUserID: cidl.MakeApiInnerUserCommunityManagerEditByUserID(),
			}
		},
	)
}

func (m *InnerUserCommunityManagerEditByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	uid := m.Params.UserID
	dbUser := db.NewMallUser()
	_, err = dbUser.EditIsCmtManager(uid, m.Ask.IsCmtManager)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "edit is_cmt_manager failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 添加组织管理员
type InnerUserUserOrgManagerAddOrUpdateImpl struct {
	cidl.ApiInnerUserUserOrgManagerAddOrUpdate
}

func AddInnerUserUserOrgManagerAddOrUpdateHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_ORG_MANAGER_ADD_OR_UPDATE,
		func() http.ApiHandler {
			return &InnerUserUserOrgManagerAddOrUpdateImpl{
				ApiInnerUserUserOrgManagerAddOrUpdate: cidl.MakeApiInnerUserUserOrgManagerAddOrUpdate(),
			}
		},
	)
}

func (m *InnerUserUserOrgManagerAddOrUpdateImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	mobile := m.Ask.Mobile
	dbUser := db.NewMallUser()
	user, err := dbUser.GetUserByMobile(mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return
	} else if err == conn.ErrNoRows { // 没有此用户
		user = nil
	}

	if user != nil { // 更新

		if user.IsCmtManager || user.IsOrgStaff || user.IsOrgManager {
			ctx.Errorf(cidl.ErrUserWasBound, "user was bound. %s", err)
			return
		}

		m.Ack.UserId = user.UserID
		strSql := `
			UPDATE usr_user
			SET
				name=?,
				mobile=?,
				nickname=?,
				id_card_number=?,
				id_card_front=?,
				id_card_back=?,
				is_org_manager=1,
				is_org_staff=1
			WHERE
				uid=?
				AND is_org_manager=0
				AND is_org_staff=0
				AND is_cmt_manager=0
		`
		result, errUpdate := dbUser.DB.Exec(strSql,
			m.Ask.Name,
			m.Ask.Mobile,
			m.Ask.Nickname,
			m.Ask.IdCardNumber,
			m.Ask.IdCardFront,
			m.Ask.IdCardBack,
			user.UserID)

		err = errUpdate
		if err != nil {
			ctx.Errorf(api.ErrDBUpdateFailed, "update org manager user failed. %s", err)
			return
		}

		rowsAffected, errUpdate := result.RowsAffected()
		err = errUpdate
		if err != nil || rowsAffected == 0 {
			ctx.Errorf(api.ErrDBUpdateFailed, "update org manager user failed. %s", err)
			return
		}

	} else { // 添加新用户
		user = &cidl.User{
			Name:         m.Ask.Name,
			Mobile:       m.Ask.Mobile,
			Nickname:     m.Ask.Nickname,
			IdCardNumber: m.Ask.IdCardNumber,
			IdCardFront:  m.Ask.IdCardFront,
			IdCardBack:   m.Ask.IdCardBack,
			IsOrgManager: true,
			IsOrgStaff:   true,
		}

		result, err := dbUser.AddUser(user)
		if err != nil {
			ctx.Errorf(api.ErrDBInsertFailed, "add user failed. %s", err)
			return
		}

		userId, err := result.LastInsertId()
		if err != nil {
			ctx.Errorf(api.ErrServer, "get add user last insert id failed. %s", err)
			return
		}

		m.Ack.UserId = fmt.Sprintf("%d", userId)
	}

	ctx.Json(m.Ack)
}

// 更新组织管理员
type InnerUserUserOrgManagerUpdateImpl struct {
	cidl.ApiInnerUserUserOrgManagerUpdate
}

func AddInnerUserUserOrgManagerUpdateHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_ORG_MANAGER_UPDATE,
		func() http.ApiHandler {
			return &InnerUserUserOrgManagerUpdateImpl{
				ApiInnerUserUserOrgManagerUpdate: cidl.MakeApiInnerUserUserOrgManagerUpdate(),
			}
		},
	)
}

func (m *InnerUserUserOrgManagerUpdateImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Ask.UserId
	mobile := m.Ask.Mobile
	dbUser := db.NewMallUser()
	strSql := `
		UPDATE usr_user
		SET
			name=?,
			nickname=?,
			id_card_number=?,
			id_card_front=?,
			id_card_back=?
		WHERE
			uid=? AND mobile=? AND is_org_manager=1 AND is_org_staff=1
	`
	_, err = dbUser.DB.Exec(strSql,
		m.Ask.Name,
		m.Ask.Nickname,
		m.Ask.IdCardNumber,
		m.Ask.IdCardFront,
		m.Ask.IdCardBack,
		userId,
		mobile,
	)

	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update org manager user failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 解绑组织管理员用户
type InnerUserUserOrgManagerUnbindImpl struct {
	cidl.ApiInnerUserUserOrgManagerUnbind
}

func AddInnerUserUserOrgManagerUnbindHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_ORG_MANAGER_UNBIND,
		func() http.ApiHandler {
			return &InnerUserUserOrgManagerUnbindImpl{
				ApiInnerUserUserOrgManagerUnbind: cidl.MakeApiInnerUserUserOrgManagerUnbind(),
			}
		},
	)
}

func (m *InnerUserUserOrgManagerUnbindImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbUser := db.NewMallUser()
	strSql := `
		UPDATE
			usr_user
		SET
			is_org_manager=0,
			is_org_staff=0
		WHERE
			uid=? AND is_org_manager=1
	`
	_, err = dbUser.DB.Exec(strSql, m.Ask.OldManagerUid)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update old org manager failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 添加组织成员
type InnerUserUserOrgStaffAddOrUpdateImpl struct {
	cidl.ApiInnerUserUserOrgStaffAddOrUpdate
}

func AddInnerUserUserOrgStaffAddOrUpdateHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_ORG_STAFF_ADD_OR_UPDATE,
		func() http.ApiHandler {
			return &InnerUserUserOrgStaffAddOrUpdateImpl{
				ApiInnerUserUserOrgStaffAddOrUpdate: cidl.MakeApiInnerUserUserOrgStaffAddOrUpdate(),
			}
		},
	)
}

func (m *InnerUserUserOrgStaffAddOrUpdateImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	mobile := m.Ask.Mobile
	dbUser := db.NewMallUser()
	user, err := dbUser.GetUserByMobile(mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return
	} else if err == conn.ErrNoRows {
		user = nil
	}

	if user != nil { // 更新
		if user.IsCmtManager || user.IsOrgStaff || user.IsOrgManager {
			ctx.Errorf(cidl.ErrUserWasBound, "user was bound. %s", err)
			return
		}

		m.Ack.UserId = user.UserID
		strSql := `
			UPDATE usr_user
			SET
				name=?,
				mobile=?,
				is_org_staff=1
			WHERE
				uid=?
				AND is_org_manager=0
				AND is_org_staff=0
				AND is_cmt_manager=0
		`
		_, err = dbUser.DB.Exec(strSql,
			m.Ask.Name,
			m.Ask.Mobile,
			user.UserID)

		if err != nil {
			ctx.Errorf(api.ErrDBUpdateFailed, "update org staff user failed. %s", err)
			return
		}

	} else { // 添加新用户
		user = &cidl.User{
			Name:       m.Ask.Name,
			Mobile:     m.Ask.Mobile,
			IsOrgStaff: true,
		}

		result, err := dbUser.AddUser(user)
		if err != nil {
			ctx.Errorf(api.ErrDBInsertFailed, "add user failed. %s", err)
			return
		}

		userId, err := result.LastInsertId()
		if err != nil {
			ctx.Errorf(api.ErrServer, "get add user last insert id failed. %s", err)
			return
		}

		m.Ack.UserId = fmt.Sprintf("%d", userId)
	}

	ctx.Json(m.Ack)
}

// 更新组织成员信息
type InnerUserUserOrgStaffUpdateByUserIDImpl struct {
	cidl.ApiInnerUserUserOrgStaffUpdateByUserID
}

func AddInnerUserUserOrgStaffUpdateByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_ORG_STAFF_UPDATE_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerUserUserOrgStaffUpdateByUserIDImpl{
				ApiInnerUserUserOrgStaffUpdateByUserID: cidl.MakeApiInnerUserUserOrgStaffUpdateByUserID(),
			}
		},
	)
}

func (m *InnerUserUserOrgStaffUpdateByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbUser := db.NewMallUser()
	userId := m.Params.UserID
	strSql := `UPDATE usr_user SET name=?, mobile=? WHERE uid=? AND is_org_staff=1`
	_, err = dbUser.DB.Exec(strSql, m.Ask.Name, m.Ask.Mobile, userId)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update org staff user failed. %s", err)
		return
	}

	// 广播更改消息
	topic, err := mq.GetTopicServiceUserService()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-user-service failed. %s", err)
		return
	}

	err = topic.ModifyUserInfo(&mq.ModifyUserInfoMessage{
		UserId: userId,
		Values: map[string]interface{}{
			"name":   m.Ask.Name,
			"mobile": m.Ask.Mobile,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "publish topic service-user-service message failed. %s", err)
		return
	}

	ctx.Succeed()
}

type InnerUserUserCmtManagerUpdateByUserIDImpl struct {
	cidl.ApiInnerUserUserCmtManagerUpdateByUserID
}

func AddInnerUserUserCmtManagerUpdateByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_CMT_MANAGER_UPDATE_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerUserUserCmtManagerUpdateByUserIDImpl{
				ApiInnerUserUserCmtManagerUpdateByUserID: cidl.MakeApiInnerUserUserCmtManagerUpdateByUserID(),
			}
		},
	)
}

func (m *InnerUserUserCmtManagerUpdateByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userId := m.Params.UserID
	dbUser := db.NewMallUser()
	strSql := `
		UPDATE usr_user
		SET
			name=?,
			id_card_number=?,
			id_card_front=?,
			id_card_back=?
		WHERE uid=? AND is_cmt_manager=1
	`
	_, err = dbUser.DB.Exec(strSql,
		m.Ask.Name,
		m.Ask.IdCardNumber,
		m.Ask.IdCardFront,
		m.Ask.IdCardBack,
		userId,
	)

	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update community manager user info failed. %s", err)
		return
	}

	// 广播更改消息
	topic, err := mq.GetTopicServiceUserService()
	if err != nil {
		ctx.Errorf(api.ErrMqConnectFailed, "get topic service-user-service failed. %s", err)
		return
	}

	err = topic.ModifyUserInfo(&mq.ModifyUserInfoMessage{
		UserId: userId,
		Values: map[string]interface{}{
			"name":           m.Ask.Name,
			"id_card_number": m.Ask.IdCardNumber,
			"id_card_front":  m.Ask.IdCardFront,
			"id_card_back":   m.Ask.IdCardBack,
		},
	})

	if err != nil {
		ctx.Errorf(api.ErrMqPublishFailed, "public topic service-user-service message failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 更改社群管理员
type InnerUserUserCmtManagerChangeImpl struct {
	cidl.ApiInnerUserUserCmtManagerChange
}

func AddInnerUserUserCmtManagerChangeHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_CMT_MANAGER_CHANGE,
		func() http.ApiHandler {
			return &InnerUserUserCmtManagerChangeImpl{
				ApiInnerUserUserCmtManagerChange: cidl.MakeApiInnerUserUserCmtManagerChange(),
			}
		},
	)
}

func (m *InnerUserUserCmtManagerChangeImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbUser := db.NewMallUser()
	strSql := `
		UPDATE
			usr_user
		SET
			is_cmt_manager=0
		WHERE uid=? AND is_cmt_manager=1
	`
	_, err = dbUser.DB.Exec(strSql, m.Ask.OldManagerUid)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update old cmt manager failed. %s", err)
		return
	}

	strSql = `
		UPDATE
			usr_user
		SET is_cmt_manager=1
		WHERE uid=? AND is_cmt_manager=0
	`
	_, err = dbUser.DB.Exec(strSql, m.Ask.NewManagerUid)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update new cmt manager failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 添加社区合伙人
type InnerUserUserCmtManagerGetOrAddImpl struct {
	cidl.ApiInnerUserUserCmtManagerGetOrAdd
}

func AddInnerUserUserCmtManagerGetOrAddHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_CMT_MANAGER_GET_OR_ADD,
		func() http.ApiHandler {
			return &InnerUserUserCmtManagerGetOrAddImpl{
				ApiInnerUserUserCmtManagerGetOrAdd: cidl.MakeApiInnerUserUserCmtManagerGetOrAdd(),
			}
		},
	)
}

func (m *InnerUserUserCmtManagerGetOrAddImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	dbUser := db.NewMallUser()
	var user *cidl.User
	user, err = dbUser.GetUserByMobile(m.Ask.Mobile)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
		return
	}

	// 如果用户不存在则新增
	if err == conn.ErrNoRows {
		user = &cidl.User{
			Name:         m.Ask.Name,
			Mobile:       m.Ask.Mobile,
			IdCardNumber: m.Ask.IdCardNumber,
			IdCardFront:  m.Ask.IdCardFront,
			IdCardBack:   m.Ask.IdCardBack,
			IsCmtManager: true,
		}

		_, err = dbUser.AddUser(user)
		if err != nil {
			ctx.Errorf(api.ErrDBInsertFailed, "add new community manager user failed. %s", err)
			return
		}

		user, err = dbUser.GetUserByMobile(m.Ask.Mobile)
		if err != nil {
			ctx.Errorf(api.ErrDbQueryFailed, "get user by mobile failed. %s", err)
			return
		}

		m.Ack.IsNew = true
		m.Ack.User = user
		ctx.Json(m.Ack)

		return
	}

	// 如果用户已经绑定用户类型
	if user.IsOrgManager || user.IsCmtManager || user.IsOrgStaff {
		ctx.Errorf(cidl.ErrUserWasBound, "bind user to cmt manager failed.")
		return
	}

	// 如果用户未绑定则绑定
	strSql := `
		UPDATE
			usr_user
		SET
			name=?,
			id_card_number=?,
			id_card_front=?,
			id_card_back=?,
			is_cmt_manager=?
		WHERE
			uid=?
			AND mobile=?
			AND is_org_manager=0
			AND is_org_staff=0
			AND is_cmt_manager=0
	`
	result, err := dbUser.DB.Exec(strSql,
		m.Ask.Name,
		m.Ask.IdCardNumber,
		m.Ask.IdCardFront,
		m.Ask.IdCardBack,
		true,
		user.UserID,
		user.Mobile)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update cmt manager user failed. %s", err)
		return
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		ctx.Errorf(api.ErrServer, "get rows affected failed. %s", err)
		return
	}

	if rowsCount == 0 {
		ctx.Errorf(api.ErrDBUpdateFailed, "update cmt manager user failed. %s", err)
		return
	}

	user, err = dbUser.GetUser(user.UserID)
	if err != nil {
		ctx.Errorf(api.ErrDbQueryFailed, "get user failed. %s", err)
		return
	}

	m.Ack.IsNew = false
	m.Ack.User = user
	ctx.Json(m.Ack)
}

// 解绑社区合伙人
type InnerUserUserCmtManagerUnbindImpl struct {
	cidl.ApiInnerUserUserCmtManagerUnbind
}

func AddInnerUserUserCmtManagerUnbindHandler() {
	AddHandler(
		cidl.META_INNER_USER_USER_CMT_MANAGER_UNBIND,
		func() http.ApiHandler {
			return &InnerUserUserCmtManagerUnbindImpl{
				ApiInnerUserUserCmtManagerUnbind: cidl.MakeApiInnerUserUserCmtManagerUnbind(),
			}
		},
	)
}

func (m *InnerUserUserCmtManagerUnbindImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	dbUser := db.NewMallUser()
	strSql := `
		UPDATE
			usr_user
		SET
			is_cmt_manager=0
		WHERE
			uid=? AND is_cmt_manager=1
	`
	_, err = dbUser.DB.Exec(strSql, m.Ask.OldManagerUid)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update old cmt manager failed. %s", err)
		return
	}

	ctx.Succeed()
}

// 微信接口, 添加社区合伙人
type InnerUserWxXcxCmtManagerGetOrAddImpl struct {
	cidl.ApiInnerUserWxXcxCmtManagerGetOrAdd
}

func AddInnerUserWxXcxCmtManagerGetOrAddHandler() {
	AddHandler(
		cidl.META_INNER_USER_WX_XCX_CMT_MANAGER_GET_OR_ADD,
		func() http.ApiHandler {
			return &InnerUserWxXcxCmtManagerGetOrAddImpl{
				ApiInnerUserWxXcxCmtManagerGetOrAdd: cidl.MakeApiInnerUserWxXcxCmtManagerGetOrAdd(),
			}
		},
	)
}

func (m *InnerUserWxXcxCmtManagerGetOrAddImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userIdOrUnionId := m.Ask.UserId
	token := m.Ask.Token
	auth, err := com.NewAuthWxXcxByToken(userIdOrUnionId, token)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxIllegalToken, "verify auth token failed. %s", err)
		return
	}

	wxUnionId := auth.SessionKey.UnionId

	dbUser := db.NewMallUser()
	var user *cidl.User
	if auth.IsVisitor {

		// 判断是否重复提交或者wx_unionId或者mobile已经绑定其他用户
		user, err = dbUser.GetUserByUnionIdOrMobile(wxUnionId, m.Ask.Mobile)
		if err != nil && err != conn.ErrNoRows {
			ctx.Errorf(api.ErrDbQueryFailed, "get user by wx_unioinid or mobile failed. %s", err)
			return
		}

		if user != nil {
			ctx.Errorf(cidl.ErrWxXcxUserExists, "user has already exist.")
			return
		}

		user = &cidl.User{
			WxUnionId:    wxUnionId,
			Nickname:     auth.WxUserInfo.Nickname,
			Avatar:       auth.WxUserInfo.AvatarUrl,
			Name:         m.Ask.Name,
			Mobile:       m.Ask.Mobile,
			IsCmtManager: true,
		}

		_, err = dbUser.AddUser(user)
		if err != nil {
			ctx.Errorf(api.ErrDBInsertFailed, "add new community manager user failed. %s", err)
			return
		}

		user, err = dbUser.GetUserByUnionId(wxUnionId)
		if err != nil {
			ctx.Errorf(api.ErrDbQueryFailed, "get user failed. %s", err)
			return
		}

		m.Ack.IsNew = true

	} else {
		user = auth.User
		strSql := `
			UPDATE
				usr_user
			SET
				nickname=?,
				avatar=?,
				name=?,
				mobile=?,
				is_cmt_manager=?
			WHERE
				uid=? AND wx_union_id=? AND is_org_manager=0 AND is_org_staff=0 AND is_cmt_manager=0
		`
		result, err := dbUser.DB.Exec(strSql,
			auth.WxUserInfo.Nickname,
			auth.WxUserInfo.AvatarUrl,
			m.Ask.Name,
			m.Ask.Mobile,
			true,
			user.UserID,
			wxUnionId)

		if err != nil {
			ctx.Errorf(api.ErrDBUpdateFailed, "edit user is_cmt_manager failed.", err)
			return
		}
		rowsCount, err := result.RowsAffected()
		if err != nil {
			ctx.Errorf(api.ErrServer, "get rows affected failed. %s", err)
			return
		}

		if rowsCount == 0 {
			ctx.Errorf(api.ErrDBUpdateFailed, "update user failed. %s", err)
			return
		}

		user, err = dbUser.GetUser(user.UserID)
		if err != nil {
			ctx.Errorf(api.ErrDbQueryFailed, "get user failed. %s", err)
			return
		}

		m.Ack.IsNew = false
	}

	m.Ack.User = user
	ctx.Json(m.Ack)
}

// 刷新token
type InnerUserWxXcxRefreshTokenImpl struct {
	cidl.ApiInnerUserWxXcxRefreshToken
}

func AddInnerUserWxXcxRefreshTokenHandler() {
	AddHandler(
		cidl.META_INNER_USER_WX_XCX_REFRESH_TOKEN,
		func() http.ApiHandler {
			return &InnerUserWxXcxRefreshTokenImpl{
				ApiInnerUserWxXcxRefreshToken: cidl.MakeApiInnerUserWxXcxRefreshToken(),
			}
		},
	)
}

func (m *InnerUserWxXcxRefreshTokenImpl) Handler(ctx *http.Context) {
	var (
		err error
	)

	userIdOrUnionId := m.Ask.UserId
	token := m.Ask.Token
	auth, err := com.NewAuthWxXcxByToken(userIdOrUnionId, token)
	if err != nil {
		ctx.Errorf(cidl.ErrWxXcxIllegalToken, "verify auth token failed. %s", err)
		return
	}

	wxUnionId := auth.SessionKey.UnionId
	dbUser := db.NewMallUser()
	isVisitor := false
	user, err := dbUser.GetUserByUnionId(wxUnionId)
	if err != nil && err != conn.ErrNoRows {
		ctx.Errorf(api.ErrDbQueryFailed, "get user by union id failed. %s", err)
		return
	} else if err == conn.ErrNoRows {
		user = nil
		isVisitor = true
	}

	auth.IsVisitor = isVisitor
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

	err = mq2.BroadcastInvalidateAuthWxXcxToken(token)
	if err != nil {
		log.Warnf("broadcast invalidate auth wx_xcx token failed. %s", err)
		return
	}

}

// 禁用用户
type InnerUserSetIsDisableByUserIDImpl struct {
	cidl.ApiInnerUserSetIsDisableByUserID
}

func AddInnerUserSetIsDisableByUserIDHandler() {
	AddHandler(
		cidl.META_INNER_USER_SET_IS_DISABLE_BY_USER_ID,
		func() http.ApiHandler {
			return &InnerUserSetIsDisableByUserIDImpl{
				ApiInnerUserSetIsDisableByUserID: cidl.MakeApiInnerUserSetIsDisableByUserID(),
			}
		},
	)
}

func (m *InnerUserSetIsDisableByUserIDImpl) Handler(ctx *http.Context) {
	var (
		err error
	)
	userId := m.Params.UserID
	userType := m.Ask.UserType
	isDisable := m.Ask.IsDisable

	dbUser := db.NewMallUser()
	_, err = dbUser.UpdateUserDisableState(userId,userType,isDisable)
	if err != nil {
		ctx.Errorf(api.ErrDBUpdateFailed, "update user db failed. %s", err)
		return
	}

	ctx.Succeed()
}

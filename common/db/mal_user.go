package db

import (
	"database/sql"
	"fmt"

	"business/user/cidl"

	"github.com/mz-eco/mz/conn"
	"github.com/mz-eco/mz/errors"
	"github.com/mz-eco/mz/log"
)

type MallUser struct {
	DB *conn.DB
}

func NewMallUser() *MallUser {
	return &MallUser{
		DB: conn.NewDB("mal_user"),
	}
}

// 获取用户
func (m *MallUser) GetUser(userID string) (user *cidl.User, err error) {
	user = &cidl.User{}
	var wx_unoin_id sql.NullString
	var mobile sql.NullString
	strSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM
			usr_user
		WHERE uid=?
	`
	queryRow, err := m.DB.QueryRow(strSql, userID)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&user.UserID,
		&wx_unoin_id,
		&user.Nickname,
		&user.Avatar,
		&user.Name,
		&mobile,
		&user.IdCardNumber,
		&user.IdCardFront,
		&user.IdCardBack,
		&user.IsOrgManager,
		&user.IsOrgStaff,
		&user.IsCmtManager,
		&user.IsDisableOrgManager,
		&user.IsDisableOrgStaff,
		&user.IsDisableCmtManger,
		&user.CreateTime,
	)

	if err != nil {
		user = nil
		if err != conn.ErrNoRows {
			log.Warnf("get user from db failed. %s", err)
		}
		return
	}

	user.WxUnionId = wx_unoin_id.String
	user.Mobile = mobile.String

	return
}

func (m *MallUser) GetUserByMobile(mobile string) (user *cidl.User, err error) {
	user = &cidl.User{}
	var wx_unoin_id sql.NullString
	var mobileNullString sql.NullString
	strSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM
			usr_user
		WHERE mobile=?
	`
	queryRow, err := m.DB.QueryRow(strSql, mobile)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&user.UserID,
		&wx_unoin_id,
		&user.Nickname,
		&user.Avatar,
		&user.Name,
		&mobileNullString,
		&user.IdCardNumber,
		&user.IdCardFront,
		&user.IdCardBack,
		&user.IsOrgManager,
		&user.IsOrgStaff,
		&user.IsCmtManager,
		&user.IsDisableOrgManager,
		&user.IsDisableOrgStaff,
		&user.IsDisableCmtManger,
		&user.CreateTime,
	)

	if err != nil {
		user = nil
		if err != conn.ErrNoRows {
			log.Warnf("get user from db failed. %s", err)
		}
		return
	}

	user.WxUnionId = wx_unoin_id.String
	user.Mobile = mobileNullString.String

	return
}

func (m *MallUser) GetUserByUserIdAndMobile(userId string, mobile string) (user *cidl.User, err error) {
	user = &cidl.User{}
	var wx_unoin_id sql.NullString
	var mobileNullString sql.NullString
	strSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM
			usr_user
		WHERE uid=? AND mobile=?
	`
	queryRow, err := m.DB.QueryRow(strSql, userId, mobile)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&user.UserID,
		&wx_unoin_id,
		&user.Nickname,
		&user.Avatar,
		&user.Name,
		&mobileNullString,
		&user.IdCardNumber,
		&user.IdCardFront,
		&user.IdCardBack,
		&user.IsOrgManager,
		&user.IsOrgStaff,
		&user.IsCmtManager,
		&user.IsDisableOrgManager,
		&user.IsDisableOrgStaff,
		&user.IsDisableCmtManger,
		&user.CreateTime,
	)

	if err != nil {
		user = nil
		if err != conn.ErrNoRows {
			log.Warnf("get user from db failed. %s", err)
		}
		return
	}

	user.WxUnionId = wx_unoin_id.String
	user.Mobile = mobileNullString.String

	return
}

func (m *MallUser) GetUserByUnionId(unionId string) (user *cidl.User, err error) {
	user = &cidl.User{}
	var wx_unoin_id sql.NullString
	var mobileNullString sql.NullString
	strSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM
			usr_user
		WHERE wx_union_id=?
	`
	queryRow, err := m.DB.QueryRow(strSql, unionId)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&user.UserID,
		&wx_unoin_id,
		&user.Nickname,
		&user.Avatar,
		&user.Name,
		&mobileNullString,
		&user.IdCardNumber,
		&user.IdCardFront,
		&user.IdCardBack,
		&user.IsOrgManager,
		&user.IsOrgStaff,
		&user.IsCmtManager,
		&user.IsDisableOrgManager,
		&user.IsDisableOrgStaff,
		&user.IsDisableCmtManger,
		&user.CreateTime,
	)

	if err != nil {
		user = nil
		if err != conn.ErrNoRows {
			log.Warnf("get user from db failed. %s", err)
		}
		return
	}

	user.WxUnionId = wx_unoin_id.String
	user.Mobile = mobileNullString.String

	return
}

func (m *MallUser) GetUserByUnionIdOrMobile(unionId string, mobile string) (user *cidl.User, err error) {
	user = &cidl.User{}
	var wx_unoin_id sql.NullString
	var mobileNullString sql.NullString
	strSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM
			usr_user
		WHERE wx_union_id=? OR mobile=?
	`
	queryRow, err := m.DB.QueryRow(strSql, unionId, mobile)
	if err != nil {
		log.Warnf("get query row failed. %s", err)
		return
	}

	err = queryRow.Scan(
		&user.UserID,
		&wx_unoin_id,
		&user.Nickname,
		&user.Avatar,
		&user.Name,
		&mobileNullString,
		&user.IdCardNumber,
		&user.IdCardFront,
		&user.IdCardBack,
		&user.IsOrgManager,
		&user.IsOrgStaff,
		&user.IsCmtManager,
		&user.IsDisableOrgManager,
		&user.IsDisableOrgStaff,
		&user.IsDisableCmtManger,
		&user.CreateTime,
	)

	if err != nil {
		user = nil
		if err != conn.ErrNoRows {
			log.Warnf("get user from db failed. %s", err)
		}
		return
	}

	user.WxUnionId = wx_unoin_id.String
	user.Mobile = mobileNullString.String

	return
}

// 添加用户
func (m *MallUser) AddUser(user *cidl.User) (result sql.Result, err error) {
	var (
		wx_union_id sql.NullString
		mobile      sql.NullString
	)

	if user.WxUnionId != "" {
		wx_union_id = sql.NullString{
			String: user.WxUnionId,
			Valid:  true,
		}
	}

	if user.Mobile != "" {
		mobile = sql.NullString{
			String: user.Mobile,
			Valid:  true,
		}
	}

	strSql := `
		INSERT INTO
			usr_user (
				wx_union_id,
				nickname,
				avatar,
				name,
				mobile,
				id_card_number,
				id_card_front,
				id_card_back,
				is_org_manager,
				is_org_staff,
				is_cmt_manager,
				is_disable_org_manager,
				is_disable_org_staff,
				is_disable_cmt_manager
			)
		VALUES(
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?,
				?
		)
	`
	result, err = m.DB.Exec(strSql,
		wx_union_id,
		user.Nickname,
		user.Avatar,
		user.Name,
		mobile,
		user.IdCardNumber,
		user.IdCardFront,
		user.IdCardBack,
		user.IsOrgManager,
		user.IsOrgStaff,
		user.IsCmtManager,
		user.IsDisableOrgManager,
		user.IsDisableOrgStaff,
		user.IsDisableCmtManger,
	)
	if err != nil {
		log.Warnf("insert user failed. %s", err)
		return
	}

	return
}

// 更新用户基本信息
func (m *MallUser) UpdateUserBasic(userId string, nickname string, mobile string) (result sql.Result, err error) {
	strSql := `UPDATE usr_user SET nickname=?, mobile=? WHERE uid=?`
	result, err = m.DB.Exec(strSql, nickname, mobile, userId)
	return
}

// 组织管理员数目
func (m *MallUser) OrgManagerCount() (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_org_manager=1`
	err = m.DB.Get(&count, countSql)
	return
}

// 组织管理员用户列表
func (m *MallUser) OrgManagerList(page uint32, pageSize uint32, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_org_manager=1
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	rows, err := m.DB.Query(listSql, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}
		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String
		users = append(users, &user)
	}

	return
}

// 组织管理员数目
func (m *MallUser) OrgManagerSearchCount(search string) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_org_manager=1 AND (mobile LIKE ? OR name LIKE ?)`
	search = "%" + search + "%"
	err = m.DB.Get(&count, countSql, search, search)
	return
}

// 组织管理员用户列表
func (m *MallUser) OrgManagerSearchList(page uint32, pageSize uint32, search string, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_org_manager=1 AND (mobile LIKE ? OR name LIKE ?)
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	search = "%" + search + "%"
	rows, err := m.DB.Query(listSql, search, search, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}
		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String
		users = append(users, &user)
	}

	return
}

// 组织成员列表数目
func (m *MallUser) OrgStaffCount() (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_org_staff=1`
	err = m.DB.Get(&count, countSql)
	return
}

// 组织成员列表
func (m *MallUser) OrgStaffList(page uint32, pageSize uint32, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_org_staff=1
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	rows, err := m.DB.Query(listSql, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}
		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String
		users = append(users, &user)
	}

	return
}

// 组织成员列表数目
func (m *MallUser) OrgStaffSearchCount(search string) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_org_staff=1 AND (mobile LIKE ? OR name LIKE ?)`
	search = "%" + search + "%"
	err = m.DB.Get(&count, countSql, search, search)
	return
}

// 组织成员列表
func (m *MallUser) OrgStaffSearchList(page uint32, pageSize uint32, search string, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_org_staff=1 AND (mobile LIKE ? OR name LIKE ?)
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	search = "%" + search + "%"
	rows, err := m.DB.Query(listSql, search, search, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}

		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String

		users = append(users, &user)
	}

	return
}

// 社区合伙人列表
func (m *MallUser) CmtManagerCount() (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_cmt_manager=1`
	err = m.DB.Get(&count, countSql)
	return
}

// 社区合伙人列表
func (m *MallUser) CmtManagerList(page uint32, pageSize uint32, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_cmt_manager=1
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	rows, err := m.DB.Query(listSql, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}
		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String
		users = append(users, &user)
	}

	return
}

// 社区合伙人列表
func (m *MallUser) CmtManagerSearchCount(search string) (count uint32, err error) {
	countSql := `SELECT COUNT(*) FROM usr_user WHERE is_cmt_manager=1 AND (mobile LIKE ? OR name LIKE ?)`
	search = "%" + search + "%"
	err = m.DB.Get(&count, countSql, search, search)
	return
}

// 社区合伙人列表
func (m *MallUser) CmtManagerSearchList(page uint32, pageSize uint32, search string, idAsc bool) (users []*cidl.User, err error) {
	if page <= 0 || pageSize <= 0 {
		err = errors.New("page or pageSize should be greater than 0.")
		return
	}

	offset := (page - 1) * pageSize
	strOrderBy := "ASC"
	if false == idAsc {
		strOrderBy = "DESC"
	}
	listSql := `
		SELECT
			uid,
			wx_union_id,
			nickname,
			avatar,
			name,
			mobile,
			id_card_number,
			id_card_front,
			id_card_back,
			is_org_manager,
			is_org_staff,
			is_cmt_manager,
			is_disable_org_manager,
			is_disable_org_staff,
			is_disable_cmt_manager,
			create_time
		FROM usr_user
		WHERE is_cmt_manager=1 AND (mobile LIKE ? OR name LIKE ?)
		ORDER BY uid %s
		LIMIT ? OFFSET ?`

	listSql = fmt.Sprintf(listSql, strOrderBy)
	search = "%" + search + "%"
	rows, err := m.DB.Query(listSql, search, search, pageSize, offset)
	if err != nil {
		log.Warnf("query org manager list failed. %s", err)
		return
	}

	for rows.Next() {
		var user cidl.User
		var wx_unoin_id sql.NullString
		var mobile sql.NullString
		err = rows.Scan(
			&user.UserID,
			&wx_unoin_id,
			&user.Nickname,
			&user.Avatar,
			&user.Name,
			&mobile,
			&user.IdCardNumber,
			&user.IdCardFront,
			&user.IdCardBack,
			&user.IsOrgManager,
			&user.IsOrgStaff,
			&user.IsCmtManager,
			&user.IsDisableOrgManager,
			&user.IsDisableOrgStaff,
			&user.IsDisableCmtManger,
			&user.CreateTime,
		)
		if err != nil {
			log.Warnf("scan user failed. %s", err)
			return
		}
		user.WxUnionId = wx_unoin_id.String
		user.Mobile = mobile.String
	}

	return
}

// 编辑用户社区合伙人
func (m *MallUser) EditIsCmtManager(userId string, isCmtManager bool) (result sql.Result, err error) {
	editSql := `UPDATE usr_user SET is_cmt_manager=? WHERE uid=?`
	result, err = m.DB.Exec(editSql, isCmtManager, userId)
	return
}

// 更新用户禁用标志
func (m *MallUser) UpdateUserDisableState(userId string, userType cidl.UserType, isDisable bool) (result sql.Result, err error) {

	var strSql string
	if userType == cidl.OrgManager {
		strSql = `UPDATE usr_user SET is_disable_org_manager = ? WHERE uid=?`
	} else if userType == cidl.OrgStaff {
		
		strSql = `UPDATE usr_user SET is_disable_org_staff = ? WHERE uid=?`
	} else {
		strSql = `UPDATE usr_user SET is_disable_cmt_manager = ? WHERE uid=?`
	}
	result, err = m.DB.Exec(strSql, isDisable, userId)
	return
}


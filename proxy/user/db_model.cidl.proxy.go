package user

import "time"

// 用户
type User struct {
	UserID              string    `db:"uid"`
	WxUnionId           string    `db:"wx_union_id"`
	Nickname            string    `db:"nickname"`
	Avatar              string    `db:"avatar"`
	Name                string    `db:"name"`
	Mobile              string    `db:"mobile"`
	IdCardNumber        string    `db:"id_card_number"`
	IdCardFront         string    `db:"id_card_front"`
	IdCardBack          string    `db:"id_card_back"`
	IsOrgManager        bool      `db:"is_org_manager"`
	IsOrgStaff          bool      `db:"is_org_staff"`
	IsCmtManager        bool      `db:"is_cmt_manager"`
	IsDisableOrgManager bool      `db:"is_disable_org_manager"`
	IsDisableOrgStaff   bool      `db:"is_disable_org_staff"`
	IsDisableCmtManger  bool      `db:"is_disable_cmt_manger"`
	CreateTime          time.Time `db:"create_time"`
	UpdateTime          time.Time `db:"update_time"`
}

func NewUser() *User {
	return &User{}
}

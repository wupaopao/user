package cidl

type MetaApiInnerUserInfoByUserID struct {
}

var META_INNER_USER_INFO_BY_USER_ID = &MetaApiInnerUserInfoByUserID{}

func (m *MetaApiInnerUserInfoByUserID) GetMethod() string { return "GET" }
func (m *MetaApiInnerUserInfoByUserID) GetURL() string    { return "/inner/user/user/info/:user_id" }
func (m *MetaApiInnerUserInfoByUserID) GetName() string   { return "InnerUserInfoByUserID" }
func (m *MetaApiInnerUserInfoByUserID) GetType() string   { return "json" }

// 获取用户
type ApiInnerUserInfoByUserID struct {
	MetaApiInnerUserInfoByUserID
	Ack    *User
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerUserInfoByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserInfoByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerUserInfoByUserID) GetAsk() interface{}    { return nil }
func (m *ApiInnerUserInfoByUserID) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserInfoByUserID() ApiInnerUserInfoByUserID {
	return ApiInnerUserInfoByUserID{
		Ack: NewUser(),
	}
}

type AckInnerUserUserInfoByMobile struct {
	Exist bool  `db:"Exist"`
	User  *User `db:"User"`
}

func NewAckInnerUserUserInfoByMobile() *AckInnerUserUserInfoByMobile {
	return &AckInnerUserUserInfoByMobile{
		User: NewUser(),
	}
}

type MetaApiInnerUserUserInfoByMobile struct {
}

var META_INNER_USER_USER_INFO_BY_MOBILE = &MetaApiInnerUserUserInfoByMobile{}

func (m *MetaApiInnerUserUserInfoByMobile) GetMethod() string { return "GET" }
func (m *MetaApiInnerUserUserInfoByMobile) GetURL() string {
	return "/inner/user/user/info_by_mobile/:mobile"
}
func (m *MetaApiInnerUserUserInfoByMobile) GetName() string { return "InnerUserUserInfoByMobile" }
func (m *MetaApiInnerUserUserInfoByMobile) GetType() string { return "json" }

// 通过手机号获取用户
type ApiInnerUserUserInfoByMobile struct {
	MetaApiInnerUserUserInfoByMobile
	Ack    *AckInnerUserUserInfoByMobile
	Params struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiInnerUserUserInfoByMobile) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserInfoByMobile) GetParams() interface{} { return &m.Params }
func (m *ApiInnerUserUserInfoByMobile) GetAsk() interface{}    { return nil }
func (m *ApiInnerUserUserInfoByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserUserInfoByMobile() ApiInnerUserUserInfoByMobile {
	return ApiInnerUserUserInfoByMobile{
		Ack: NewAckInnerUserUserInfoByMobile(),
	}
}

type AskInnerUserCommunityManagerEditByUserID struct {
	IsCmtManager bool `binding:"required" db:"IsCmtManager"`
}

func NewAskInnerUserCommunityManagerEditByUserID() *AskInnerUserCommunityManagerEditByUserID {
	return &AskInnerUserCommunityManagerEditByUserID{}
}

type MetaApiInnerUserCommunityManagerEditByUserID struct {
}

var META_INNER_USER_COMMUNITY_MANAGER_EDIT_BY_USER_ID = &MetaApiInnerUserCommunityManagerEditByUserID{}

func (m *MetaApiInnerUserCommunityManagerEditByUserID) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserCommunityManagerEditByUserID) GetURL() string {
	return "/inner/user/user/community_manager/edit/:user_id"
}
func (m *MetaApiInnerUserCommunityManagerEditByUserID) GetName() string {
	return "InnerUserCommunityManagerEditByUserID"
}
func (m *MetaApiInnerUserCommunityManagerEditByUserID) GetType() string { return "json" }

// 设置用户为社区合伙人
type ApiInnerUserCommunityManagerEditByUserID struct {
	MetaApiInnerUserCommunityManagerEditByUserID
	Ask    *AskInnerUserCommunityManagerEditByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerUserCommunityManagerEditByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserCommunityManagerEditByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerUserCommunityManagerEditByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserCommunityManagerEditByUserID) GetAck() interface{}    { return nil }
func MakeApiInnerUserCommunityManagerEditByUserID() ApiInnerUserCommunityManagerEditByUserID {
	return ApiInnerUserCommunityManagerEditByUserID{
		Ask: NewAskInnerUserCommunityManagerEditByUserID(),
	}
}

type AskInnerUserUserOrgManagerAddOrUpdate struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	Nickname     string `binding:"required,lte=64" db:"Nickname"`
	IdCardNumber string `binding:"required,lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"required,lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"required,lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserOrgManagerAddOrUpdate() *AskInnerUserUserOrgManagerAddOrUpdate {
	return &AskInnerUserUserOrgManagerAddOrUpdate{}
}

type AckInnerUserUserOrgManagerAddOrUpdate struct {
	UserId string `db:"UserId"`
}

func NewAckInnerUserUserOrgManagerAddOrUpdate() *AckInnerUserUserOrgManagerAddOrUpdate {
	return &AckInnerUserUserOrgManagerAddOrUpdate{}
}

type MetaApiInnerUserUserOrgManagerAddOrUpdate struct {
}

var META_INNER_USER_USER_ORG_MANAGER_ADD_OR_UPDATE = &MetaApiInnerUserUserOrgManagerAddOrUpdate{}

func (m *MetaApiInnerUserUserOrgManagerAddOrUpdate) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserOrgManagerAddOrUpdate) GetURL() string {
	return "/inner/user/user/org_manager/add_or_update"
}
func (m *MetaApiInnerUserUserOrgManagerAddOrUpdate) GetName() string {
	return "InnerUserUserOrgManagerAddOrUpdate"
}
func (m *MetaApiInnerUserUserOrgManagerAddOrUpdate) GetType() string { return "json" }

// 添加组织管理员用户
type ApiInnerUserUserOrgManagerAddOrUpdate struct {
	MetaApiInnerUserUserOrgManagerAddOrUpdate
	Ask *AskInnerUserUserOrgManagerAddOrUpdate
	Ack *AckInnerUserUserOrgManagerAddOrUpdate
}

func (m *ApiInnerUserUserOrgManagerAddOrUpdate) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserOrgManagerAddOrUpdate) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserOrgManagerAddOrUpdate) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserOrgManagerAddOrUpdate) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserUserOrgManagerAddOrUpdate() ApiInnerUserUserOrgManagerAddOrUpdate {
	return ApiInnerUserUserOrgManagerAddOrUpdate{
		Ask: NewAskInnerUserUserOrgManagerAddOrUpdate(),
		Ack: NewAckInnerUserUserOrgManagerAddOrUpdate(),
	}
}

type AskInnerUserUserOrgManagerUpdate struct {
	UserId       string `binding:"required" db:"UserId"`
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	Nickname     string `binding:"required,lte=64" db:"Nickname"`
	IdCardNumber string `binding:"required,lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"required,lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"required,lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserOrgManagerUpdate() *AskInnerUserUserOrgManagerUpdate {
	return &AskInnerUserUserOrgManagerUpdate{}
}

type MetaApiInnerUserUserOrgManagerUpdate struct {
}

var META_INNER_USER_USER_ORG_MANAGER_UPDATE = &MetaApiInnerUserUserOrgManagerUpdate{}

func (m *MetaApiInnerUserUserOrgManagerUpdate) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserOrgManagerUpdate) GetURL() string {
	return "/inner/user/user/org_manager/update"
}
func (m *MetaApiInnerUserUserOrgManagerUpdate) GetName() string {
	return "InnerUserUserOrgManagerUpdate"
}
func (m *MetaApiInnerUserUserOrgManagerUpdate) GetType() string { return "json" }

// 更新组织管理员用户
type ApiInnerUserUserOrgManagerUpdate struct {
	MetaApiInnerUserUserOrgManagerUpdate
	Ask *AskInnerUserUserOrgManagerUpdate
}

func (m *ApiInnerUserUserOrgManagerUpdate) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserOrgManagerUpdate) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserOrgManagerUpdate) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserOrgManagerUpdate) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserOrgManagerUpdate() ApiInnerUserUserOrgManagerUpdate {
	return ApiInnerUserUserOrgManagerUpdate{
		Ask: NewAskInnerUserUserOrgManagerUpdate(),
	}
}

type AskInnerUserUserOrgManagerUnbind struct {
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserOrgManagerUnbind() *AskInnerUserUserOrgManagerUnbind {
	return &AskInnerUserUserOrgManagerUnbind{}
}

type MetaApiInnerUserUserOrgManagerUnbind struct {
}

var META_INNER_USER_USER_ORG_MANAGER_UNBIND = &MetaApiInnerUserUserOrgManagerUnbind{}

func (m *MetaApiInnerUserUserOrgManagerUnbind) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserOrgManagerUnbind) GetURL() string {
	return "/inner/user/user/org_manager/unbind"
}
func (m *MetaApiInnerUserUserOrgManagerUnbind) GetName() string {
	return "InnerUserUserOrgManagerUnbind"
}
func (m *MetaApiInnerUserUserOrgManagerUnbind) GetType() string { return "json" }

// 解绑组织管理员用户
type ApiInnerUserUserOrgManagerUnbind struct {
	MetaApiInnerUserUserOrgManagerUnbind
	Ask *AskInnerUserUserOrgManagerUnbind
}

func (m *ApiInnerUserUserOrgManagerUnbind) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserOrgManagerUnbind) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserOrgManagerUnbind) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserOrgManagerUnbind) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserOrgManagerUnbind() ApiInnerUserUserOrgManagerUnbind {
	return ApiInnerUserUserOrgManagerUnbind{
		Ask: NewAskInnerUserUserOrgManagerUnbind(),
	}
}

type AskInnerUserUserOrgStaffAddOrUpdate struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserUserOrgStaffAddOrUpdate() *AskInnerUserUserOrgStaffAddOrUpdate {
	return &AskInnerUserUserOrgStaffAddOrUpdate{}
}

type AckInnerUserUserOrgStaffAddOrUpdate struct {
	UserId string `db:"UserId"`
}

func NewAckInnerUserUserOrgStaffAddOrUpdate() *AckInnerUserUserOrgStaffAddOrUpdate {
	return &AckInnerUserUserOrgStaffAddOrUpdate{}
}

type MetaApiInnerUserUserOrgStaffAddOrUpdate struct {
}

var META_INNER_USER_USER_ORG_STAFF_ADD_OR_UPDATE = &MetaApiInnerUserUserOrgStaffAddOrUpdate{}

func (m *MetaApiInnerUserUserOrgStaffAddOrUpdate) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserOrgStaffAddOrUpdate) GetURL() string {
	return "/inner/user/user/org_staff/add_or_update"
}
func (m *MetaApiInnerUserUserOrgStaffAddOrUpdate) GetName() string {
	return "InnerUserUserOrgStaffAddOrUpdate"
}
func (m *MetaApiInnerUserUserOrgStaffAddOrUpdate) GetType() string { return "json" }

// 添加组织成员
type ApiInnerUserUserOrgStaffAddOrUpdate struct {
	MetaApiInnerUserUserOrgStaffAddOrUpdate
	Ask *AskInnerUserUserOrgStaffAddOrUpdate
	Ack *AckInnerUserUserOrgStaffAddOrUpdate
}

func (m *ApiInnerUserUserOrgStaffAddOrUpdate) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserOrgStaffAddOrUpdate) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserOrgStaffAddOrUpdate) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserOrgStaffAddOrUpdate) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserUserOrgStaffAddOrUpdate() ApiInnerUserUserOrgStaffAddOrUpdate {
	return ApiInnerUserUserOrgStaffAddOrUpdate{
		Ask: NewAskInnerUserUserOrgStaffAddOrUpdate(),
		Ack: NewAckInnerUserUserOrgStaffAddOrUpdate(),
	}
}

type AskInnerUserUserOrgStaffUpdateByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserUserOrgStaffUpdateByUserID() *AskInnerUserUserOrgStaffUpdateByUserID {
	return &AskInnerUserUserOrgStaffUpdateByUserID{}
}

type MetaApiInnerUserUserOrgStaffUpdateByUserID struct {
}

var META_INNER_USER_USER_ORG_STAFF_UPDATE_BY_USER_ID = &MetaApiInnerUserUserOrgStaffUpdateByUserID{}

func (m *MetaApiInnerUserUserOrgStaffUpdateByUserID) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserOrgStaffUpdateByUserID) GetURL() string {
	return "/inner/user/user/org_staff/update/:user_id"
}
func (m *MetaApiInnerUserUserOrgStaffUpdateByUserID) GetName() string {
	return "InnerUserUserOrgStaffUpdateByUserID"
}
func (m *MetaApiInnerUserUserOrgStaffUpdateByUserID) GetType() string { return "json" }

// 更新组织成员
type ApiInnerUserUserOrgStaffUpdateByUserID struct {
	MetaApiInnerUserUserOrgStaffUpdateByUserID
	Ask    *AskInnerUserUserOrgStaffUpdateByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerUserUserOrgStaffUpdateByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserOrgStaffUpdateByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerUserUserOrgStaffUpdateByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserOrgStaffUpdateByUserID) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserOrgStaffUpdateByUserID() ApiInnerUserUserOrgStaffUpdateByUserID {
	return ApiInnerUserUserOrgStaffUpdateByUserID{
		Ask: NewAskInnerUserUserOrgStaffUpdateByUserID(),
	}
}

type AskInnerUserUserCmtManagerUpdateByUserID struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	IdCardNumber string `binding:"lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserCmtManagerUpdateByUserID() *AskInnerUserUserCmtManagerUpdateByUserID {
	return &AskInnerUserUserCmtManagerUpdateByUserID{}
}

type MetaApiInnerUserUserCmtManagerUpdateByUserID struct {
}

var META_INNER_USER_USER_CMT_MANAGER_UPDATE_BY_USER_ID = &MetaApiInnerUserUserCmtManagerUpdateByUserID{}

func (m *MetaApiInnerUserUserCmtManagerUpdateByUserID) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserCmtManagerUpdateByUserID) GetURL() string {
	return "/inner/user/user/cmt_manager/update/:user_id"
}
func (m *MetaApiInnerUserUserCmtManagerUpdateByUserID) GetName() string {
	return "InnerUserUserCmtManagerUpdateByUserID"
}
func (m *MetaApiInnerUserUserCmtManagerUpdateByUserID) GetType() string { return "json" }

// 更新社区合伙人信息
type ApiInnerUserUserCmtManagerUpdateByUserID struct {
	MetaApiInnerUserUserCmtManagerUpdateByUserID
	Ask    *AskInnerUserUserCmtManagerUpdateByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiInnerUserUserCmtManagerUpdateByUserID) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserCmtManagerUpdateByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiInnerUserUserCmtManagerUpdateByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserCmtManagerUpdateByUserID) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserCmtManagerUpdateByUserID() ApiInnerUserUserCmtManagerUpdateByUserID {
	return ApiInnerUserUserCmtManagerUpdateByUserID{
		Ask: NewAskInnerUserUserCmtManagerUpdateByUserID(),
	}
}

type AskInnerUserUserCmtManagerChange struct {
	NewManagerUid string `binding:"required" db:"NewManagerUid"`
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserCmtManagerChange() *AskInnerUserUserCmtManagerChange {
	return &AskInnerUserUserCmtManagerChange{}
}

type MetaApiInnerUserUserCmtManagerChange struct {
}

var META_INNER_USER_USER_CMT_MANAGER_CHANGE = &MetaApiInnerUserUserCmtManagerChange{}

func (m *MetaApiInnerUserUserCmtManagerChange) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserCmtManagerChange) GetURL() string {
	return "/inner/user/user/cmt_manager/change"
}
func (m *MetaApiInnerUserUserCmtManagerChange) GetName() string {
	return "InnerUserUserCmtManagerChange"
}
func (m *MetaApiInnerUserUserCmtManagerChange) GetType() string { return "json" }

// 更改社区合伙人
type ApiInnerUserUserCmtManagerChange struct {
	MetaApiInnerUserUserCmtManagerChange
	Ask *AskInnerUserUserCmtManagerChange
}

func (m *ApiInnerUserUserCmtManagerChange) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserCmtManagerChange) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserCmtManagerChange) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserCmtManagerChange) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserCmtManagerChange() ApiInnerUserUserCmtManagerChange {
	return ApiInnerUserUserCmtManagerChange{
		Ask: NewAskInnerUserUserCmtManagerChange(),
	}
}

type AskInnerUserUserCmtManagerGetOrAdd struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	IdCardNumber string `binding:"lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserCmtManagerGetOrAdd() *AskInnerUserUserCmtManagerGetOrAdd {
	return &AskInnerUserUserCmtManagerGetOrAdd{}
}

type AckInnerUserUserCmtManagerGetOrAdd struct {
	IsNew bool  `db:"IsNew"`
	User  *User `db:"User"`
}

func NewAckInnerUserUserCmtManagerGetOrAdd() *AckInnerUserUserCmtManagerGetOrAdd {
	return &AckInnerUserUserCmtManagerGetOrAdd{
		User: NewUser(),
	}
}

type MetaApiInnerUserUserCmtManagerGetOrAdd struct {
}

var META_INNER_USER_USER_CMT_MANAGER_GET_OR_ADD = &MetaApiInnerUserUserCmtManagerGetOrAdd{}

func (m *MetaApiInnerUserUserCmtManagerGetOrAdd) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserCmtManagerGetOrAdd) GetURL() string {
	return "/inner/user/user/cmt_manager/get_or_add"
}
func (m *MetaApiInnerUserUserCmtManagerGetOrAdd) GetName() string {
	return "InnerUserUserCmtManagerGetOrAdd"
}
func (m *MetaApiInnerUserUserCmtManagerGetOrAdd) GetType() string { return "json" }

// 添加社区合伙人
type ApiInnerUserUserCmtManagerGetOrAdd struct {
	MetaApiInnerUserUserCmtManagerGetOrAdd
	Ask *AskInnerUserUserCmtManagerGetOrAdd
	Ack *AckInnerUserUserCmtManagerGetOrAdd
}

func (m *ApiInnerUserUserCmtManagerGetOrAdd) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserCmtManagerGetOrAdd) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserCmtManagerGetOrAdd) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserCmtManagerGetOrAdd) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserUserCmtManagerGetOrAdd() ApiInnerUserUserCmtManagerGetOrAdd {
	return ApiInnerUserUserCmtManagerGetOrAdd{
		Ask: NewAskInnerUserUserCmtManagerGetOrAdd(),
		Ack: NewAckInnerUserUserCmtManagerGetOrAdd(),
	}
}

type AskInnerUserUserCmtManagerUnbind struct {
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserCmtManagerUnbind() *AskInnerUserUserCmtManagerUnbind {
	return &AskInnerUserUserCmtManagerUnbind{}
}

type MetaApiInnerUserUserCmtManagerUnbind struct {
}

var META_INNER_USER_USER_CMT_MANAGER_UNBIND = &MetaApiInnerUserUserCmtManagerUnbind{}

func (m *MetaApiInnerUserUserCmtManagerUnbind) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserUserCmtManagerUnbind) GetURL() string {
	return "/inner/user/user/cmt_manager/unbind"
}
func (m *MetaApiInnerUserUserCmtManagerUnbind) GetName() string {
	return "InnerUserUserCmtManagerUnbind"
}
func (m *MetaApiInnerUserUserCmtManagerUnbind) GetType() string { return "json" }

// 解绑社区合伙人
type ApiInnerUserUserCmtManagerUnbind struct {
	MetaApiInnerUserUserCmtManagerUnbind
	Ask *AskInnerUserUserCmtManagerUnbind
}

func (m *ApiInnerUserUserCmtManagerUnbind) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserUserCmtManagerUnbind) GetParams() interface{} { return nil }
func (m *ApiInnerUserUserCmtManagerUnbind) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserUserCmtManagerUnbind) GetAck() interface{}    { return nil }
func MakeApiInnerUserUserCmtManagerUnbind() ApiInnerUserUserCmtManagerUnbind {
	return ApiInnerUserUserCmtManagerUnbind{
		Ask: NewAskInnerUserUserCmtManagerUnbind(),
	}
}

type AskInnerUserWxXcxCmtManagerGetOrAdd struct {
	UserId string `binding:"required" db:"UserId"`
	Token  string `binding:"required" db:"Token"`
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserWxXcxCmtManagerGetOrAdd() *AskInnerUserWxXcxCmtManagerGetOrAdd {
	return &AskInnerUserWxXcxCmtManagerGetOrAdd{}
}

type AckInnerUserWxXcxCmtManagerGetOrAdd struct {
	IsNew bool  `db:"IsNew"`
	User  *User `db:"User"`
}

func NewAckInnerUserWxXcxCmtManagerGetOrAdd() *AckInnerUserWxXcxCmtManagerGetOrAdd {
	return &AckInnerUserWxXcxCmtManagerGetOrAdd{
		User: NewUser(),
	}
}

type MetaApiInnerUserWxXcxCmtManagerGetOrAdd struct {
}

var META_INNER_USER_WX_XCX_CMT_MANAGER_GET_OR_ADD = &MetaApiInnerUserWxXcxCmtManagerGetOrAdd{}

func (m *MetaApiInnerUserWxXcxCmtManagerGetOrAdd) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserWxXcxCmtManagerGetOrAdd) GetURL() string {
	return "/inner/user/wx_xcx/cmt_manager/get_or_add"
}
func (m *MetaApiInnerUserWxXcxCmtManagerGetOrAdd) GetName() string {
	return "InnerUserWxXcxCmtManagerGetOrAdd"
}
func (m *MetaApiInnerUserWxXcxCmtManagerGetOrAdd) GetType() string { return "json" }

// 微信小程序专用内部接口
// 添加社区合伙人
type ApiInnerUserWxXcxCmtManagerGetOrAdd struct {
	MetaApiInnerUserWxXcxCmtManagerGetOrAdd
	Ask *AskInnerUserWxXcxCmtManagerGetOrAdd
	Ack *AckInnerUserWxXcxCmtManagerGetOrAdd
}

func (m *ApiInnerUserWxXcxCmtManagerGetOrAdd) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserWxXcxCmtManagerGetOrAdd) GetParams() interface{} { return nil }
func (m *ApiInnerUserWxXcxCmtManagerGetOrAdd) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserWxXcxCmtManagerGetOrAdd) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserWxXcxCmtManagerGetOrAdd() ApiInnerUserWxXcxCmtManagerGetOrAdd {
	return ApiInnerUserWxXcxCmtManagerGetOrAdd{
		Ask: NewAskInnerUserWxXcxCmtManagerGetOrAdd(),
		Ack: NewAckInnerUserWxXcxCmtManagerGetOrAdd(),
	}
}

type AskInnerUserWxXcxRefreshToken struct {
	UserId string `binding:"required" db:"UserId"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserWxXcxRefreshToken() *AskInnerUserWxXcxRefreshToken {
	return &AskInnerUserWxXcxRefreshToken{}
}

type MetaApiInnerUserWxXcxRefreshToken struct {
}

var META_INNER_USER_WX_XCX_REFRESH_TOKEN = &MetaApiInnerUserWxXcxRefreshToken{}

func (m *MetaApiInnerUserWxXcxRefreshToken) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserWxXcxRefreshToken) GetURL() string    { return "/inner/user/wx_xcx/refresh_token" }
func (m *MetaApiInnerUserWxXcxRefreshToken) GetName() string   { return "InnerUserWxXcxRefreshToken" }
func (m *MetaApiInnerUserWxXcxRefreshToken) GetType() string   { return "json" }

// 刷新token
type ApiInnerUserWxXcxRefreshToken struct {
	MetaApiInnerUserWxXcxRefreshToken
	Ask *AskInnerUserWxXcxRefreshToken
}

func (m *ApiInnerUserWxXcxRefreshToken) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserWxXcxRefreshToken) GetParams() interface{} { return nil }
func (m *ApiInnerUserWxXcxRefreshToken) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserWxXcxRefreshToken) GetAck() interface{}    { return nil }
func MakeApiInnerUserWxXcxRefreshToken() ApiInnerUserWxXcxRefreshToken {
	return ApiInnerUserWxXcxRefreshToken{
		Ask: NewAskInnerUserWxXcxRefreshToken(),
	}
}

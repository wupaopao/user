package cidl

type MetaApiOrgUserInfoByUserId struct {
}

var META_ORG_USER_INFO_BY_USER_ID = &MetaApiOrgUserInfoByUserId{}

func (m *MetaApiOrgUserInfoByUserId) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserInfoByUserId) GetURL() string    { return "/user/org/user/info/:user_id" }
func (m *MetaApiOrgUserInfoByUserId) GetName() string   { return "OrgUserInfoByUserId" }
func (m *MetaApiOrgUserInfoByUserId) GetType() string   { return "json" }

// 用户信息
type ApiOrgUserInfoByUserId struct {
	MetaApiOrgUserInfoByUserId
	Ack    *User
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiOrgUserInfoByUserId) GetQuery() interface{}  { return nil }
func (m *ApiOrgUserInfoByUserId) GetParams() interface{} { return &m.Params }
func (m *ApiOrgUserInfoByUserId) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserInfoByUserId) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserInfoByUserId() ApiOrgUserInfoByUserId {
	return ApiOrgUserInfoByUserId{
		Ack: NewUser(),
	}
}

type AckOrgUserOrganizationManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserOrganizationManagerList() *AckOrgUserOrganizationManagerList {
	return &AckOrgUserOrganizationManagerList{
		List: make([]*User, 0),
	}
}

type MetaApiOrgUserOrganizationManagerList struct {
}

var META_ORG_USER_ORGANIZATION_MANAGER_LIST = &MetaApiOrgUserOrganizationManagerList{}

func (m *MetaApiOrgUserOrganizationManagerList) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserOrganizationManagerList) GetURL() string {
	return "/user/org/user/organization_manager/list"
}
func (m *MetaApiOrgUserOrganizationManagerList) GetName() string {
	return "OrgUserOrganizationManagerList"
}
func (m *MetaApiOrgUserOrganizationManagerList) GetType() string { return "json" }

// 城市合伙人用户
type ApiOrgUserOrganizationManagerList struct {
	MetaApiOrgUserOrganizationManagerList
	Ack   *AckOrgUserOrganizationManagerList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiOrgUserOrganizationManagerList) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgUserOrganizationManagerList) GetParams() interface{} { return nil }
func (m *ApiOrgUserOrganizationManagerList) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserOrganizationManagerList) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserOrganizationManagerList() ApiOrgUserOrganizationManagerList {
	return ApiOrgUserOrganizationManagerList{
		Ack: NewAckOrgUserOrganizationManagerList(),
	}
}

type AckOrgUserCommunityManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserCommunityManagerList() *AckOrgUserCommunityManagerList {
	return &AckOrgUserCommunityManagerList{
		List: make([]*User, 0),
	}
}

type MetaApiOrgUserCommunityManagerList struct {
}

var META_ORG_USER_COMMUNITY_MANAGER_LIST = &MetaApiOrgUserCommunityManagerList{}

func (m *MetaApiOrgUserCommunityManagerList) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserCommunityManagerList) GetURL() string {
	return "/user/org/user/community_manager/list"
}
func (m *MetaApiOrgUserCommunityManagerList) GetName() string { return "OrgUserCommunityManagerList" }
func (m *MetaApiOrgUserCommunityManagerList) GetType() string { return "json" }

// 社区合伙人
type ApiOrgUserCommunityManagerList struct {
	MetaApiOrgUserCommunityManagerList
	Ack   *AckOrgUserCommunityManagerList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiOrgUserCommunityManagerList) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgUserCommunityManagerList) GetParams() interface{} { return nil }
func (m *ApiOrgUserCommunityManagerList) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserCommunityManagerList) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserCommunityManagerList() ApiOrgUserCommunityManagerList {
	return ApiOrgUserCommunityManagerList{
		Ack: NewAckOrgUserCommunityManagerList(),
	}
}

type AckOrgUserOrganizationStaffList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserOrganizationStaffList() *AckOrgUserOrganizationStaffList {
	return &AckOrgUserOrganizationStaffList{
		List: make([]*User, 0),
	}
}

type MetaApiOrgUserOrganizationStaffList struct {
}

var META_ORG_USER_ORGANIZATION_STAFF_LIST = &MetaApiOrgUserOrganizationStaffList{}

func (m *MetaApiOrgUserOrganizationStaffList) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserOrganizationStaffList) GetURL() string {
	return "/user/org/user/organization_staff/list"
}
func (m *MetaApiOrgUserOrganizationStaffList) GetName() string { return "OrgUserOrganizationStaffList" }
func (m *MetaApiOrgUserOrganizationStaffList) GetType() string { return "json" }

// 组织成员
type ApiOrgUserOrganizationStaffList struct {
	MetaApiOrgUserOrganizationStaffList
	Ack   *AckOrgUserOrganizationStaffList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiOrgUserOrganizationStaffList) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgUserOrganizationStaffList) GetParams() interface{} { return nil }
func (m *ApiOrgUserOrganizationStaffList) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserOrganizationStaffList) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserOrganizationStaffList() ApiOrgUserOrganizationStaffList {
	return ApiOrgUserOrganizationStaffList{
		Ack: NewAckOrgUserOrganizationStaffList(),
	}
}

type AskOrgUserIDCardPicToken struct {
	FileNames []string `db:"FileNames"`
}

func NewAskOrgUserIDCardPicToken() *AskOrgUserIDCardPicToken {
	return &AskOrgUserIDCardPicToken{
		FileNames: make([]string, 0),
	}
}

type AckOrgUserIDCardPicToken struct {
	Tokens []*AckPicToken `db:"Tokens"`
}

func NewAckOrgUserIDCardPicToken() *AckOrgUserIDCardPicToken {
	return &AckOrgUserIDCardPicToken{
		Tokens: make([]*AckPicToken, 0),
	}
}

type MetaApiOrgUserIDCardPicToken struct {
}

var META_ORG_USER_ID_CARD_PIC_TOKEN = &MetaApiOrgUserIDCardPicToken{}

func (m *MetaApiOrgUserIDCardPicToken) GetMethod() string { return "POST" }
func (m *MetaApiOrgUserIDCardPicToken) GetURL() string    { return "/user/org/user/id_card_pic_token" }
func (m *MetaApiOrgUserIDCardPicToken) GetName() string   { return "OrgUserIDCardPicToken" }
func (m *MetaApiOrgUserIDCardPicToken) GetType() string   { return "json" }

// 身份证token
type ApiOrgUserIDCardPicToken struct {
	MetaApiOrgUserIDCardPicToken
	Ask *AskOrgUserIDCardPicToken
	Ack *AckOrgUserIDCardPicToken
}

func (m *ApiOrgUserIDCardPicToken) GetQuery() interface{}  { return nil }
func (m *ApiOrgUserIDCardPicToken) GetParams() interface{} { return nil }
func (m *ApiOrgUserIDCardPicToken) GetAsk() interface{}    { return m.Ask }
func (m *ApiOrgUserIDCardPicToken) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserIDCardPicToken() ApiOrgUserIDCardPicToken {
	return ApiOrgUserIDCardPicToken{
		Ask: NewAskOrgUserIDCardPicToken(),
		Ack: NewAckOrgUserIDCardPicToken(),
	}
}

type MetaApiOrgUserAccessIDCardPic struct {
}

var META_ORG_USER_ACCESS_ID_CARD_PIC = &MetaApiOrgUserAccessIDCardPic{}

func (m *MetaApiOrgUserAccessIDCardPic) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserAccessIDCardPic) GetURL() string    { return "/user/org/user/access_id_card_pic" }
func (m *MetaApiOrgUserAccessIDCardPic) GetName() string   { return "OrgUserAccessIDCardPic" }
func (m *MetaApiOrgUserAccessIDCardPic) GetType() string   { return "json" }

// 访问身份证图片
type ApiOrgUserAccessIDCardPic struct {
	MetaApiOrgUserAccessIDCardPic
	Query struct {
		Uri string `form:"uri" binding:"required" db:"Uri"`
	}
}

func (m *ApiOrgUserAccessIDCardPic) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgUserAccessIDCardPic) GetParams() interface{} { return nil }
func (m *ApiOrgUserAccessIDCardPic) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserAccessIDCardPic) GetAck() interface{}    { return nil }
func MakeApiOrgUserAccessIDCardPic() ApiOrgUserAccessIDCardPic {
	return ApiOrgUserAccessIDCardPic{}
}

type AckOrgUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckOrgUserCheckBindMobileByMobile() *AckOrgUserCheckBindMobileByMobile {
	return &AckOrgUserCheckBindMobileByMobile{}
}

type MetaApiOrgUserCheckBindMobileByMobile struct {
}

var META_ORG_USER_CHECK_BIND_MOBILE_BY_MOBILE = &MetaApiOrgUserCheckBindMobileByMobile{}

func (m *MetaApiOrgUserCheckBindMobileByMobile) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserCheckBindMobileByMobile) GetURL() string {
	return "/user/org/user/check_bind_mobile/:mobile"
}
func (m *MetaApiOrgUserCheckBindMobileByMobile) GetName() string {
	return "OrgUserCheckBindMobileByMobile"
}
func (m *MetaApiOrgUserCheckBindMobileByMobile) GetType() string { return "json" }

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
type ApiOrgUserCheckBindMobileByMobile struct {
	MetaApiOrgUserCheckBindMobileByMobile
	Ack    *AckOrgUserCheckBindMobileByMobile
	Params struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiOrgUserCheckBindMobileByMobile) GetQuery() interface{}  { return nil }
func (m *ApiOrgUserCheckBindMobileByMobile) GetParams() interface{} { return &m.Params }
func (m *ApiOrgUserCheckBindMobileByMobile) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserCheckBindMobileByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiOrgUserCheckBindMobileByMobile() ApiOrgUserCheckBindMobileByMobile {
	return ApiOrgUserCheckBindMobileByMobile{
		Ack: NewAckOrgUserCheckBindMobileByMobile(),
	}
}

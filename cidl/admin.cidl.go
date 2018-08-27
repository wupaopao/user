package cidl

type MetaApiAdminUserInfoByUserId struct {
}

var META_ADMIN_USER_INFO_BY_USER_ID = &MetaApiAdminUserInfoByUserId{}

func (m *MetaApiAdminUserInfoByUserId) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserInfoByUserId) GetURL() string    { return "/user/admin/user/info/:user_id" }
func (m *MetaApiAdminUserInfoByUserId) GetName() string   { return "AdminUserInfoByUserId" }
func (m *MetaApiAdminUserInfoByUserId) GetType() string   { return "json" }

// 用户信息
type ApiAdminUserInfoByUserId struct {
	MetaApiAdminUserInfoByUserId
	Ack    *User
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiAdminUserInfoByUserId) GetQuery() interface{}  { return nil }
func (m *ApiAdminUserInfoByUserId) GetParams() interface{} { return &m.Params }
func (m *ApiAdminUserInfoByUserId) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserInfoByUserId) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserInfoByUserId() ApiAdminUserInfoByUserId {
	return ApiAdminUserInfoByUserId{
		Ack: NewUser(),
	}
}

type MetaApiAdminUserInfoByMobile struct {
}

var META_ADMIN_USER_INFO_BY_MOBILE = &MetaApiAdminUserInfoByMobile{}

func (m *MetaApiAdminUserInfoByMobile) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserInfoByMobile) GetURL() string    { return "/user/admin/user/info_by_mobile" }
func (m *MetaApiAdminUserInfoByMobile) GetName() string   { return "AdminUserInfoByMobile" }
func (m *MetaApiAdminUserInfoByMobile) GetType() string   { return "json" }

// 通过手机号获取用户信息
type ApiAdminUserInfoByMobile struct {
	MetaApiAdminUserInfoByMobile
	Ack   *User
	Query struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiAdminUserInfoByMobile) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminUserInfoByMobile) GetParams() interface{} { return nil }
func (m *ApiAdminUserInfoByMobile) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserInfoByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserInfoByMobile() ApiAdminUserInfoByMobile {
	return ApiAdminUserInfoByMobile{
		Ack: NewUser(),
	}
}

type AckAdminUserOrganizationManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserOrganizationManagerList() *AckAdminUserOrganizationManagerList {
	return &AckAdminUserOrganizationManagerList{
		List: make([]*User, 0),
	}
}

type MetaApiAdminUserOrganizationManagerList struct {
}

var META_ADMIN_USER_ORGANIZATION_MANAGER_LIST = &MetaApiAdminUserOrganizationManagerList{}

func (m *MetaApiAdminUserOrganizationManagerList) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserOrganizationManagerList) GetURL() string {
	return "/user/admin/user/organization_manager/list"
}
func (m *MetaApiAdminUserOrganizationManagerList) GetName() string {
	return "AdminUserOrganizationManagerList"
}
func (m *MetaApiAdminUserOrganizationManagerList) GetType() string { return "json" }

// 城市合伙人用户列表
type ApiAdminUserOrganizationManagerList struct {
	MetaApiAdminUserOrganizationManagerList
	Ack   *AckAdminUserOrganizationManagerList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiAdminUserOrganizationManagerList) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminUserOrganizationManagerList) GetParams() interface{} { return nil }
func (m *ApiAdminUserOrganizationManagerList) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserOrganizationManagerList) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserOrganizationManagerList() ApiAdminUserOrganizationManagerList {
	return ApiAdminUserOrganizationManagerList{
		Ack: NewAckAdminUserOrganizationManagerList(),
	}
}

type AckAdminUserCommunityManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserCommunityManagerList() *AckAdminUserCommunityManagerList {
	return &AckAdminUserCommunityManagerList{
		List: make([]*User, 0),
	}
}

type MetaApiAdminUserCommunityManagerList struct {
}

var META_ADMIN_USER_COMMUNITY_MANAGER_LIST = &MetaApiAdminUserCommunityManagerList{}

func (m *MetaApiAdminUserCommunityManagerList) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserCommunityManagerList) GetURL() string {
	return "/user/admin/user/community_manager/list"
}
func (m *MetaApiAdminUserCommunityManagerList) GetName() string {
	return "AdminUserCommunityManagerList"
}
func (m *MetaApiAdminUserCommunityManagerList) GetType() string { return "json" }

// 社区合伙人列表
type ApiAdminUserCommunityManagerList struct {
	MetaApiAdminUserCommunityManagerList
	Ack   *AckAdminUserCommunityManagerList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiAdminUserCommunityManagerList) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminUserCommunityManagerList) GetParams() interface{} { return nil }
func (m *ApiAdminUserCommunityManagerList) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserCommunityManagerList) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserCommunityManagerList() ApiAdminUserCommunityManagerList {
	return ApiAdminUserCommunityManagerList{
		Ack: NewAckAdminUserCommunityManagerList(),
	}
}

type AckAdminUserOrganizationStaffList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserOrganizationStaffList() *AckAdminUserOrganizationStaffList {
	return &AckAdminUserOrganizationStaffList{
		List: make([]*User, 0),
	}
}

type MetaApiAdminUserOrganizationStaffList struct {
}

var META_ADMIN_USER_ORGANIZATION_STAFF_LIST = &MetaApiAdminUserOrganizationStaffList{}

func (m *MetaApiAdminUserOrganizationStaffList) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserOrganizationStaffList) GetURL() string {
	return "/user/admin/user/organization_staff/list"
}
func (m *MetaApiAdminUserOrganizationStaffList) GetName() string {
	return "AdminUserOrganizationStaffList"
}
func (m *MetaApiAdminUserOrganizationStaffList) GetType() string { return "json" }

// 组织成员列表
type ApiAdminUserOrganizationStaffList struct {
	MetaApiAdminUserOrganizationStaffList
	Ack   *AckAdminUserOrganizationStaffList
	Query struct {
		Page     uint32 `form:"page" binding:"required,gt=0" db:"Page"`
		PageSize uint32 `form:"page_size" binding:"required,gt=0,lt=50" db:"PageSize"`
		Search   string `form:"search" db:"Search"`
	}
}

func (m *ApiAdminUserOrganizationStaffList) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminUserOrganizationStaffList) GetParams() interface{} { return nil }
func (m *ApiAdminUserOrganizationStaffList) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserOrganizationStaffList) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserOrganizationStaffList() ApiAdminUserOrganizationStaffList {
	return ApiAdminUserOrganizationStaffList{
		Ack: NewAckAdminUserOrganizationStaffList(),
	}
}

// 身份证图片Token获取
type AckPicToken struct {
	OriginalFileName string `db:"OriginalFileName"`
	Token            string `db:"Token"`
	Key              string `db:"Key"`
	StoreUrl         string `db:"StoreUrl"`
	AccessUrl        string `db:"AccessUrl"`
}

func NewAckPicToken() *AckPicToken {
	return &AckPicToken{}
}

type AskAdminUserIDCardPicToken struct {
	FileNames []string `db:"FileNames"`
}

func NewAskAdminUserIDCardPicToken() *AskAdminUserIDCardPicToken {
	return &AskAdminUserIDCardPicToken{
		FileNames: make([]string, 0),
	}
}

type AckAdminUserIDCardPicToken struct {
	Tokens []*AckPicToken `db:"Tokens"`
}

func NewAckAdminUserIDCardPicToken() *AckAdminUserIDCardPicToken {
	return &AckAdminUserIDCardPicToken{
		Tokens: make([]*AckPicToken, 0),
	}
}

type MetaApiAdminUserIDCardPicToken struct {
}

var META_ADMIN_USER_ID_CARD_PIC_TOKEN = &MetaApiAdminUserIDCardPicToken{}

func (m *MetaApiAdminUserIDCardPicToken) GetMethod() string { return "POST" }
func (m *MetaApiAdminUserIDCardPicToken) GetURL() string    { return "/user/admin/user/id_card_pic_token" }
func (m *MetaApiAdminUserIDCardPicToken) GetName() string   { return "AdminUserIDCardPicToken" }
func (m *MetaApiAdminUserIDCardPicToken) GetType() string   { return "json" }

// 身份证token
type ApiAdminUserIDCardPicToken struct {
	MetaApiAdminUserIDCardPicToken
	Ask *AskAdminUserIDCardPicToken
	Ack *AckAdminUserIDCardPicToken
}

func (m *ApiAdminUserIDCardPicToken) GetQuery() interface{}  { return nil }
func (m *ApiAdminUserIDCardPicToken) GetParams() interface{} { return nil }
func (m *ApiAdminUserIDCardPicToken) GetAsk() interface{}    { return m.Ask }
func (m *ApiAdminUserIDCardPicToken) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserIDCardPicToken() ApiAdminUserIDCardPicToken {
	return ApiAdminUserIDCardPicToken{
		Ask: NewAskAdminUserIDCardPicToken(),
		Ack: NewAckAdminUserIDCardPicToken(),
	}
}

type MetaApiAdminUserAccessIDCardPic struct {
}

var META_ADMIN_USER_ACCESS_ID_CARD_PIC = &MetaApiAdminUserAccessIDCardPic{}

func (m *MetaApiAdminUserAccessIDCardPic) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserAccessIDCardPic) GetURL() string {
	return "/user/admin/user/access_id_card_pic"
}
func (m *MetaApiAdminUserAccessIDCardPic) GetName() string { return "AdminUserAccessIDCardPic" }
func (m *MetaApiAdminUserAccessIDCardPic) GetType() string { return "json" }

// 访问身份证图片
type ApiAdminUserAccessIDCardPic struct {
	MetaApiAdminUserAccessIDCardPic
	Query struct {
		Uri string `form:"uri" binding:"required" db:"Uri"`
	}
}

func (m *ApiAdminUserAccessIDCardPic) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminUserAccessIDCardPic) GetParams() interface{} { return nil }
func (m *ApiAdminUserAccessIDCardPic) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserAccessIDCardPic) GetAck() interface{}    { return nil }
func MakeApiAdminUserAccessIDCardPic() ApiAdminUserAccessIDCardPic {
	return ApiAdminUserAccessIDCardPic{}
}

type AckAdminUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckAdminUserCheckBindMobileByMobile() *AckAdminUserCheckBindMobileByMobile {
	return &AckAdminUserCheckBindMobileByMobile{}
}

type MetaApiAdminUserCheckBindMobileByMobile struct {
}

var META_ADMIN_USER_CHECK_BIND_MOBILE_BY_MOBILE = &MetaApiAdminUserCheckBindMobileByMobile{}

func (m *MetaApiAdminUserCheckBindMobileByMobile) GetMethod() string { return "GET" }
func (m *MetaApiAdminUserCheckBindMobileByMobile) GetURL() string {
	return "/user/admin/user/check_bind_mobile/:mobile"
}
func (m *MetaApiAdminUserCheckBindMobileByMobile) GetName() string {
	return "AdminUserCheckBindMobileByMobile"
}
func (m *MetaApiAdminUserCheckBindMobileByMobile) GetType() string { return "json" }

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
type ApiAdminUserCheckBindMobileByMobile struct {
	MetaApiAdminUserCheckBindMobileByMobile
	Ack    *AckAdminUserCheckBindMobileByMobile
	Params struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiAdminUserCheckBindMobileByMobile) GetQuery() interface{}  { return nil }
func (m *ApiAdminUserCheckBindMobileByMobile) GetParams() interface{} { return &m.Params }
func (m *ApiAdminUserCheckBindMobileByMobile) GetAsk() interface{}    { return nil }
func (m *ApiAdminUserCheckBindMobileByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiAdminUserCheckBindMobileByMobile() ApiAdminUserCheckBindMobileByMobile {
	return ApiAdminUserCheckBindMobileByMobile{
		Ack: NewAckAdminUserCheckBindMobileByMobile(),
	}
}

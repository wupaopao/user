package cidl

type MetaApiAdminWxWorkLoginWxWorkAuth struct {
}

var META_ADMIN_WX_WORK_LOGIN_WX_WORK_AUTH = &MetaApiAdminWxWorkLoginWxWorkAuth{}

func (m *MetaApiAdminWxWorkLoginWxWorkAuth) GetMethod() string { return "GET" }
func (m *MetaApiAdminWxWorkLoginWxWorkAuth) GetURL() string {
	return "/user/admin/wx_work/login/wx_work_auth"
}
func (m *MetaApiAdminWxWorkLoginWxWorkAuth) GetName() string { return "AdminWxWorkLoginWxWorkAuth" }
func (m *MetaApiAdminWxWorkLoginWxWorkAuth) GetType() string { return "json" }

// 获取企业微信授权页面跳转
type ApiAdminWxWorkLoginWxWorkAuth struct {
	MetaApiAdminWxWorkLoginWxWorkAuth
	Query struct {
		RedirectUri string `form:"redirect_uri" binding:"required" db:"RedirectUri"`
	}
}

func (m *ApiAdminWxWorkLoginWxWorkAuth) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminWxWorkLoginWxWorkAuth) GetParams() interface{} { return nil }
func (m *ApiAdminWxWorkLoginWxWorkAuth) GetAsk() interface{}    { return nil }
func (m *ApiAdminWxWorkLoginWxWorkAuth) GetAck() interface{}    { return nil }
func MakeApiAdminWxWorkLoginWxWorkAuth() ApiAdminWxWorkLoginWxWorkAuth {
	return ApiAdminWxWorkLoginWxWorkAuth{}
}

// 企业微信授权参数
type AckAdminWxWorkLoginWxWorkAuthData struct {
	AppId       string `json:"appid" db:"AppId"`
	AgentId     int32  `json:"agentid" db:"AgentId"`
	RedirectUri string `json:"redirect_uri" db:"RedirectUri"`
	State       string `json:"state" db:"State"`
}

func NewAckAdminWxWorkLoginWxWorkAuthData() *AckAdminWxWorkLoginWxWorkAuthData {
	return &AckAdminWxWorkLoginWxWorkAuthData{}
}

type MetaApiAdminWxWorkLoginWxWorkAuthData struct {
}

var META_ADMIN_WX_WORK_LOGIN_WX_WORK_AUTH_DATA = &MetaApiAdminWxWorkLoginWxWorkAuthData{}

func (m *MetaApiAdminWxWorkLoginWxWorkAuthData) GetMethod() string { return "GET" }
func (m *MetaApiAdminWxWorkLoginWxWorkAuthData) GetURL() string {
	return "/user/admin/wx_work/login/wx_work_auth_data"
}
func (m *MetaApiAdminWxWorkLoginWxWorkAuthData) GetName() string {
	return "AdminWxWorkLoginWxWorkAuthData"
}
func (m *MetaApiAdminWxWorkLoginWxWorkAuthData) GetType() string { return "json" }

// 获取企业微信授权数据
type ApiAdminWxWorkLoginWxWorkAuthData struct {
	MetaApiAdminWxWorkLoginWxWorkAuthData
	Ack   *AckAdminWxWorkLoginWxWorkAuthData
	Query struct {
		RedirectUri string `form:"redirect_uri" binding:"required" db:"RedirectUri"`
	}
}

func (m *ApiAdminWxWorkLoginWxWorkAuthData) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminWxWorkLoginWxWorkAuthData) GetParams() interface{} { return nil }
func (m *ApiAdminWxWorkLoginWxWorkAuthData) GetAsk() interface{}    { return nil }
func (m *ApiAdminWxWorkLoginWxWorkAuthData) GetAck() interface{}    { return m.Ack }
func MakeApiAdminWxWorkLoginWxWorkAuthData() ApiAdminWxWorkLoginWxWorkAuthData {
	return ApiAdminWxWorkLoginWxWorkAuthData{
		Ack: NewAckAdminWxWorkLoginWxWorkAuthData(),
	}
}

type MetaApiAdminWxWorkLoginWxWorkCallback struct {
}

var META_ADMIN_WX_WORK_LOGIN_WX_WORK_CALLBACK = &MetaApiAdminWxWorkLoginWxWorkCallback{}

func (m *MetaApiAdminWxWorkLoginWxWorkCallback) GetMethod() string { return "GET" }
func (m *MetaApiAdminWxWorkLoginWxWorkCallback) GetURL() string {
	return "/user/admin/wx_work/login/wx_work_callback"
}
func (m *MetaApiAdminWxWorkLoginWxWorkCallback) GetName() string {
	return "AdminWxWorkLoginWxWorkCallback"
}
func (m *MetaApiAdminWxWorkLoginWxWorkCallback) GetType() string { return "json" }

// 企业微信回调
type ApiAdminWxWorkLoginWxWorkCallback struct {
	MetaApiAdminWxWorkLoginWxWorkCallback
	Query struct {
		RedirectUri string `form:"redirect_uri" binding:"required" db:"RedirectUri"`
		Code        string `form:"code" db:"Code"`
		State       string `form:"state" binding:"required" db:"State"`
	}
}

func (m *ApiAdminWxWorkLoginWxWorkCallback) GetQuery() interface{}  { return &m.Query }
func (m *ApiAdminWxWorkLoginWxWorkCallback) GetParams() interface{} { return nil }
func (m *ApiAdminWxWorkLoginWxWorkCallback) GetAsk() interface{}    { return nil }
func (m *ApiAdminWxWorkLoginWxWorkCallback) GetAck() interface{}    { return nil }
func MakeApiAdminWxWorkLoginWxWorkCallback() ApiAdminWxWorkLoginWxWorkCallback {
	return ApiAdminWxWorkLoginWxWorkCallback{}
}

type MetaApiAdminWxWorkLoginLogout struct {
}

var META_ADMIN_WX_WORK_LOGIN_LOGOUT = &MetaApiAdminWxWorkLoginLogout{}

func (m *MetaApiAdminWxWorkLoginLogout) GetMethod() string { return "GET" }
func (m *MetaApiAdminWxWorkLoginLogout) GetURL() string    { return "/user/admin/wx_work/login/logout" }
func (m *MetaApiAdminWxWorkLoginLogout) GetName() string   { return "AdminWxWorkLoginLogout" }
func (m *MetaApiAdminWxWorkLoginLogout) GetType() string   { return "json" }

// 运营账户登出
type ApiAdminWxWorkLoginLogout struct {
	MetaApiAdminWxWorkLoginLogout
}

func (m *ApiAdminWxWorkLoginLogout) GetQuery() interface{}  { return nil }
func (m *ApiAdminWxWorkLoginLogout) GetParams() interface{} { return nil }
func (m *ApiAdminWxWorkLoginLogout) GetAsk() interface{}    { return nil }
func (m *ApiAdminWxWorkLoginLogout) GetAck() interface{}    { return nil }
func MakeApiAdminWxWorkLoginLogout() ApiAdminWxWorkLoginLogout {
	return ApiAdminWxWorkLoginLogout{}
}

type AckAdminWxWorkUserBasicInfo struct {
	UserID string `json:"uid" db:"UserID"`
	Name   string `json:"name" db:"Name"`
	Avatar string `json:"avatar" db:"Avatar"`
}

func NewAckAdminWxWorkUserBasicInfo() *AckAdminWxWorkUserBasicInfo {
	return &AckAdminWxWorkUserBasicInfo{}
}

type MetaApiAdminWxWorkUserBasicInfo struct {
}

var META_ADMIN_WX_WORK_USER_BASIC_INFO = &MetaApiAdminWxWorkUserBasicInfo{}

func (m *MetaApiAdminWxWorkUserBasicInfo) GetMethod() string { return "GET" }
func (m *MetaApiAdminWxWorkUserBasicInfo) GetURL() string {
	return "/user/admin/wx_work/user/basic_info"
}
func (m *MetaApiAdminWxWorkUserBasicInfo) GetName() string { return "AdminWxWorkUserBasicInfo" }
func (m *MetaApiAdminWxWorkUserBasicInfo) GetType() string { return "json" }

// 获取用户基本信息
type ApiAdminWxWorkUserBasicInfo struct {
	MetaApiAdminWxWorkUserBasicInfo
	Ack *AckAdminWxWorkUserBasicInfo
}

func (m *ApiAdminWxWorkUserBasicInfo) GetQuery() interface{}  { return nil }
func (m *ApiAdminWxWorkUserBasicInfo) GetParams() interface{} { return nil }
func (m *ApiAdminWxWorkUserBasicInfo) GetAsk() interface{}    { return nil }
func (m *ApiAdminWxWorkUserBasicInfo) GetAck() interface{}    { return m.Ack }
func MakeApiAdminWxWorkUserBasicInfo() ApiAdminWxWorkUserBasicInfo {
	return ApiAdminWxWorkUserBasicInfo{
		Ack: NewAckAdminWxWorkUserBasicInfo(),
	}
}

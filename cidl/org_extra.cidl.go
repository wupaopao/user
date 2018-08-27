package cidl

type AskUserOrgUserLogin struct {
	Mobile     string `binding:"required,numeric" db:"Mobile"`
	VerifyCode string `binding:"required" db:"VerifyCode"`
}

func NewAskUserOrgUserLogin() *AskUserOrgUserLogin {
	return &AskUserOrgUserLogin{}
}

type AckUserOrgUserLogin struct {
	AuthCityStaff
	UserId string         `db:"UserId"`
	Token  string         `db:"Token"`
	User   *User          `db:"User"`
	Staff  *AuthCityStaff `db:"Staff"`
}

func NewAckUserOrgUserLogin() *AckUserOrgUserLogin {
	return &AckUserOrgUserLogin{
		User:  NewUser(),
		Staff: NewAuthCityStaff(),
	}
}

type MetaApiUserOrgUserLogin struct {
}

var META_USER_ORG_USER_LOGIN = &MetaApiUserOrgUserLogin{}

func (m *MetaApiUserOrgUserLogin) GetMethod() string { return "POST" }
func (m *MetaApiUserOrgUserLogin) GetURL() string    { return "/user/org/user/login" }
func (m *MetaApiUserOrgUserLogin) GetName() string   { return "UserOrgUserLogin" }
func (m *MetaApiUserOrgUserLogin) GetType() string   { return "json" }

// 登陆
type ApiUserOrgUserLogin struct {
	MetaApiUserOrgUserLogin
	Ask *AskUserOrgUserLogin
	Ack *AckUserOrgUserLogin
}

func (m *ApiUserOrgUserLogin) GetQuery() interface{}  { return nil }
func (m *ApiUserOrgUserLogin) GetParams() interface{} { return nil }
func (m *ApiUserOrgUserLogin) GetAsk() interface{}    { return m.Ask }
func (m *ApiUserOrgUserLogin) GetAck() interface{}    { return m.Ack }
func MakeApiUserOrgUserLogin() ApiUserOrgUserLogin {
	return ApiUserOrgUserLogin{
		Ask: NewAskUserOrgUserLogin(),
		Ack: NewAckUserOrgUserLogin(),
	}
}

type MetaApiUserOrgUserLogout struct {
}

var META_USER_ORG_USER_LOGOUT = &MetaApiUserOrgUserLogout{}

func (m *MetaApiUserOrgUserLogout) GetMethod() string { return "POST" }
func (m *MetaApiUserOrgUserLogout) GetURL() string    { return "/user/org/user/logout" }
func (m *MetaApiUserOrgUserLogout) GetName() string   { return "UserOrgUserLogout" }
func (m *MetaApiUserOrgUserLogout) GetType() string   { return "json" }

// 退出登陆
type ApiUserOrgUserLogout struct {
	MetaApiUserOrgUserLogout
}

func (m *ApiUserOrgUserLogout) GetQuery() interface{}  { return nil }
func (m *ApiUserOrgUserLogout) GetParams() interface{} { return nil }
func (m *ApiUserOrgUserLogout) GetAsk() interface{}    { return nil }
func (m *ApiUserOrgUserLogout) GetAck() interface{}    { return nil }
func MakeApiUserOrgUserLogout() ApiUserOrgUserLogout {
	return ApiUserOrgUserLogout{}
}

type MetaApiOrgUserLoginMobileVerify struct {
}

var META_ORG_USER_LOGIN_MOBILE_VERIFY = &MetaApiOrgUserLoginMobileVerify{}

func (m *MetaApiOrgUserLoginMobileVerify) GetMethod() string { return "GET" }
func (m *MetaApiOrgUserLoginMobileVerify) GetURL() string    { return "/user/org/user/login_mobile_verify" }
func (m *MetaApiOrgUserLoginMobileVerify) GetName() string   { return "OrgUserLoginMobileVerify" }
func (m *MetaApiOrgUserLoginMobileVerify) GetType() string   { return "json" }

// 登陆获取手机验证码
type ApiOrgUserLoginMobileVerify struct {
	MetaApiOrgUserLoginMobileVerify
	Query struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiOrgUserLoginMobileVerify) GetQuery() interface{}  { return &m.Query }
func (m *ApiOrgUserLoginMobileVerify) GetParams() interface{} { return nil }
func (m *ApiOrgUserLoginMobileVerify) GetAsk() interface{}    { return nil }
func (m *ApiOrgUserLoginMobileVerify) GetAck() interface{}    { return nil }
func MakeApiOrgUserLoginMobileVerify() ApiOrgUserLoginMobileVerify {
	return ApiOrgUserLoginMobileVerify{}
}

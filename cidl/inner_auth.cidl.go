package cidl

type AskInnerUserAuthTokenInfoAdmin struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoAdmin() *AskInnerUserAuthTokenInfoAdmin {
	return &AskInnerUserAuthTokenInfoAdmin{}
}

type MetaApiInnerUserAuthTokenInfoAdmin struct {
}

var META_INNER_USER_AUTH_TOKEN_INFO_ADMIN = &MetaApiInnerUserAuthTokenInfoAdmin{}

func (m *MetaApiInnerUserAuthTokenInfoAdmin) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserAuthTokenInfoAdmin) GetURL() string {
	return "/inner/user/auth/token_info/admin"
}
func (m *MetaApiInnerUserAuthTokenInfoAdmin) GetName() string { return "InnerUserAuthTokenInfoAdmin" }
func (m *MetaApiInnerUserAuthTokenInfoAdmin) GetType() string { return "json" }

type ApiInnerUserAuthTokenInfoAdmin struct {
	MetaApiInnerUserAuthTokenInfoAdmin
	Ask *AskInnerUserAuthTokenInfoAdmin
	Ack *AuthAdmin
}

func (m *ApiInnerUserAuthTokenInfoAdmin) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserAuthTokenInfoAdmin) GetParams() interface{} { return nil }
func (m *ApiInnerUserAuthTokenInfoAdmin) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserAuthTokenInfoAdmin) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserAuthTokenInfoAdmin() ApiInnerUserAuthTokenInfoAdmin {
	return ApiInnerUserAuthTokenInfoAdmin{
		Ask: NewAskInnerUserAuthTokenInfoAdmin(),
		Ack: NewAuthAdmin(),
	}
}

type AskInnerUserAuthTokenInfoOrg struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoOrg() *AskInnerUserAuthTokenInfoOrg {
	return &AskInnerUserAuthTokenInfoOrg{}
}

type MetaApiInnerUserAuthTokenInfoOrg struct {
}

var META_INNER_USER_AUTH_TOKEN_INFO_ORG = &MetaApiInnerUserAuthTokenInfoOrg{}

func (m *MetaApiInnerUserAuthTokenInfoOrg) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserAuthTokenInfoOrg) GetURL() string    { return "/inner/user/auth/token_info/org" }
func (m *MetaApiInnerUserAuthTokenInfoOrg) GetName() string   { return "InnerUserAuthTokenInfoOrg" }
func (m *MetaApiInnerUserAuthTokenInfoOrg) GetType() string   { return "json" }

// org授权信息
type ApiInnerUserAuthTokenInfoOrg struct {
	MetaApiInnerUserAuthTokenInfoOrg
	Ask *AskInnerUserAuthTokenInfoOrg
	Ack *AuthCity
}

func (m *ApiInnerUserAuthTokenInfoOrg) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserAuthTokenInfoOrg) GetParams() interface{} { return nil }
func (m *ApiInnerUserAuthTokenInfoOrg) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserAuthTokenInfoOrg) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserAuthTokenInfoOrg() ApiInnerUserAuthTokenInfoOrg {
	return ApiInnerUserAuthTokenInfoOrg{
		Ask: NewAskInnerUserAuthTokenInfoOrg(),
		Ack: NewAuthCity(),
	}
}

type AskInnerUserAuthTokenInfoWxXcx struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoWxXcx() *AskInnerUserAuthTokenInfoWxXcx {
	return &AskInnerUserAuthTokenInfoWxXcx{}
}

type MetaApiInnerUserAuthTokenInfoWxXcx struct {
}

var META_INNER_USER_AUTH_TOKEN_INFO_WX_XCX = &MetaApiInnerUserAuthTokenInfoWxXcx{}

func (m *MetaApiInnerUserAuthTokenInfoWxXcx) GetMethod() string { return "POST" }
func (m *MetaApiInnerUserAuthTokenInfoWxXcx) GetURL() string {
	return "/inner/user/auth/token_info/wx_xcx"
}
func (m *MetaApiInnerUserAuthTokenInfoWxXcx) GetName() string { return "InnerUserAuthTokenInfoWxXcx" }
func (m *MetaApiInnerUserAuthTokenInfoWxXcx) GetType() string { return "json" }

// 微信小程序授权信息
type ApiInnerUserAuthTokenInfoWxXcx struct {
	MetaApiInnerUserAuthTokenInfoWxXcx
	Ask *AskInnerUserAuthTokenInfoWxXcx
	Ack *AuthWxXcx
}

func (m *ApiInnerUserAuthTokenInfoWxXcx) GetQuery() interface{}  { return nil }
func (m *ApiInnerUserAuthTokenInfoWxXcx) GetParams() interface{} { return nil }
func (m *ApiInnerUserAuthTokenInfoWxXcx) GetAsk() interface{}    { return m.Ask }
func (m *ApiInnerUserAuthTokenInfoWxXcx) GetAck() interface{}    { return m.Ack }
func MakeApiInnerUserAuthTokenInfoWxXcx() ApiInnerUserAuthTokenInfoWxXcx {
	return ApiInnerUserAuthTokenInfoWxXcx{
		Ask: NewAskInnerUserAuthTokenInfoWxXcx(),
		Ack: NewAuthWxXcx(),
	}
}

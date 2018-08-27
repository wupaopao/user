package cidl

type AskWxXcxUserAuth struct {
	Code          string `binding:"required" db:"Code"`
	EncryptedData string `binding:"required" db:"EncryptedData"`
	InitVector    string `binding:"required" db:"InitVector"`
}

func NewAskWxXcxUserAuth() *AskWxXcxUserAuth {
	return &AskWxXcxUserAuth{}
}

type AckWxXcxUserAuth struct {
	IsVisitor        bool                       `db:"IsVisitor"`
	Token            string                     `db:"Token"`
	UserId           string                     `db:"UserId"`
	User             *User                      `db:"User"`
	Organization     *AuthWxXcxOrganization     `db:"Organization"`
	CommunityManager *AuthWxXcxCommunityManager `db:"CommunityManager"`
}

func NewAckWxXcxUserAuth() *AckWxXcxUserAuth {
	return &AckWxXcxUserAuth{
		User:             NewUser(),
		Organization:     NewAuthWxXcxOrganization(),
		CommunityManager: NewAuthWxXcxCommunityManager(),
	}
}

type MetaApiWxXcxUserAuth struct {
}

var META_WX_XCX_USER_AUTH = &MetaApiWxXcxUserAuth{}

func (m *MetaApiWxXcxUserAuth) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserAuth) GetURL() string    { return "/user/wx_xcx/user/auth" }
func (m *MetaApiWxXcxUserAuth) GetName() string   { return "WxXcxUserAuth" }
func (m *MetaApiWxXcxUserAuth) GetType() string   { return "json" }

// 微信用户信息
// 微信小程序授权登陆
type ApiWxXcxUserAuth struct {
	MetaApiWxXcxUserAuth
	Ask *AskWxXcxUserAuth
	Ack *AckWxXcxUserAuth
}

func (m *ApiWxXcxUserAuth) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserAuth) GetParams() interface{} { return nil }
func (m *ApiWxXcxUserAuth) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxUserAuth) GetAck() interface{}    { return m.Ack }
func MakeApiWxXcxUserAuth() ApiWxXcxUserAuth {
	return ApiWxXcxUserAuth{
		Ask: NewAskWxXcxUserAuth(),
		Ack: NewAckWxXcxUserAuth(),
	}
}

type AskWxXcxUserAuthFake struct {
	FakeUnionId   string `binding:"required" db:"FakeUnionId"`
	Code          string `binding:"required" db:"Code"`
	EncryptedData string `binding:"required" db:"EncryptedData"`
	InitVector    string `binding:"required" db:"InitVector"`
}

func NewAskWxXcxUserAuthFake() *AskWxXcxUserAuthFake {
	return &AskWxXcxUserAuthFake{}
}

type AckWxXcxUserAuthFake struct {
	IsVisitor        bool                       `db:"IsVisitor"`
	Token            string                     `db:"Token"`
	UserId           string                     `db:"UserId"`
	User             *User                      `db:"User"`
	Organization     *AuthWxXcxOrganization     `db:"Organization"`
	CommunityManager *AuthWxXcxCommunityManager `db:"CommunityManager"`
}

func NewAckWxXcxUserAuthFake() *AckWxXcxUserAuthFake {
	return &AckWxXcxUserAuthFake{
		User:             NewUser(),
		Organization:     NewAuthWxXcxOrganization(),
		CommunityManager: NewAuthWxXcxCommunityManager(),
	}
}

type MetaApiWxXcxUserAuthFake struct {
}

var META_WX_XCX_USER_AUTH_FAKE = &MetaApiWxXcxUserAuthFake{}

func (m *MetaApiWxXcxUserAuthFake) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserAuthFake) GetURL() string    { return "/user/wx_xcx/user/auth_fake" }
func (m *MetaApiWxXcxUserAuthFake) GetName() string   { return "WxXcxUserAuthFake" }
func (m *MetaApiWxXcxUserAuthFake) GetType() string   { return "json" }

// TODO 测试用，微信假授权接口，待删除
type ApiWxXcxUserAuthFake struct {
	MetaApiWxXcxUserAuthFake
	Ask *AskWxXcxUserAuthFake
	Ack *AckWxXcxUserAuthFake
}

func (m *ApiWxXcxUserAuthFake) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserAuthFake) GetParams() interface{} { return nil }
func (m *ApiWxXcxUserAuthFake) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxUserAuthFake) GetAck() interface{}    { return m.Ack }
func MakeApiWxXcxUserAuthFake() ApiWxXcxUserAuthFake {
	return ApiWxXcxUserAuthFake{
		Ask: NewAskWxXcxUserAuthFake(),
		Ack: NewAckWxXcxUserAuthFake(),
	}
}

// TODO 测试用，微信假用户，待删除
type FakeUser struct {
	UnionId  string `db:"union_id"`
	Nickname string `db:"nickname"`
}

func NewFakeUser() *FakeUser {
	return &FakeUser{}
}

type AckWxXcxUserFakeUserList struct {
	Users []*FakeUser `db:"Users"`
}

func NewAckWxXcxUserFakeUserList() *AckWxXcxUserFakeUserList {
	return &AckWxXcxUserFakeUserList{
		Users: make([]*FakeUser, 0),
	}
}

type MetaApiWxXcxUserFakeUserList struct {
}

var META_WX_XCX_USER_FAKE_USER_LIST = &MetaApiWxXcxUserFakeUserList{}

func (m *MetaApiWxXcxUserFakeUserList) GetMethod() string { return "GET" }
func (m *MetaApiWxXcxUserFakeUserList) GetURL() string    { return "/user/wx_xcx/user/fake_user_list" }
func (m *MetaApiWxXcxUserFakeUserList) GetName() string   { return "WxXcxUserFakeUserList" }
func (m *MetaApiWxXcxUserFakeUserList) GetType() string   { return "json" }

type ApiWxXcxUserFakeUserList struct {
	MetaApiWxXcxUserFakeUserList
	Ack *AckWxXcxUserFakeUserList
}

func (m *ApiWxXcxUserFakeUserList) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserFakeUserList) GetParams() interface{} { return nil }
func (m *ApiWxXcxUserFakeUserList) GetAsk() interface{}    { return nil }
func (m *ApiWxXcxUserFakeUserList) GetAck() interface{}    { return m.Ack }
func MakeApiWxXcxUserFakeUserList() ApiWxXcxUserFakeUserList {
	return ApiWxXcxUserFakeUserList{
		Ack: NewAckWxXcxUserFakeUserList(),
	}
}

type AskWxXcxUserEditBasicByUserID struct {
	Nickname string `binding:"required,lte=64" db:"Nickname"`
	Mobile   string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskWxXcxUserEditBasicByUserID() *AskWxXcxUserEditBasicByUserID {
	return &AskWxXcxUserEditBasicByUserID{}
}

type MetaApiWxXcxUserEditBasicByUserID struct {
}

var META_WX_XCX_USER_EDIT_BASIC_BY_USER_ID = &MetaApiWxXcxUserEditBasicByUserID{}

func (m *MetaApiWxXcxUserEditBasicByUserID) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserEditBasicByUserID) GetURL() string {
	return "/user/wx_xcx/user/edit_basic/:user_id"
}
func (m *MetaApiWxXcxUserEditBasicByUserID) GetName() string { return "WxXcxUserEditBasicByUserID" }
func (m *MetaApiWxXcxUserEditBasicByUserID) GetType() string { return "json" }

// 修改基本信息
type ApiWxXcxUserEditBasicByUserID struct {
	MetaApiWxXcxUserEditBasicByUserID
	Ask    *AskWxXcxUserEditBasicByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiWxXcxUserEditBasicByUserID) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserEditBasicByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxUserEditBasicByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxUserEditBasicByUserID) GetAck() interface{}    { return nil }
func MakeApiWxXcxUserEditBasicByUserID() ApiWxXcxUserEditBasicByUserID {
	return ApiWxXcxUserEditBasicByUserID{
		Ask: NewAskWxXcxUserEditBasicByUserID(),
	}
}

type AskWxXcxUserEditNicknameByUserID struct {
	Nickname string `binding:"required,lte=64" db:"Nickname"`
}

func NewAskWxXcxUserEditNicknameByUserID() *AskWxXcxUserEditNicknameByUserID {
	return &AskWxXcxUserEditNicknameByUserID{}
}

type MetaApiWxXcxUserEditNicknameByUserID struct {
}

var META_WX_XCX_USER_EDIT_NICKNAME_BY_USER_ID = &MetaApiWxXcxUserEditNicknameByUserID{}

func (m *MetaApiWxXcxUserEditNicknameByUserID) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserEditNicknameByUserID) GetURL() string {
	return "/user/wx_xcx/user/edit_nickname/:user_id"
}
func (m *MetaApiWxXcxUserEditNicknameByUserID) GetName() string {
	return "WxXcxUserEditNicknameByUserID"
}
func (m *MetaApiWxXcxUserEditNicknameByUserID) GetType() string { return "json" }

// 修改昵称
type ApiWxXcxUserEditNicknameByUserID struct {
	MetaApiWxXcxUserEditNicknameByUserID
	Ask    *AskWxXcxUserEditNicknameByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiWxXcxUserEditNicknameByUserID) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserEditNicknameByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxUserEditNicknameByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxUserEditNicknameByUserID) GetAck() interface{}    { return nil }
func MakeApiWxXcxUserEditNicknameByUserID() ApiWxXcxUserEditNicknameByUserID {
	return ApiWxXcxUserEditNicknameByUserID{
		Ask: NewAskWxXcxUserEditNicknameByUserID(),
	}
}

type MetaApiWxXcxUserMobileVerify struct {
}

var META_WX_XCX_USER_MOBILE_VERIFY = &MetaApiWxXcxUserMobileVerify{}

func (m *MetaApiWxXcxUserMobileVerify) GetMethod() string { return "GET" }
func (m *MetaApiWxXcxUserMobileVerify) GetURL() string    { return "/user/wx_xcx/user/mobile_verify" }
func (m *MetaApiWxXcxUserMobileVerify) GetName() string   { return "WxXcxUserMobileVerify" }
func (m *MetaApiWxXcxUserMobileVerify) GetType() string   { return "json" }

// 修改手机获取手机验证码
type ApiWxXcxUserMobileVerify struct {
	MetaApiWxXcxUserMobileVerify
	Query struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiWxXcxUserMobileVerify) GetQuery() interface{}  { return &m.Query }
func (m *ApiWxXcxUserMobileVerify) GetParams() interface{} { return nil }
func (m *ApiWxXcxUserMobileVerify) GetAsk() interface{}    { return nil }
func (m *ApiWxXcxUserMobileVerify) GetAck() interface{}    { return nil }
func MakeApiWxXcxUserMobileVerify() ApiWxXcxUserMobileVerify {
	return ApiWxXcxUserMobileVerify{}
}

type AskWxXcxUserEditMobileByUserID struct {
	Mobile     string `binding:"required,numeric" db:"Mobile"`
	VerifyCode string `binding:"required" db:"VerifyCode"`
}

func NewAskWxXcxUserEditMobileByUserID() *AskWxXcxUserEditMobileByUserID {
	return &AskWxXcxUserEditMobileByUserID{}
}

type MetaApiWxXcxUserEditMobileByUserID struct {
}

var META_WX_XCX_USER_EDIT_MOBILE_BY_USER_ID = &MetaApiWxXcxUserEditMobileByUserID{}

func (m *MetaApiWxXcxUserEditMobileByUserID) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserEditMobileByUserID) GetURL() string {
	return "/user/wx_xcx/user/edit_mobile/:user_id"
}
func (m *MetaApiWxXcxUserEditMobileByUserID) GetName() string { return "WxXcxUserEditMobileByUserID" }
func (m *MetaApiWxXcxUserEditMobileByUserID) GetType() string { return "json" }

// 修改手机
type ApiWxXcxUserEditMobileByUserID struct {
	MetaApiWxXcxUserEditMobileByUserID
	Ask    *AskWxXcxUserEditMobileByUserID
	Params struct {
		UserID string `form:"user_id" db:"UserID"`
	}
}

func (m *ApiWxXcxUserEditMobileByUserID) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserEditMobileByUserID) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxUserEditMobileByUserID) GetAsk() interface{}    { return m.Ask }
func (m *ApiWxXcxUserEditMobileByUserID) GetAck() interface{}    { return nil }
func MakeApiWxXcxUserEditMobileByUserID() ApiWxXcxUserEditMobileByUserID {
	return ApiWxXcxUserEditMobileByUserID{
		Ask: NewAskWxXcxUserEditMobileByUserID(),
	}
}

type MetaApiUserWxXcxUserWxBindMobileVerify struct {
}

var META_USER_WX_XCX_USER_WX_BIND_MOBILE_VERIFY = &MetaApiUserWxXcxUserWxBindMobileVerify{}

func (m *MetaApiUserWxXcxUserWxBindMobileVerify) GetMethod() string { return "GET" }
func (m *MetaApiUserWxXcxUserWxBindMobileVerify) GetURL() string {
	return "/user/wx_xcx/user/wx_bind_mobile_verify"
}
func (m *MetaApiUserWxXcxUserWxBindMobileVerify) GetName() string {
	return "UserWxXcxUserWxBindMobileVerify"
}
func (m *MetaApiUserWxXcxUserWxBindMobileVerify) GetType() string { return "json" }

// 微信号绑定手机号手机验证码
type ApiUserWxXcxUserWxBindMobileVerify struct {
	MetaApiUserWxXcxUserWxBindMobileVerify
	Query struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiUserWxXcxUserWxBindMobileVerify) GetQuery() interface{}  { return &m.Query }
func (m *ApiUserWxXcxUserWxBindMobileVerify) GetParams() interface{} { return nil }
func (m *ApiUserWxXcxUserWxBindMobileVerify) GetAsk() interface{}    { return nil }
func (m *ApiUserWxXcxUserWxBindMobileVerify) GetAck() interface{}    { return nil }
func MakeApiUserWxXcxUserWxBindMobileVerify() ApiUserWxXcxUserWxBindMobileVerify {
	return ApiUserWxXcxUserWxBindMobileVerify{}
}

type AskUserWxXcxUserWxBindMobile struct {
	Mobile     string `binding:"required,numeric" db:"Mobile"`
	VerifyCode string `binding:"required" db:"VerifyCode"`
}

func NewAskUserWxXcxUserWxBindMobile() *AskUserWxXcxUserWxBindMobile {
	return &AskUserWxXcxUserWxBindMobile{}
}

type MetaApiUserWxXcxUserWxBindMobile struct {
}

var META_USER_WX_XCX_USER_WX_BIND_MOBILE = &MetaApiUserWxXcxUserWxBindMobile{}

func (m *MetaApiUserWxXcxUserWxBindMobile) GetMethod() string { return "POST" }
func (m *MetaApiUserWxXcxUserWxBindMobile) GetURL() string    { return "/user/wx_xcx/user/wx_bind_mobile" }
func (m *MetaApiUserWxXcxUserWxBindMobile) GetName() string   { return "UserWxXcxUserWxBindMobile" }
func (m *MetaApiUserWxXcxUserWxBindMobile) GetType() string   { return "json" }

// 微信绑定手机号
type ApiUserWxXcxUserWxBindMobile struct {
	MetaApiUserWxXcxUserWxBindMobile
	Ask *AskUserWxXcxUserWxBindMobile
}

func (m *ApiUserWxXcxUserWxBindMobile) GetQuery() interface{}  { return nil }
func (m *ApiUserWxXcxUserWxBindMobile) GetParams() interface{} { return nil }
func (m *ApiUserWxXcxUserWxBindMobile) GetAsk() interface{}    { return m.Ask }
func (m *ApiUserWxXcxUserWxBindMobile) GetAck() interface{}    { return nil }
func MakeApiUserWxXcxUserWxBindMobile() ApiUserWxXcxUserWxBindMobile {
	return ApiUserWxXcxUserWxBindMobile{
		Ask: NewAskUserWxXcxUserWxBindMobile(),
	}
}

type AckWxXcxUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckWxXcxUserCheckBindMobileByMobile() *AckWxXcxUserCheckBindMobileByMobile {
	return &AckWxXcxUserCheckBindMobileByMobile{}
}

type MetaApiWxXcxUserCheckBindMobileByMobile struct {
}

var META_WX_XCX_USER_CHECK_BIND_MOBILE_BY_MOBILE = &MetaApiWxXcxUserCheckBindMobileByMobile{}

func (m *MetaApiWxXcxUserCheckBindMobileByMobile) GetMethod() string { return "GET" }
func (m *MetaApiWxXcxUserCheckBindMobileByMobile) GetURL() string {
	return "/user/wx_xcx/user/check_bind_mobile/:mobile"
}
func (m *MetaApiWxXcxUserCheckBindMobileByMobile) GetName() string {
	return "WxXcxUserCheckBindMobileByMobile"
}
func (m *MetaApiWxXcxUserCheckBindMobileByMobile) GetType() string { return "json" }

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
type ApiWxXcxUserCheckBindMobileByMobile struct {
	MetaApiWxXcxUserCheckBindMobileByMobile
	Ack    *AckWxXcxUserCheckBindMobileByMobile
	Params struct {
		Mobile string `form:"mobile" binding:"required,numeric" db:"Mobile"`
	}
}

func (m *ApiWxXcxUserCheckBindMobileByMobile) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserCheckBindMobileByMobile) GetParams() interface{} { return &m.Params }
func (m *ApiWxXcxUserCheckBindMobileByMobile) GetAsk() interface{}    { return nil }
func (m *ApiWxXcxUserCheckBindMobileByMobile) GetAck() interface{}    { return m.Ack }
func MakeApiWxXcxUserCheckBindMobileByMobile() ApiWxXcxUserCheckBindMobileByMobile {
	return ApiWxXcxUserCheckBindMobileByMobile{
		Ack: NewAckWxXcxUserCheckBindMobileByMobile(),
	}
}

type MetaApiWxXcxUserUnbindWx struct {
}

var META_WX_XCX_USER_UNBIND_WX = &MetaApiWxXcxUserUnbindWx{}

func (m *MetaApiWxXcxUserUnbindWx) GetMethod() string { return "POST" }
func (m *MetaApiWxXcxUserUnbindWx) GetURL() string    { return "/user/wx_xcx/user/unbind_wx" }
func (m *MetaApiWxXcxUserUnbindWx) GetName() string   { return "WxXcxUserUnbindWx" }
func (m *MetaApiWxXcxUserUnbindWx) GetType() string   { return "json" }

// 用户解绑微信
type ApiWxXcxUserUnbindWx struct {
	MetaApiWxXcxUserUnbindWx
}

func (m *ApiWxXcxUserUnbindWx) GetQuery() interface{}  { return nil }
func (m *ApiWxXcxUserUnbindWx) GetParams() interface{} { return nil }
func (m *ApiWxXcxUserUnbindWx) GetAsk() interface{}    { return nil }
func (m *ApiWxXcxUserUnbindWx) GetAck() interface{}    { return nil }
func MakeApiWxXcxUserUnbindWx() ApiWxXcxUserUnbindWx {
	return ApiWxXcxUserUnbindWx{}
}

package user

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

// 微信用户信息
// 微信小程序授权登陆
func (m *Proxy) WxXcxUserAuth(ask *AskWxXcxUserAuth,
) (*AckWxXcxUserAuth, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckWxXcxUserAuth
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/auth",
		ask,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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

// TODO 测试用，微信假授权接口，待删除
func (m *Proxy) WxXcxUserAuthFake(ask *AskWxXcxUserAuthFake,
) (*AckWxXcxUserAuthFake, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckWxXcxUserAuthFake
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/auth_fake",
		ask,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
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
func (m *Proxy) WxXcxUserFakeUserList() (*AckWxXcxUserFakeUserList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckWxXcxUserFakeUserList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/wx_xcx/user/fake_user_list",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AskWxXcxUserEditBasicByUserID struct {
	Nickname string `binding:"required,lte=64" db:"Nickname"`
	Mobile   string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskWxXcxUserEditBasicByUserID() *AskWxXcxUserEditBasicByUserID {
	return &AskWxXcxUserEditBasicByUserID{}
}

// 修改基本信息
func (m *Proxy) WxXcxUserEditBasicByUserID(UserID string,
	ask *AskWxXcxUserEditBasicByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/edit_basic/:user_id",
		ask,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AskWxXcxUserEditNicknameByUserID struct {
	Nickname string `binding:"required,lte=64" db:"Nickname"`
}

func NewAskWxXcxUserEditNicknameByUserID() *AskWxXcxUserEditNicknameByUserID {
	return &AskWxXcxUserEditNicknameByUserID{}
}

// 修改昵称
func (m *Proxy) WxXcxUserEditNicknameByUserID(UserID string,
	ask *AskWxXcxUserEditNicknameByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/edit_nickname/:user_id",
		ask,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

// 修改手机获取手机验证码
func (m *Proxy) WxXcxUserMobileVerify() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/wx_xcx/user/mobile_verify",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AskWxXcxUserEditMobileByUserID struct {
	Mobile     string `binding:"required,numeric" db:"Mobile"`
	VerifyCode string `binding:"required" db:"VerifyCode"`
}

func NewAskWxXcxUserEditMobileByUserID() *AskWxXcxUserEditMobileByUserID {
	return &AskWxXcxUserEditMobileByUserID{}
}

// 修改手机
func (m *Proxy) WxXcxUserEditMobileByUserID(UserID string,
	ask *AskWxXcxUserEditMobileByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/edit_mobile/:user_id",
		ask,
		ack,
		map[string]interface{}{
			"user_id": UserID,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

// 微信号绑定手机号手机验证码
func (m *Proxy) UserWxXcxUserWxBindMobileVerify() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/wx_xcx/user/wx_bind_mobile_verify",
		nil,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AskUserWxXcxUserWxBindMobile struct {
	Mobile     string `binding:"required,numeric" db:"Mobile"`
	VerifyCode string `binding:"required" db:"VerifyCode"`
}

func NewAskUserWxXcxUserWxBindMobile() *AskUserWxXcxUserWxBindMobile {
	return &AskUserWxXcxUserWxBindMobile{}
}

// 微信绑定手机号
func (m *Proxy) UserWxXcxUserWxBindMobile(ask *AskUserWxXcxUserWxBindMobile,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/wx_xcx/user/wx_bind_mobile",
		ask,
		ack,
		nil,
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

type AckWxXcxUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckWxXcxUserCheckBindMobileByMobile() *AckWxXcxUserCheckBindMobileByMobile {
	return &AckWxXcxUserCheckBindMobileByMobile{}
}

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
func (m *Proxy) WxXcxUserCheckBindMobileByMobile(Mobile string,
) (*AckWxXcxUserCheckBindMobileByMobile, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckWxXcxUserCheckBindMobileByMobile
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/wx_xcx/user/check_bind_mobile/:mobile",
		nil,
		ack,
		map[string]interface{}{
			"mobile": Mobile,
		},
	)
	if err != nil {
		return nil, err
	}
	if ack.Code != 0 {
		return nil, m.Error(ack.Code, ack.Message)
	}
	return ack.Data, nil
}

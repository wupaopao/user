package user

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

// 登陆
func (m *Proxy) UserOrgUserLogin(ask *AskUserOrgUserLogin,
) (*AckUserOrgUserLogin, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckUserOrgUserLogin
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/org/user/login",
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

// 退出登陆
func (m *Proxy) UserOrgUserLogout() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/org/user/logout",
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

// 登陆获取手机验证码
func (m *Proxy) OrgUserLoginMobileVerify() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/login_mobile_verify",
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

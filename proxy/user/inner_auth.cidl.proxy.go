package user

type AskInnerUserAuthTokenInfoAdmin struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoAdmin() *AskInnerUserAuthTokenInfoAdmin {
	return &AskInnerUserAuthTokenInfoAdmin{}
}
func (m *Proxy) InnerUserAuthTokenInfoAdmin(ask *AskInnerUserAuthTokenInfoAdmin,
) (*AuthAdmin, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AuthAdmin
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/auth/token_info/admin",
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

type AskInnerUserAuthTokenInfoOrg struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoOrg() *AskInnerUserAuthTokenInfoOrg {
	return &AskInnerUserAuthTokenInfoOrg{}
}

// org授权信息
func (m *Proxy) InnerUserAuthTokenInfoOrg(ask *AskInnerUserAuthTokenInfoOrg,
) (*AuthCity, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AuthCity
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/auth/token_info/org",
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

type AskInnerUserAuthTokenInfoWxXcx struct {
	UserID string `binding:"required" db:"UserID"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserAuthTokenInfoWxXcx() *AskInnerUserAuthTokenInfoWxXcx {
	return &AskInnerUserAuthTokenInfoWxXcx{}
}

// 微信小程序授权信息
func (m *Proxy) InnerUserAuthTokenInfoWxXcx(ask *AskInnerUserAuthTokenInfoWxXcx,
) (*AuthWxXcx, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AuthWxXcx
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/auth/token_info/wx_xcx",
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

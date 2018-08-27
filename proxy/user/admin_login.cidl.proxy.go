package user

// 获取企业微信授权页面跳转
func (m *Proxy) AdminWxWorkLoginWxWorkAuth() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/wx_work/login/wx_work_auth",
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

// 获取企业微信授权数据
func (m *Proxy) AdminWxWorkLoginWxWorkAuthData() (*AckAdminWxWorkLoginWxWorkAuthData, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminWxWorkLoginWxWorkAuthData
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/wx_work/login/wx_work_auth_data",
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

// 企业微信回调
func (m *Proxy) AdminWxWorkLoginWxWorkCallback() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/wx_work/login/wx_work_callback",
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

// 运营账户登出
func (m *Proxy) AdminWxWorkLoginLogout() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/wx_work/login/logout",
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

type AckAdminWxWorkUserBasicInfo struct {
	UserID string `json:"uid" db:"UserID"`
	Name   string `json:"name" db:"Name"`
	Avatar string `json:"avatar" db:"Avatar"`
}

func NewAckAdminWxWorkUserBasicInfo() *AckAdminWxWorkUserBasicInfo {
	return &AckAdminWxWorkUserBasicInfo{}
}

// 获取用户基本信息
func (m *Proxy) AdminWxWorkUserBasicInfo() (*AckAdminWxWorkUserBasicInfo, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminWxWorkUserBasicInfo
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/wx_work/user/basic_info",
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

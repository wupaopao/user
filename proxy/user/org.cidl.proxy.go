package user

// 用户信息
func (m *Proxy) OrgUserInfoByUserId(UserID string,
) (*User, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *User
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/info/:user_id",
		nil,
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

type AckOrgUserOrganizationManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserOrganizationManagerList() *AckOrgUserOrganizationManagerList {
	return &AckOrgUserOrganizationManagerList{
		List: make([]*User, 0),
	}
}

// 城市合伙人用户
func (m *Proxy) OrgUserOrganizationManagerList() (*AckOrgUserOrganizationManagerList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgUserOrganizationManagerList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/organization_manager/list",
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

type AckOrgUserCommunityManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserCommunityManagerList() *AckOrgUserCommunityManagerList {
	return &AckOrgUserCommunityManagerList{
		List: make([]*User, 0),
	}
}

// 社区合伙人
func (m *Proxy) OrgUserCommunityManagerList() (*AckOrgUserCommunityManagerList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgUserCommunityManagerList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/community_manager/list",
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

type AckOrgUserOrganizationStaffList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckOrgUserOrganizationStaffList() *AckOrgUserOrganizationStaffList {
	return &AckOrgUserOrganizationStaffList{
		List: make([]*User, 0),
	}
}

// 组织成员
func (m *Proxy) OrgUserOrganizationStaffList() (*AckOrgUserOrganizationStaffList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgUserOrganizationStaffList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/organization_staff/list",
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

// 身份证token
func (m *Proxy) OrgUserIDCardPicToken(ask *AskOrgUserIDCardPicToken,
) (*AckOrgUserIDCardPicToken, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgUserIDCardPicToken
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/org/user/id_card_pic_token",
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

// 访问身份证图片
func (m *Proxy) OrgUserAccessIDCardPic() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/access_id_card_pic",
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

type AckOrgUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckOrgUserCheckBindMobileByMobile() *AckOrgUserCheckBindMobileByMobile {
	return &AckOrgUserCheckBindMobileByMobile{}
}

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
func (m *Proxy) OrgUserCheckBindMobileByMobile(Mobile string,
) (*AckOrgUserCheckBindMobileByMobile, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckOrgUserCheckBindMobileByMobile
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/org/user/check_bind_mobile/:mobile",
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

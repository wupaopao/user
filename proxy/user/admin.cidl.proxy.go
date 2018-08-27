package user

// 用户信息
func (m *Proxy) AdminUserInfoByUserId(UserID string,
) (*User, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *User
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/info/:user_id",
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

// 通过手机号获取用户信息
func (m *Proxy) AdminUserInfoByMobile() (*User, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *User
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/info_by_mobile",
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

type AckAdminUserOrganizationManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserOrganizationManagerList() *AckAdminUserOrganizationManagerList {
	return &AckAdminUserOrganizationManagerList{
		List: make([]*User, 0),
	}
}

// 城市合伙人用户列表
func (m *Proxy) AdminUserOrganizationManagerList() (*AckAdminUserOrganizationManagerList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminUserOrganizationManagerList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/organization_manager/list",
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

type AckAdminUserCommunityManagerList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserCommunityManagerList() *AckAdminUserCommunityManagerList {
	return &AckAdminUserCommunityManagerList{
		List: make([]*User, 0),
	}
}

// 社区合伙人列表
func (m *Proxy) AdminUserCommunityManagerList() (*AckAdminUserCommunityManagerList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminUserCommunityManagerList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/community_manager/list",
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

type AckAdminUserOrganizationStaffList struct {
	Count uint32  `db:"Count"`
	List  []*User `db:"List"`
}

func NewAckAdminUserOrganizationStaffList() *AckAdminUserOrganizationStaffList {
	return &AckAdminUserOrganizationStaffList{
		List: make([]*User, 0),
	}
}

// 组织成员列表
func (m *Proxy) AdminUserOrganizationStaffList() (*AckAdminUserOrganizationStaffList, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminUserOrganizationStaffList
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/organization_staff/list",
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

// 身份证token
func (m *Proxy) AdminUserIDCardPicToken(ask *AskAdminUserIDCardPicToken,
) (*AckAdminUserIDCardPicToken, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminUserIDCardPicToken
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/user/admin/user/id_card_pic_token",
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
func (m *Proxy) AdminUserAccessIDCardPic() (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/access_id_card_pic",
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

type AckAdminUserCheckBindMobileByMobile struct {
	CanBeBound bool   `db:"CanBeBound"`
	UserExist  bool   `db:"UserExist"`
	UserName   string `db:"UserName"`
}

func NewAckAdminUserCheckBindMobileByMobile() *AckAdminUserCheckBindMobileByMobile {
	return &AckAdminUserCheckBindMobileByMobile{}
}

// 检查手机是否已经被绑定为团购组织管理员、组织成员或者社区合伙人
func (m *Proxy) AdminUserCheckBindMobileByMobile(Mobile string,
) (*AckAdminUserCheckBindMobileByMobile, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckAdminUserCheckBindMobileByMobile
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/user/admin/user/check_bind_mobile/:mobile",
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

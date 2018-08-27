package user

// 获取用户
func (m *Proxy) InnerUserInfoByUserID(UserID string,
) (*User, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *User
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/user/user/info/:user_id",
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

type AckInnerUserUserInfoByMobile struct {
	Exist bool  `db:"Exist"`
	User  *User `db:"User"`
}

func NewAckInnerUserUserInfoByMobile() *AckInnerUserUserInfoByMobile {
	return &AckInnerUserUserInfoByMobile{
		User: NewUser(),
	}
}

// 通过手机号获取用户
func (m *Proxy) InnerUserUserInfoByMobile(Mobile string,
) (*AckInnerUserUserInfoByMobile, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerUserUserInfoByMobile
	}
	ack := &Ack{}
	err := m.Invoke(
		"GET",
		"/inner/user/user/info_by_mobile/:mobile",
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

type AskInnerUserCommunityManagerEditByUserID struct {
	IsCmtManager bool `binding:"required" db:"IsCmtManager"`
}

func NewAskInnerUserCommunityManagerEditByUserID() *AskInnerUserCommunityManagerEditByUserID {
	return &AskInnerUserCommunityManagerEditByUserID{}
}

// 设置用户为社区合伙人
func (m *Proxy) InnerUserCommunityManagerEditByUserID(UserID string,
	ask *AskInnerUserCommunityManagerEditByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/community_manager/edit/:user_id",
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

type AskInnerUserUserOrgManagerAddOrUpdate struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	Nickname     string `binding:"required,lte=64" db:"Nickname"`
	IdCardNumber string `binding:"required,lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"required,lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"required,lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserOrgManagerAddOrUpdate() *AskInnerUserUserOrgManagerAddOrUpdate {
	return &AskInnerUserUserOrgManagerAddOrUpdate{}
}

type AckInnerUserUserOrgManagerAddOrUpdate struct {
	UserId string `db:"UserId"`
}

func NewAckInnerUserUserOrgManagerAddOrUpdate() *AckInnerUserUserOrgManagerAddOrUpdate {
	return &AckInnerUserUserOrgManagerAddOrUpdate{}
}

// 添加组织管理员用户
func (m *Proxy) InnerUserUserOrgManagerAddOrUpdate(ask *AskInnerUserUserOrgManagerAddOrUpdate,
) (*AckInnerUserUserOrgManagerAddOrUpdate, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerUserUserOrgManagerAddOrUpdate
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/org_manager/add_or_update",
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

type AskInnerUserUserOrgManagerUpdate struct {
	UserId       string `binding:"required" db:"UserId"`
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	Nickname     string `binding:"required,lte=64" db:"Nickname"`
	IdCardNumber string `binding:"required,lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"required,lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"required,lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserOrgManagerUpdate() *AskInnerUserUserOrgManagerUpdate {
	return &AskInnerUserUserOrgManagerUpdate{}
}

// 更新组织管理员用户
func (m *Proxy) InnerUserUserOrgManagerUpdate(ask *AskInnerUserUserOrgManagerUpdate,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/org_manager/update",
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

type AskInnerUserUserOrgManagerUnbind struct {
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserOrgManagerUnbind() *AskInnerUserUserOrgManagerUnbind {
	return &AskInnerUserUserOrgManagerUnbind{}
}

// 解绑组织管理员用户
func (m *Proxy) InnerUserUserOrgManagerUnbind(ask *AskInnerUserUserOrgManagerUnbind,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/org_manager/unbind",
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

type AskInnerUserUserOrgStaffAddOrUpdate struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserUserOrgStaffAddOrUpdate() *AskInnerUserUserOrgStaffAddOrUpdate {
	return &AskInnerUserUserOrgStaffAddOrUpdate{}
}

type AckInnerUserUserOrgStaffAddOrUpdate struct {
	UserId string `db:"UserId"`
}

func NewAckInnerUserUserOrgStaffAddOrUpdate() *AckInnerUserUserOrgStaffAddOrUpdate {
	return &AckInnerUserUserOrgStaffAddOrUpdate{}
}

// 添加组织成员
func (m *Proxy) InnerUserUserOrgStaffAddOrUpdate(ask *AskInnerUserUserOrgStaffAddOrUpdate,
) (*AckInnerUserUserOrgStaffAddOrUpdate, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerUserUserOrgStaffAddOrUpdate
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/org_staff/add_or_update",
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

type AskInnerUserUserOrgStaffUpdateByUserID struct {
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserUserOrgStaffUpdateByUserID() *AskInnerUserUserOrgStaffUpdateByUserID {
	return &AskInnerUserUserOrgStaffUpdateByUserID{}
}

// 更新组织成员
func (m *Proxy) InnerUserUserOrgStaffUpdateByUserID(UserID string,
	ask *AskInnerUserUserOrgStaffUpdateByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/org_staff/update/:user_id",
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

type AskInnerUserUserCmtManagerUpdateByUserID struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	IdCardNumber string `binding:"lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserCmtManagerUpdateByUserID() *AskInnerUserUserCmtManagerUpdateByUserID {
	return &AskInnerUserUserCmtManagerUpdateByUserID{}
}

// 更新社区合伙人信息
func (m *Proxy) InnerUserUserCmtManagerUpdateByUserID(UserID string,
	ask *AskInnerUserUserCmtManagerUpdateByUserID,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/cmt_manager/update/:user_id",
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

type AskInnerUserUserCmtManagerChange struct {
	NewManagerUid string `binding:"required" db:"NewManagerUid"`
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserCmtManagerChange() *AskInnerUserUserCmtManagerChange {
	return &AskInnerUserUserCmtManagerChange{}
}

// 更改社区合伙人
func (m *Proxy) InnerUserUserCmtManagerChange(ask *AskInnerUserUserCmtManagerChange,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/cmt_manager/change",
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

type AskInnerUserUserCmtManagerGetOrAdd struct {
	Name         string `binding:"required,lte=64" db:"Name"`
	Mobile       string `binding:"required,numeric,lte=11" db:"Mobile"`
	IdCardNumber string `binding:"lte=18" db:"IdCardNumber"`
	IdCardFront  string `binding:"lte=255" db:"IdCardFront"`
	IdCardBack   string `binding:"lte=255" db:"IdCardBack"`
}

func NewAskInnerUserUserCmtManagerGetOrAdd() *AskInnerUserUserCmtManagerGetOrAdd {
	return &AskInnerUserUserCmtManagerGetOrAdd{}
}

type AckInnerUserUserCmtManagerGetOrAdd struct {
	IsNew bool  `db:"IsNew"`
	User  *User `db:"User"`
}

func NewAckInnerUserUserCmtManagerGetOrAdd() *AckInnerUserUserCmtManagerGetOrAdd {
	return &AckInnerUserUserCmtManagerGetOrAdd{
		User: NewUser(),
	}
}

// 添加社区合伙人
func (m *Proxy) InnerUserUserCmtManagerGetOrAdd(ask *AskInnerUserUserCmtManagerGetOrAdd,
) (*AckInnerUserUserCmtManagerGetOrAdd, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerUserUserCmtManagerGetOrAdd
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/cmt_manager/get_or_add",
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

type AskInnerUserUserCmtManagerUnbind struct {
	OldManagerUid string `binding:"required" db:"OldManagerUid"`
}

func NewAskInnerUserUserCmtManagerUnbind() *AskInnerUserUserCmtManagerUnbind {
	return &AskInnerUserUserCmtManagerUnbind{}
}

// 解绑社区合伙人
func (m *Proxy) InnerUserUserCmtManagerUnbind(ask *AskInnerUserUserCmtManagerUnbind,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/user/cmt_manager/unbind",
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

type AskInnerUserWxXcxCmtManagerGetOrAdd struct {
	UserId string `binding:"required" db:"UserId"`
	Token  string `binding:"required" db:"Token"`
	Name   string `binding:"required,lte=64" db:"Name"`
	Mobile string `binding:"required,numeric,lte=11" db:"Mobile"`
}

func NewAskInnerUserWxXcxCmtManagerGetOrAdd() *AskInnerUserWxXcxCmtManagerGetOrAdd {
	return &AskInnerUserWxXcxCmtManagerGetOrAdd{}
}

type AckInnerUserWxXcxCmtManagerGetOrAdd struct {
	IsNew bool  `db:"IsNew"`
	User  *User `db:"User"`
}

func NewAckInnerUserWxXcxCmtManagerGetOrAdd() *AckInnerUserWxXcxCmtManagerGetOrAdd {
	return &AckInnerUserWxXcxCmtManagerGetOrAdd{
		User: NewUser(),
	}
}

// 微信小程序专用内部接口
// 添加社区合伙人
func (m *Proxy) InnerUserWxXcxCmtManagerGetOrAdd(ask *AskInnerUserWxXcxCmtManagerGetOrAdd,
) (*AckInnerUserWxXcxCmtManagerGetOrAdd, error) {
	type Ack struct {
		Code    int
		Message string
		Data    *AckInnerUserWxXcxCmtManagerGetOrAdd
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/wx_xcx/cmt_manager/get_or_add",
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

type AskInnerUserWxXcxRefreshToken struct {
	UserId string `binding:"required" db:"UserId"`
	Token  string `binding:"required" db:"Token"`
}

func NewAskInnerUserWxXcxRefreshToken() *AskInnerUserWxXcxRefreshToken {
	return &AskInnerUserWxXcxRefreshToken{}
}

// 刷新token
func (m *Proxy) InnerUserWxXcxRefreshToken(ask *AskInnerUserWxXcxRefreshToken,
) (map[string]interface{}, error) {
	type Ack struct {
		Code    int
		Message string
		Data    map[string]interface{}
	}
	ack := &Ack{}
	err := m.Invoke(
		"POST",
		"/inner/user/wx_xcx/refresh_token",
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

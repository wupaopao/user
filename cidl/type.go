package cidl

import (
	"business/agency/proxy/agency"

	wx2 "github.com/mz-eco/mz/wx"
)

// 成员信息
type AuthCityStaff struct {
	// 团购组织ID
	OrganizationId uint32

	// 团购组织名称
	OrganizationName string

	// 团购模式
	GroupBuyingMode uint32

	// 角色ID
	RoleId uint32

	// 角色名称
	RoleName string

	// 角色权限
	RoleAuthorization *agency.RoleAuthorizationMap
}

func NewAuthCityStaff() *AuthCityStaff {
	return &AuthCityStaff{}
}

// 微信社群管理员
type AuthWxXcxCommunityManager struct {
	OrganizationId   uint32
	OrganizationName string
	GroupId          uint32
	GroupName        string
}

func NewAuthWxXcxCommunityManager() *AuthWxXcxCommunityManager {
	return &AuthWxXcxCommunityManager{}
}

type AuthWxXcxOrganizationManager struct {
	OrganizationId   uint32
	OrganizationName string
}

func NewAuthWxXcxOrganizationManager() *AuthWxXcxOrganizationManager {
	return &AuthWxXcxOrganizationManager{}
}

type AuthWxXcxOrganization struct {
	OrganizationId uint32
	Name           string
	Logo           string

	// 团购模式
	GroupBuyingMode uint32
}

func NewAuthWxXcxOrganization() *AuthWxXcxOrganization {
	return &AuthWxXcxOrganization{}
}

// admin授权存储结构
type AuthAdmin struct {
	UserId string
	User   *wx2.WxWorkUser
}

func NewAuthAdmin() *AuthAdmin {
	return &AuthAdmin{}
}

// city授权存储结构
type AuthCity struct {
	UserId string         // 用户ID
	User   *User          // 用户信息
	Staff  *AuthCityStaff // 组织成员信息
}

func NewAuthCity() *AuthCity {
	return &AuthCity{}
}

// wx_xcx授权存储结构
type AuthWxXcx struct {
	IsVisitor  bool   // 是否是访客，如果为true，则User字段为空
	UserId     string // 用户ID，如果IsVisitor=true, userId为unionID
	SessionKey *wx2.AckSessionKey
	WxUserInfo *wx2.UserInfo

	User *User

	Organization        *AuthWxXcxOrganization
	CommunityManager    *AuthWxXcxCommunityManager
	OrganizationManager *AuthWxXcxOrganizationManager
}

func NewAuthWxXcx() *AuthWxXcx {
	return &AuthWxXcx{}
}

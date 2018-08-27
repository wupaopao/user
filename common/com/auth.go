package com

import (
	"encoding/json"
	"time"

	"business/agency/proxy/agency"
	"business/auth/common/mq"
	"business/community/proxy/community"
	"business/user/cidl"
	"business/user/common/cache"

	cache2 "github.com/mz-eco/mz/cache"
	"github.com/mz-eco/mz/errors"
	"github.com/mz-eco/mz/log"
	"github.com/mz-eco/mz/utils"
	wx2 "github.com/mz-eco/mz/wx"
)

/**
admin授权信息
*/
type AuthAdmin cidl.AuthAdmin

func NewAuthAdmin(user *wx2.WxWorkUser) *AuthAdmin {
	return &AuthAdmin{
		UserId: user.UserID,
		User:   user,
	}
}

func NewAuthAdminByToken(userId string, token string) (auth *AuthAdmin, err error) {
	adminCache := cache.NewUserCache()
	tokenInfo, err := adminCache.GetAuthAdminToken(token)
	if err != nil {
		auth = nil
		if err != cache2.Nil {
			log.Warnf("get auth admin token info from cache failed. %s", err)
		}
		return
	}

	auth = &AuthAdmin{}
	err = json.Unmarshal([]byte(tokenInfo), auth)
	if err != nil {
		auth = nil
		log.Warnf("unmarshal auth admin token failed. %s", err)
		return
	}

	if auth.UserId != userId {
		err = errors.New("user id does not match auth")
		return
	}

	return
}

func DeleteAuthAdminByToken(token string) (err error) {
	adminCache := cache.NewUserCache()
	err = adminCache.DeleteAuthAdminToken(token)
	if err != nil {
		log.Warnf("delete auth admin token failed. %s", err)
		return
	}

	err = mq.BroadcastInvalidateAuthAdminToken(token)
	if err != nil {
		log.Warnf("broadcast invalidate auth admin token failed. %s", err)
		return
	}

	return
}

const TTL_AUTH_ADMIN_TOKEN = time.Duration(24) * time.Hour // 24h

/**
创建新的Token
*/
func (m *AuthAdmin) NewToken() (token string, err error) {
	token = utils.UniqueID()
	adminCache := cache.NewUserCache()
	byteAuth, err := json.Marshal(m)
	if err != nil {
		log.Warnf("marshal wx work user failed. %s", err)
		return
	}

	err = adminCache.SetAuthAdminToken(token, byteAuth, TTL_AUTH_ADMIN_TOKEN)
	if err != nil {
		log.Warnf("set auth token info cache failed. %s", err)
		return
	}

	return
}

/**
城市合伙人系统授权
*/
type AuthCity cidl.AuthCity

func NewAuthCityByToken(userId string, token string) (auth *AuthCity, err error) {
	userCache := cache.NewUserCache()
	tokenInfo, err := userCache.Cache.Get(userCache.KeyAuthCityToken(token))
	if err != nil {
		if err != cache2.Nil {
			log.Warnf("get cit auth info failed. %s", err)
		}
		return
	}

	auth = &AuthCity{}
	err = json.Unmarshal([]byte(tokenInfo), auth)
	if err != nil {
		auth = nil
		log.Warnf("unmarshal auth city token failed. %s", err)
		return
	}

	return
}

func DeleteAuthCityByToken(token string) (err error) {
	userCache := cache.NewUserCache()
	_, err = userCache.Cache.Delete(userCache.KeyAuthCityToken(token))
	if err != nil {
		log.Warnf("delete auth city token failed. %s", err)
		return
	}

	err = mq.BroadcastInvalidateAuthOrgToken(token)
	if err != nil {
		log.Warnf("broadcast invalidate auth org token failed. %s", err)
		return
	}

	return
}

const TTL_AUTH_CITY_TOKEN = time.Duration(24) * time.Hour // 24h

func (m *AuthCity) NewToken() (token string, err error) {
	token = utils.UniqueID()
	err = m.SetToken(token)
	return
}

func (m *AuthCity) SetToken(token string) (err error) {
	userCache := cache.NewUserCache()
	byteAuth, err := json.Marshal(m)
	if err != nil {
		log.Warnf("marshal auth city failed. %s", err)
		return
	}

	err = userCache.Cache.Set(userCache.KeyAuthCityToken(token), byteAuth, TTL_AUTH_CITY_TOKEN)
	if err != nil {
		log.Warnf("set auth city token cache failed. %s", err)
		return
	}

	return
}

/**
小程序授权
*/
type AuthWxXcx cidl.AuthWxXcx

func NewAuthWxXcxByToken(userId string, token string) (auth *AuthWxXcx, err error) {
	userCache := cache.NewUserCache()
	tokenInfo, err := userCache.Cache.Get(userCache.KeyAuthWxXcxToken(token))
	if err != nil {
		if err != cache2.Nil {
			log.Warnf("get auth wx xcx token info from cache failed. %s", err)
		}
		return
	}

	auth = &AuthWxXcx{}
	err = json.Unmarshal([]byte(tokenInfo), auth)
	if err != nil {
		auth = nil
		log.Warnf("unmarshal auth wx xcx token failed. %s", err)
		return
	}

	if auth.UserId != userId {
		err = errors.New("user id does not match auth")
		return
	}

	return
}

func DeleteAuthWxXcxByToken(token string) (err error) {
	userCache := cache.NewUserCache()
	_, err = userCache.Cache.Delete(userCache.KeyAuthWxXcxToken(token))
	if err != nil {
		log.Warnf("delete auth wx xcx token failed. %s", err)
		return
	}

	err = mq.BroadcastInvalidateAuthWxXcxToken(token)
	if err != nil {
		log.Warnf("broadcast invalidate auth wx_xcx token failed. %s", err)
		return
	}

	return
}

const TTL_AUTH_WX_XCX_TOKEN = time.Duration(24) * time.Hour // 24h

func (m *AuthWxXcx) NewToken() (token string, err error) {
	token = utils.UniqueID()
	err = m.SetToken(token)
	return
}

func (m *AuthWxXcx) SetToken(token string) (err error) {
	userCache := cache.NewUserCache()
	byteAuth, err := json.Marshal(m)
	if err != nil {
		log.Warnf("marshal auth wx xcx failed. %s", err)
		return
	}

	err = userCache.Cache.Set(userCache.KeyAuthWxXcxToken(token), byteAuth, TTL_AUTH_WX_XCX_TOKEN)
	if err != nil {
		log.Warnf("set auth wx xcx token cache failed. %s", err)
		return
	}

	return
}

func (m *AuthWxXcx) SetUserInfo(user *cidl.User) (err error) {
	m.User = user

	//是否被禁用
	/*info, errProxy := agency.NewProxy("agency-service").InnerAgencyStaffIsDisableByUserID(user.UserID)	
	if errProxy != nil {
		err = errProxy
		log.Warnf("get agency staff is_disable from proxy failed. %s", err)
		return
	}
	if info.IsDisable {
		err = errors.New("account is disable.") 
		log.Warnf("set user info failed. %s", err)
		return
	}
	*/

	// 是否是社区合伙人
	if user.IsCmtManager {
		group, errProxy := community.NewProxy("community-service").InnerCommunityGroupInfoByUserIDByUserID(user.UserID)
		if errProxy != nil {
			err = errProxy
			log.Warnf("get community group from proxy failed. %s", err)
			return
		}

		m.CommunityManager = &cidl.AuthWxXcxCommunityManager{
			OrganizationId:   group.OrganizationId,
			OrganizationName: group.OrganizationName,
			GroupId:          group.GroupId,
			GroupName:        group.Name,
		}

		organization, errProxy := agency.NewProxy("agency-service").InnerAgencyOrganizationInfoByOrganizationID(group.OrganizationId)
		if errProxy != nil {
			err = errProxy
			log.Warnf("get organization from proxy failed. %s", err)
			return
		}

		m.Organization = &cidl.AuthWxXcxOrganization{
			OrganizationId:  organization.OrganizationId,
			Name:            organization.Name,
			Logo:            organization.Logo,
			GroupBuyingMode: uint32(organization.GroupBuyingMode),
		}

	} else {
		// 是否是城市合伙人
		if user.IsOrgManager {
			organization, errProxy := agency.NewProxy("agency-service").InnerAgencyOrganizationInfoByUserIDByUserID(user.UserID)
			if errProxy != nil {
				err = errProxy
				log.Warnf("get organization from proxy failed. %s", err)
				return
			}

			m.OrganizationManager = &cidl.AuthWxXcxOrganizationManager{
				OrganizationId:   organization.OrganizationId,
				OrganizationName: organization.Name,
			}

			m.Organization = &cidl.AuthWxXcxOrganization{
				OrganizationId:  organization.OrganizationId,
				Name:            organization.Name,
				Logo:            organization.Logo,
				GroupBuyingMode: uint32(organization.GroupBuyingMode),
			}

		} else if user.IsOrgStaff {
			// ...预留
		}

	}

	return
}

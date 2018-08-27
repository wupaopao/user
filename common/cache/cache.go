package cache

import (
	"fmt"
	"time"

	"github.com/mz-eco/mz/cache"
)

type UserCache struct {
	Cache *cache.Cache
}

func NewUserCache() *UserCache {
	return &UserCache{
		Cache: cache.NewRedis("user"),
	}
}

// 接口授权token
// 微信小程序TOKEN
func (m *UserCache) KeyAuthWxXcxToken(token string) string {
	return fmt.Sprintf("usr:auth:tkn:wx_xcx:%s", token)
}

// 城市合伙人端TOKEN
func (m *UserCache) KeyAuthCityToken(token string) string {
	return fmt.Sprintf("usr:auth:tkn:city:%s", token)
}

func (m *UserCache) KeyAuthAdminToken(token string) string {
	return fmt.Sprintf("usr:auth:tkn:admin:%s", token)
}

func (m *UserCache) GetAuthAdminToken(token string) (string, error) {
	return m.Cache.Get(m.KeyAuthAdminToken(token))
}

func (m *UserCache) SetAuthAdminToken(token string, content []byte, ttl time.Duration) error {
	return m.Cache.Set(m.KeyAuthAdminToken(token), content, ttl)
}

func (m *UserCache) DeleteAuthAdminToken(token string) (err error) {
	_, err = m.Cache.Delete(m.KeyAuthAdminToken(token))
	return
}

// 验证码
func (m *UserCache) KeyMobileVerifyCityLogin(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:city:login:%s", mobile)
}

func (m *UserCache) KeyMobileVerifyCityLoginInterval(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:city:lgnitv:%s", mobile)
}

func (m *UserCache) KeyMobileVerifyWxXcxBindMobile(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:wxxcx:bindmob:%s", mobile)
}

func (m *UserCache) KeyMobileVerifyWxXcxBindMobileInterval(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:wxxcx:bindmobitv:%s", mobile)
}

func (m *UserCache) KeyMobileVerifyWxXcxModifyMobile(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:wxxcx:mdfmob:%s", mobile)
}

func (m *UserCache) KeyMobileVerifyWxXcxModifyMobileInterval(mobile string) string {
	return fmt.Sprintf("usr:mobvrf:wxxcx:mdfitv:%s", mobile)
}

// 企业微信access token
func (m *UserCache) keyWxWorkAccessToken() string {
	return "usr:wxw:tkn"
}

func (m *UserCache) GetWxWorkAccessToken() (accessToken string, err error) {
	key := m.keyWxWorkAccessToken()
	accessToken, err = m.Cache.Get(key)
	return
}

func (m *UserCache) SetWxWorkAccessToken(accessToken []byte, ttl time.Duration) (err error) {
	key := m.keyWxWorkAccessToken()
	return m.Cache.Set(key, accessToken, ttl)
}

// 企业微信授权session
func (m *UserCache) keyWxWorkAuthSession(sessionId string) string {
	return fmt.Sprintf("usr:wxw:auth:ses:%s", sessionId)
}

func (m *UserCache) GetWxWorkAuthSession(sessionId string) (session string, err error) {
	key := m.keyWxWorkAuthSession(sessionId)
	return m.Cache.Get(key)
}

func (m *UserCache) SetWxWorkAuthSession(sessionId string, session []byte, ttl time.Duration) (err error) {
	key := m.keyWxWorkAuthSession(sessionId)
	return m.Cache.Set(key, session, ttl)
}

func (m *UserCache) DeleteWxWorkAuthSession(sessionId string) (err error) {
	key := m.keyWxWorkAuthSession(sessionId)
	_, err = m.Cache.Delete(key)
	return
}

// 企业微信用户缓存
const TTL_WxWorkUser = time.Duration(24) * time.Hour

func (m *UserCache) KeyWxWorkUser(userID string) string {
	return fmt.Sprintf("usr:wxw:usr:%s", userID)
}

func (m *UserCache) GetWxWorkUser(userID string) (string, error) {
	return m.Cache.Get(m.KeyWxWorkUser(userID))
}

func (m *UserCache) SetWxWorkUser(userID string, user []byte) (err error) {
	return m.Cache.Set(m.KeyWxWorkUser(userID), user, TTL_WxWorkUser)
}

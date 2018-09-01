package cidl

import "fmt"

type autoErrorserrorsTcidl int

const (
	ErrGenerateAuthWxXcxTokenFailed autoErrorserrorsTcidl = 3001 //生成微信小程序用户TOKEN失败
	ErrUserWasBound                 autoErrorserrorsTcidl = 3002 //用户已经被绑定
	ErrWxXcxEmptyUnionId            autoErrorserrorsTcidl = 3102 //获取UnionId为空
	ErrWxXcxIllegalToken            autoErrorserrorsTcidl = 3103 //无效token
	ErrWxXcxUpdateTokenFailed       autoErrorserrorsTcidl = 3104 //更新token失败
	ErrWxXcxUserExists              autoErrorserrorsTcidl = 3105 //微信或者手机绑定的用户已经存在
	ErrWxXcxDecryptUserInfoFailed   autoErrorserrorsTcidl = 3106 //解码用户信息失败
	ErrWxXcxMobileNotExist          autoErrorserrorsTcidl = 3107 //手机用户未存在
	ErrWxXcxMobileHasBindWxUnionId  autoErrorserrorsTcidl = 3108 //手机已经绑定其他微信
	ErrWxXcxMobileNotChange         autoErrorserrorsTcidl = 3109 //手机未发生更改
	ErrOrgIllegalToken              autoErrorserrorsTcidl = 3200 //无效City授权Token
	ErrOrgGenerateTokenFailed       autoErrorserrorsTcidl = 3201 //生成团购组织端TOKEN失败
	ErrOrgLoginUserNotExist         autoErrorserrorsTcidl = 3202 //用户不存在或无权进入
	ErrLoginUserForbidden           autoErrorserrorsTcidl = 3203 //账号被禁用
	ErrAdminIllegalToken            autoErrorserrorsTcidl = 3300 //无效Admin授权Token
)

func (m autoErrorserrorsTcidl) Number() int { return int(m) }
func (m autoErrorserrorsTcidl) Message() string {
	switch m {

	case ErrGenerateAuthWxXcxTokenFailed:
		return "生成微信小程序用户TOKEN失败"
	case ErrUserWasBound:
		return "用户已经被绑定"
	case ErrWxXcxEmptyUnionId:
		return "获取UnionId为空"
	case ErrWxXcxIllegalToken:
		return "无效token"
	case ErrWxXcxUpdateTokenFailed:
		return "更新token失败"
	case ErrWxXcxUserExists:
		return "微信或者手机绑定的用户已经存在"
	case ErrWxXcxDecryptUserInfoFailed:
		return "解码用户信息失败"
	case ErrWxXcxMobileNotExist:
		return "手机用户未存在"
	case ErrWxXcxMobileHasBindWxUnionId:
		return "手机已经绑定其他微信"
	case ErrWxXcxMobileNotChange:
		return "手机未发生更改"
	case ErrOrgIllegalToken:
		return "无效City授权Token"
	case ErrOrgGenerateTokenFailed:
		return "生成团购组织端TOKEN失败"
	case ErrOrgLoginUserNotExist:
		return "用户不存在或无权进入"
	case ErrLoginUserForbidden:
		return "账号被禁用"
	case ErrAdminIllegalToken:
		return "无效Admin授权Token"
	default:
		return "UNKNOWN_MESSAGE_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) Name() string {
	switch m {

	case ErrGenerateAuthWxXcxTokenFailed:
		return "ErrGenerateAuthWxXcxTokenFailed"
	case ErrUserWasBound:
		return "ErrUserWasBound"
	case ErrWxXcxEmptyUnionId:
		return "ErrWxXcxEmptyUnionId"
	case ErrWxXcxIllegalToken:
		return "ErrWxXcxIllegalToken"
	case ErrWxXcxUpdateTokenFailed:
		return "ErrWxXcxUpdateTokenFailed"
	case ErrWxXcxUserExists:
		return "ErrWxXcxUserExists"
	case ErrWxXcxDecryptUserInfoFailed:
		return "ErrWxXcxDecryptUserInfoFailed"
	case ErrWxXcxMobileNotExist:
		return "ErrWxXcxMobileNotExist"
	case ErrWxXcxMobileHasBindWxUnionId:
		return "ErrWxXcxMobileHasBindWxUnionId"
	case ErrWxXcxMobileNotChange:
		return "ErrWxXcxMobileNotChange"
	case ErrOrgIllegalToken:
		return "ErrOrgIllegalToken"
	case ErrOrgGenerateTokenFailed:
		return "ErrOrgGenerateTokenFailed"
	case ErrOrgLoginUserNotExist:
		return "ErrOrgLoginUserNotExist"
	case ErrLoginUserForbidden:
		return "ErrLoginUserForbidden"
	case ErrAdminIllegalToken:
		return "ErrAdminIllegalToken"
	default:
		return "UNKNOWN_Name_autoErrorserrorsTcidl"
	}
}
func (m autoErrorserrorsTcidl) String() string {
	return fmt.Sprintf("[%d:%s]%s", m, m.Name(), m.Message())

}

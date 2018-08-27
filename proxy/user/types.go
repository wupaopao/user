package user

import "business/user/cidl"

type AuthCityStaff cidl.AuthCityStaff

func NewAuthCityStaff() *AuthCityStaff {
	return &AuthCityStaff{}
}

type AuthWxXcxOrganization = cidl.AuthWxXcxOrganization

func NewAuthWxXcxOrganization() *AuthWxXcxOrganization {
	return &AuthWxXcxOrganization{}
}

type AuthWxXcxCommunityManager = cidl.AuthWxXcxCommunityManager

func NewAuthWxXcxCommunityManager() *AuthWxXcxCommunityManager {
	return &AuthWxXcxCommunityManager{}
}

type AuthAdmin = cidl.AuthAdmin
type AuthCity = cidl.AuthCity
type AuthWxXcx = cidl.AuthWxXcx

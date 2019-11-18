package service

import "short/app/entity"

var _ SSOAccount = (*SSOAccountFake)(nil)

type SSOAccountFake struct {
	user entity.SSOUser
}

func (s SSOAccountFake) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	return s.user, nil
}

func NewSSOAccountFake(user entity.SSOUser) SSOAccountFake {
	return SSOAccountFake{
		user: user,
	}
}

package service

import "short/app/entity"

var _ SSOAccount = (*SSOAccountFake)(nil)

// SSOAccountFake represents in memory implementation of account service that
// access account data from the identity provider.
type SSOAccountFake struct {
	user entity.SSOUser
}

// GetSingleSignOnUser retrieves user information from identity provider using
// access token.
func (s SSOAccountFake) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	return s.user, nil
}

// NewSSOAccountFake creates fake SSOAccount service.
func NewSSOAccountFake(user entity.SSOUser) SSOAccountFake {
	return SSOAccountFake{
		user: user,
	}
}

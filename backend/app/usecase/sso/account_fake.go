package sso

import "github.com/short-d/short/backend/app/entity"

var _ Account = (*AccountFake)(nil)

// AccountFake represents in memory implementation of account service that
// access account data from the identity provider.
type AccountFake struct {
	user entity.SSOUser
}

// GetSingleSignOnUser retrieves user information from identity provider using
// access token.
func (a AccountFake) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	return a.user, nil
}

// NewAccountFake creates fake Account service.
func NewAccountFake(user entity.SSOUser) AccountFake {
	return AccountFake{
		user: user,
	}
}

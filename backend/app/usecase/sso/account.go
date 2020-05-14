package sso

import "github.com/short-d/short/backend/app/entity"

// Account accesses account data from the identity provider.
type Account interface {
	GetSingleSignOnUser(accessToken string) (entity.SSOUser, error)
}

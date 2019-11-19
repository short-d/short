package service

import "short/app/entity"

// SSOAccount accesses account data from the identity provider.
type SSOAccount interface {
	GetSingleSignOnUser(accessToken string) (entity.SSOUser, error)
}

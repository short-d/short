package repository

import "github.com/short-d/short/app/entity"

// AccountMapping accesses account mapping between SSOUser and internal User
// from storage media, such as database.
type AccountMapping interface {
	IsSSOUserExist(ssoUser entity.SSOUser) (bool, error)
	CreateMapping(ssoUser entity.SSOUser, user entity.User) error
}

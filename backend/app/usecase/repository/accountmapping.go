package repository

import "short/app/entity"

type AccountMapping interface {
	IsSSOUserExist(ssoUser entity.SSOUser) (bool, error)
	CreateMapping(ssoUser entity.SSOUser, user entity.User) error
}

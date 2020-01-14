package repository

import (
	"errors"

	"github.com/short-d/short/app/entity"
)

var _ AccountMapping = (*AccountMappingFake)(nil)

// AccountMappingFake represents in memory implementation of AccountMapping
// repository.
type AccountMappingFake struct {
	ssoUsers []entity.SSOUser
	users    []entity.User
}

// IsSSOUserExist checks whether a external user is linked to any internal
// user.
func (a AccountMappingFake) IsSSOUserExist(ssoUser entity.SSOUser) (bool, error) {
	for _, currSSOUser := range a.ssoUsers {
		if currSSOUser.ID == ssoUser.ID {
			return true, nil
		}
	}
	return false, nil
}

// IsRelationExist checks whether a given external user is linked to a given
// internal user.
func (a AccountMappingFake) IsRelationExist(ssoUser entity.SSOUser, user entity.User) bool {
	for idx, currSSOUser := range a.ssoUsers {
		if currSSOUser.ID != ssoUser.ID {
			continue
		}

		if a.users[idx].ID == user.ID {
			return true
		}
	}
	return false
}

// CreateMapping links an external user with an internal user.
func (a *AccountMappingFake) CreateMapping(ssoUser entity.SSOUser, user entity.User) error {
	isExist := a.IsRelationExist(ssoUser, user)
	if isExist {
		return errors.New("mapping exists")
	}
	a.users = append(a.users, user)
	a.ssoUsers = append(a.ssoUsers, ssoUser)
	return nil
}

// NewAccountMappingFake creates in memory implementation of AccountMapping
// repository.
func NewAccountMappingFake(
	ssoUsers []entity.SSOUser,
	users []entity.User) (AccountMappingFake, error) {
	if len(ssoUsers) != len(users) {
		return AccountMappingFake{}, errors.New("account length does not match")
	}
	return AccountMappingFake{
		ssoUsers: ssoUsers,
		users:    users,
	}, nil
}

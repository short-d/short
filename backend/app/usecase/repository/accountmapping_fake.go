package repository

import (
	"errors"
	"short/app/entity"
)

var _ AccountMapping = (*AccountMappingFake)(nil)

type AccountMappingFake struct {
	ssoUsers []entity.SSOUser
	users    []entity.User
}

func (a AccountMappingFake) IsSSOUserExist(ssoUser entity.SSOUser) (bool, error) {
	for _, currSSOUser := range a.ssoUsers {
		if currSSOUser.ID == ssoUser.ID {
			return true, nil
		}
	}
	return false, nil
}

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

func (a *AccountMappingFake) CreateMapping(ssoUser entity.SSOUser, user entity.User) error {
	isExist := a.IsRelationExist(ssoUser, user)
	if isExist {
		return errors.New("mapping exists")
	}
	a.users = append(a.users, user)
	a.ssoUsers = append(a.ssoUsers, ssoUser)
	return nil
}

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

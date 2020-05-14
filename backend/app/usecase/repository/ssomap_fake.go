package repository

import (
	"errors"

	"github.com/short-d/short/backend/app/entity"
)

var _ SSOMap = (*SSOMapFake)(nil)

// SSOMapFake represents in memory implementation of SSOMap repository.
type SSOMapFake struct {
	ssoUsers []entity.SSOUser
	users    []entity.User
}

func (s SSOMapFake) GetShortUserID(ssoUserID string) (string, error) {
	for idx, currSSOUser := range s.ssoUsers {
		if currSSOUser.ID == ssoUserID {
			return s.users[idx].ID, nil
		}
	}
	return "", ErrEntryNotFound("user not found")
}

// IsSSOUserExist checks whether a external user is linked to any internal
// user.
func (s SSOMapFake) IsSSOUserExist(ssoUserID string) (bool, error) {
	for _, currSSOUser := range s.ssoUsers {
		if currSSOUser.ID == ssoUserID {
			return true, nil
		}
	}
	return false, nil
}

// IsRelationExist checks whether a given external user is linked to a given
// internal user.
func (s SSOMapFake) IsRelationExist(ssoUserID string, userID string) bool {
	for idx, currSSOUser := range s.ssoUsers {
		if currSSOUser.ID != ssoUserID {
			continue
		}

		if s.users[idx].ID == userID {
			return true
		}
	}
	return false
}

// CreateMapping links an external user with an internal user.
func (s *SSOMapFake) CreateMapping(ssoUserID string, userID string) error {
	isExist := s.IsRelationExist(ssoUserID, userID)
	if isExist {
		return nil
	}
	s.users = append(s.users, entity.User{ID: userID})
	s.ssoUsers = append(s.ssoUsers, entity.SSOUser{ID: ssoUserID})
	return nil
}

// NewsSSOMapFake creates in memory implementation of SSOMapFake repository.
func NewsSSOMapFake(
	ssoUsers []entity.SSOUser,
	users []entity.User) (SSOMapFake, error) {
	if len(ssoUsers) != len(users) {
		return SSOMapFake{}, errors.New("account length does not match")
	}
	return SSOMapFake{
		ssoUsers: ssoUsers,
		users:    users,
	}, nil
}

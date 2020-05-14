package repository

import (
	"errors"
	"fmt"
)

var _ SSOMap = (*SSOMapFake)(nil)

// SSOMapFake represents in memory implementation of SSOMap repository.
type SSOMapFake struct {
	ssoUserIDs []string
	userIDs    []string
}

// GetShortUserID retrieves the internal user ID that is linked to the external
// user.
func (s SSOMapFake) GetShortUserID(ssoUserID string) (string, error) {
	for idx, currSSOUserID := range s.ssoUserIDs {
		fmt.Println(currSSOUserID)
		if currSSOUserID == ssoUserID {
			return s.userIDs[idx], nil
		}
	}
	return "", ErrEntryNotFound("user ID not found")
}

// IsSSOUserExist checks whether a external user is linked to any internal
// user.
func (s SSOMapFake) IsSSOUserExist(ssoUserID string) (bool, error) {
	for _, currSSOUserID := range s.ssoUserIDs {
		if currSSOUserID == ssoUserID {
			return true, nil
		}
	}
	return false, nil
}

// IsRelationExist checks whether a given external user is linked to a given
// internal user.
func (s SSOMapFake) IsRelationExist(ssoUserID string, userID string) bool {
	for idx, currSSOUserID := range s.ssoUserIDs {
		if currSSOUserID != ssoUserID {
			continue
		}

		if s.userIDs[idx] == userID {
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
	s.userIDs = append(s.userIDs, userID)
	s.ssoUserIDs = append(s.ssoUserIDs, ssoUserID)
	return nil
}

// NewsSSOMapFake creates in memory implementation of SSOMapFake repository.
func NewsSSOMapFake(
	ssoUserIDs []string,
	userIDs []string) (SSOMapFake, error) {
	if len(ssoUserIDs) != len(userIDs) {
		return SSOMapFake{}, errors.New("account length does not match")
	}
	return SSOMapFake{
		ssoUserIDs: ssoUserIDs,
		userIDs:    userIDs,
	}, nil
}

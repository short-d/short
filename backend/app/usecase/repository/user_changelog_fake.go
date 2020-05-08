package repository

import (
	"errors"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ UserChangeLog = (*UserChangeLogFake)(nil)

// UserChangeLogFake represents in memory implementation of UserChangeLog repository.
type UserChangeLogFake struct {
	lastViewedAt map[string]time.Time
}

// GetLastViewedAt retrieves lastViewedAt for user
func (u UserChangeLogFake) GetLastViewedAt(user entity.User) (time.Time, error) {
	if lastViewedAt, exists := u.lastViewedAt[user.Email]; exists {
		return lastViewedAt.UTC(), nil
	}

	return time.Time{}, ErrEntryNotFound("user not found")
}

// UpdateLastViewedAt updates lastViewedAt for user to currentTime
func (u *UserChangeLogFake) UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error) {
	if _, exists := u.lastViewedAt[user.Email]; exists {
		u.lastViewedAt[user.Email] = currentTime.UTC()
		return currentTime.UTC(), nil
	}

	return time.Time{}, ErrEntryNotFound("user not found")
}

// CreateRelation inserts new entry into UserChangeLog repository
func (u *UserChangeLogFake) CreateRelation(user entity.User, currentTime time.Time) error {
	if _, exists := u.lastViewedAt[user.Email]; exists {
		return errors.New("user exists")
	}

	u.lastViewedAt[user.Email] = currentTime.UTC()
	return nil
}

// NewUserChangeLogFake creates in memory UserChangeLog repository
func NewUserChangeLogFake(lastViewedAt map[string]time.Time) UserChangeLogFake {
	return UserChangeLogFake{
		lastViewedAt: lastViewedAt,
	}
}

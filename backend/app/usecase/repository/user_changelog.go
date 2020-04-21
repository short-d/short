package repository

import (
	"time"

	"github.com/short-d/short/app/entity"
)

// ErrEntryNotFound represents entry unavailable error
type ErrEntryNotFound string

func (e ErrEntryNotFound) Error() string {
	return string(e)
}

// UserChangeLog accesses user-changelog information from storage, such as database.
type UserChangeLog interface {
	GetLastViewedAt(user entity.User) (time.Time, error)
	UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error)
	CreateRelation(user entity.User, currentTime time.Time) error
}

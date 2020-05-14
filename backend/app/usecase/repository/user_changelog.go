package repository

import (
	"time"

	"github.com/short-d/short/backend/app/entity"
)

// UserChangeLog accesses user-changelog information from storage, such as database.
type UserChangeLog interface {
	GetLastViewedAt(user entity.User) (time.Time, error)
	UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error)
	CreateRelation(user entity.User, currentTime time.Time) error
}

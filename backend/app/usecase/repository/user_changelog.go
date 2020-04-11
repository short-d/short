package repository

import (
	"time"

	"github.com/short-d/short/app/entity"
)

type UserChangeLog interface {
	GetLastViewedAt(user entity.User) (time.Time, error)
	UpdateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error)
	CreateLastViewedAt(user entity.User, currentTime time.Time) (time.Time, error)
}

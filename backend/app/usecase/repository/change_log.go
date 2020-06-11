package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// ChangeLog accesses changelog from storage, such as database.
type ChangeLog interface {
	GetChangeLog() ([]entity.Change, error)
	CreateChange(newChange entity.Change) (entity.Change, error)
	DeleteChange(id string) error
	UpdateChange(newChange entity.Change) (entity.Change, error)
}

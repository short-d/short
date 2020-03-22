package repository

import (
	"github.com/short-d/short/app/entity"
)

// Change accesses changelog information from storage, such as database.
type Changelog interface {
	GetChangeLog() ([]entity.Change, error)
	CreateChange(newChange entity.Change) (entity.Change, error)
}

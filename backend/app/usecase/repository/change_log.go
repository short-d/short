package repository

import (
	"github.com/short-d/short/app/entity"
)

type ChangeLog interface {
	GetChangeLog() ([]entity.Change, error)
	CreateChange(newChange entity.Change) (entity.Change, error)
}

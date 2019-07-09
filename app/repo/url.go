package repo

import (
	"tinyURL/app/entity"
)

type Url interface {
	GetByAlias(alias string) (entity.Url, error)
}
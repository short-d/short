package repo

import "short/app/entity"

type Url interface {
	IsAliasExist(alias string) bool
	GetByAlias(alias string) (entity.Url, error)
	Create(url entity.Url) error
}

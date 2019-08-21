package repo

import "short/app/entity"

type Url interface {
	IsAliasExist(alias string) (bool, error)
	GetByAlias(alias string) (entity.Url, error)
	Create(url entity.Url) error
}

package repo

import "short/app/entity"

type URL interface {
	IsAliasExist(alias string) (bool, error)
	GetByAlias(alias string) (entity.URL, error)
	Create(url entity.URL) error
}

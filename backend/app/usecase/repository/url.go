package repository

import "short/app/entity"

// URL accesses urls from storage, such as database.
type URL interface {
	IsAliasExist(alias string) (bool, error)
	GetByAlias(alias string) (entity.URL, error)
	Create(url entity.URL) error
	Update(url entity.URL) error
}

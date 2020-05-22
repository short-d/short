package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// ShortLink accesses urls from storage, such as database.
type URL interface {
	IsAliasExist(alias string) (bool, error)
	GetByAlias(alias string) (entity.ShortLink, error)
	// TODO(issue#698): change to CreateURL
	Create(url entity.ShortLink) error
	UpdateURL(oldAlias string, newURL entity.ShortLink) (entity.ShortLink, error)
	GetByAliases(aliases []string) ([]entity.ShortLink, error)
}

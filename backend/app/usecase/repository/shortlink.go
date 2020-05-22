package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// ShortLink accesses shortLinks from storage, such as database.
type ShortLink interface {
	IsAliasExist(alias string) (bool, error)
	GetByAlias(alias string) (entity.ShortLink, error)
	CreateShortLink(shortLink entity.ShortLink) error
	UpdateShortLink(oldAlias string, shortLink entity.ShortLink) (entity.ShortLink, error)
	GetByAliases(aliases []string) ([]entity.ShortLink, error)
}

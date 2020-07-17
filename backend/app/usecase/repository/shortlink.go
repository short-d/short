package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// ShortLink accesses shortLinks from storage, such as database.
type ShortLink interface {
	IsAliasExist(alias string) (bool, error)
	GetShortLinkByAlias(alias string) (entity.ShortLink, error)
	CreateShortLink(shortLinkInput entity.ShortLinkInput) error
	UpdateShortLink(oldAlias string, shortLinkInput entity.ShortLinkInput) (entity.ShortLink, error)
	GetShortLinksByAliases(aliases []string) ([]entity.ShortLink, error)
}

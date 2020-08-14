package repository

import (
	"fmt"

	"github.com/short-d/short/backend/app/entity"
)

var _ error = (*ErrAliasNotFound)(nil)

// ErrAliasNotFound represents no short link entry found
// with the given alias.
type ErrAliasNotFound struct {
	Alias string
}

func (e ErrAliasNotFound) Error() string {
	return fmt.Sprintf("short link with alias(%s) not found", e.Alias)
}

// ShortLink accesses shortLinks from storage, such as database.
type ShortLink interface {
	IsAliasExist(alias string) (bool, error)
	GetShortLinkByAlias(alias string) (entity.ShortLink, error)
	CreateShortLink(shortLinkInput entity.ShortLinkInput) error
	UpdateShortLink(oldAlias string, shortLinkInput entity.ShortLinkInput) (entity.ShortLink, error)
	DeleteShortLink(alias string) error
	GetShortLinksByAliases(aliases []string) ([]entity.ShortLink, error)
}

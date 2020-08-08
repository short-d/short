package repository

import (
	"fmt"

	"github.com/short-d/short/backend/app/entity"
)

var _ error = (*ErrShortLinkDeletionFailure)(nil)

type ErrShortLinkDeletionFailure struct {
	Alias string
}

func (e ErrShortLinkDeletionFailure) Error() string {
	return fmt.Sprintf("failed to delete user short link alias(%s)", e.Alias)
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

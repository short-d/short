package resolver

import (
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

// ShortLink retrieves requested fields of ShortLink entity.
type ShortLink struct {
	shortLink entity.ShortLink
}

// Alias retrieves the alias of ShortLink entity.
func (s ShortLink) Alias() *string {
	return &s.shortLink.Alias
}

// LongLink retrieves the long link of ShortLink entity.
func (s ShortLink) LongLink() *string {
	return &s.shortLink.LongLink
}

// ExpireAt retrieves the expiration time of ShortLink entity.
func (s ShortLink) ExpireAt() *scalar.Time {
	if s.shortLink.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *s.shortLink.ExpireAt}
}

func newShortLink(shortLink entity.ShortLink) ShortLink {
	return ShortLink{shortLink: shortLink}
}

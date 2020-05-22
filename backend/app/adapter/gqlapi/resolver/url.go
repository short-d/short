package resolver

import (
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

// ShortLink retrieves requested fields of ShortLink entity.
type URL struct {
	url entity.ShortLink
}

// Alias retrieves the alias of ShortLink entity.
func (u URL) Alias() *string {
	return &u.url.Alias
}

// LongLink retrieves the long link of ShortLink entity.
func (u URL) OriginalURL() *string {
	return &u.url.LongLink
}

// ExpireAt retrieves the expiration time of ShortLink entity.
func (u URL) ExpireAt() *scalar.Time {
	if u.url.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *u.url.ExpireAt}
}

func newURL(url entity.ShortLink) URL {
	return URL{url: url}
}

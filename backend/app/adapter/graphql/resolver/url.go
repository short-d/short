package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
)

// URL retrieves queried fields of URL entity.
type URL struct {
	url entity.URL
}

// Alias retrieves the alias of URL entity.
func (u URL) Alias() *string {
	return &u.url.Alias
}

// OriginalURL retrieves the long link of URL entity.
func (u URL) OriginalURL() *string {
	return &u.url.OriginalURL
}

// ExpireAt retrieves the expiration time of URL entity.
func (u URL) ExpireAt() *scalar.Time {
	if u.url.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *u.url.ExpireAt}
}

func newURL(url entity.URL) URL {
	return URL{url: url}
}

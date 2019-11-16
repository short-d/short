package resolver

import (
	"short/app/adapter/graphql/scalar"
	"short/app/entity"
)

type URL struct {
	url entity.URL
}

func (u URL) Alias() *string {
	return &u.url.Alias
}

func (u URL) OriginalURL() *string {
	return &u.url.OriginalURL
}

func (u URL) ExpireAt() *scalar.Time {
	if u.url.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *u.url.ExpireAt}
}

func newURL(url entity.URL) URL {
	return URL{url: url}
}

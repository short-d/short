package resolver

import (
	"tinyURL/app/entity"
	"tinyURL/app/graphql/scalar"
)

type Url struct {
	url entity.Url
}

func (u Url) Alias() *string {
	return &u.url.Alias
}

func (u Url) OriginalUrl() *string {
	return &u.url.OriginalUrl
}

func (u Url) ExpireAt() *scalar.Time {
	if u.url.ExpireAt == nil {
		return nil
	}

	return &scalar.Time{Time: *u.url.ExpireAt}
}

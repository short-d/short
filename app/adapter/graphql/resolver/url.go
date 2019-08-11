package resolver

import (
	scalar2 "short/app/adapter/graphql/scalar"
	"short/app/entity"
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

func (u Url) ExpireAt() *scalar2.Time {
	if u.url.ExpireAt == nil {
		return nil
	}

	return &scalar2.Time{Time: *u.url.ExpireAt}
}

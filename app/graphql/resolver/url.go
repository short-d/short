package resolver

import "tinyURL/app/entity"

type Url struct {
	url entity.Url
}

func (u Url) Alias() *string {
	return &u.url.Alias
}

func (u Url) OriginalUrl() *string {
	return &u.url.OriginalUrl
}

func (u Url) ExpireAt() *string {
	timeStr := u.url.ExpireAt.String()
	return &timeStr
}

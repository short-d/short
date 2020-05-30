package order

import "github.com/short-d/short/backend/app/entity"

var _ Order = (*Unchanged)(nil)

type Unchanged struct{}

func (u Unchanged) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	return shortLinks
}

func (u Unchanged) ArrangeUsers(users []entity.User) []entity.User {
	return users
}

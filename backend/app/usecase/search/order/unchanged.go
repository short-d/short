package order

import "github.com/short-d/short/backend/app/entity"

var _ Order = (*Unchanged)(nil)

// Unchanged represents the order where nothing is changed
type Unchanged struct{}

// ArrangeShortLinks keeps the same order of users
func (u Unchanged) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	return shortLinks
}

// ArrangeUsers keeps the same order of users
func (u Unchanged) ArrangeUsers(users []entity.User) []entity.User {
	return users
}

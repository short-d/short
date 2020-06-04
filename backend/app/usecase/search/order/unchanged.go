package order

import "github.com/short-d/short/backend/app/entity"

var _ Order = (*Unchanged)(nil)

// Unchanged keeps the order of search results untouched.
type Unchanged struct{}

// ArrangeShortLinks keeps the arrangement of shortLinks untouched.
func (u Unchanged) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	return shortLinks
}

// ArrangeUsers keeps the arrangement of users untouched.
func (u Unchanged) ArrangeUsers(users []entity.User) []entity.User {
	return users
}

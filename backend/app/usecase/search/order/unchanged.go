package order

import "github.com/short-d/short/backend/app/entity"

var _ Order = (*Unchanged)(nil)

// Unchanged represents the Order with unchanged arrangement of searchable resources.
type Unchanged struct{}

// ArrangeShortLinks keeps the same arrangement of shortLinks.
func (u Unchanged) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	return shortLinks
}

// ArrangeUsers keeps the same arrangement of users.
func (u Unchanged) ArrangeUsers(users []entity.User) []entity.User {
	return users
}

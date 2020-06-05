package order

import (
	"sort"

	"github.com/short-d/short/backend/app/entity"
)

var _ Order = (*CreatedTime)(nil)

// CreatedTime arranges searchable resources based on their created time.
type CreatedTime struct {
}

// ArrangeShortLinks arranges shortLinks based on CreatedAt.
func (c CreatedTime) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	sort.SliceStable(shortLinks, func(i, j int) bool {
		if shortLinks[j].CreatedAt == nil {
			return true
		} else if shortLinks[i].CreatedAt == nil {
			return false
		}
		return shortLinks[i].CreatedAt.Before(*shortLinks[j].CreatedAt)
	})
	return shortLinks
}

// ArrangeUsers arranges users based on CreatedAt.
func (c CreatedTime) ArrangeUsers(users []entity.User) []entity.User {
	sort.SliceStable(users, func(i, j int) bool {
		if users[j].CreatedAt == nil {
			return true
		} else if users[i].CreatedAt == nil {
			return false
		}
		return users[i].CreatedAt.Before(*users[j].CreatedAt)
	})
	return users
}

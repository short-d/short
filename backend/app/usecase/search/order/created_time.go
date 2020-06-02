package order

import (
	"sort"

	"github.com/short-d/short/backend/app/entity"
)

var _ Order = (*CreatedTime)(nil)

// CreatedTime represents the order of resources based on created time
type CreatedTime struct {
}

// ArrangeShortLinks sorts users based on time of creation
func (c CreatedTime) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	sort.SliceStable(shortLinks, func(i, j int) bool {
		return shortLinks[i].CreatedAt.Before(*shortLinks[j].CreatedAt)
	})
	return shortLinks
}

// ArrangeUsers sorts users based on time of creation
func (c CreatedTime) ArrangeUsers(users []entity.User) []entity.User {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].CreatedAt.Before(*users[j].CreatedAt)
	})
	return users
}

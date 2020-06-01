package order

import (
	"sort"

	"github.com/short-d/short/backend/app/entity"
)

var _ Order = (*CreatedTime)(nil)

type CreatedTime struct {
}

func (c CreatedTime) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	sort.SliceStable(shortLinks, func(i, j int) bool {
		return shortLinks[i].CreatedAt.Before(*shortLinks[j].CreatedAt)
	})
	return shortLinks
}

func (c CreatedTime) ArrangeUsers(users []entity.User) []entity.User {
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].CreatedAt.Before(*users[j].CreatedAt)
	})
	return users
}

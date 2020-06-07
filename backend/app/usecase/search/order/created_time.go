package order

import (
	"sort"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ Order = (*CreatedTime)(nil)

// CreatedTime arranges searchable resources based on when they are created.
type CreatedTime struct {
}

// ArrangeShortLinks arranges shortLinks based on CreatedAt.
func (c CreatedTime) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	sort.SliceStable(shortLinks, func(firstIdx, secondIdx int) bool {
		return lessTime(shortLinks[firstIdx].CreatedAt, shortLinks[secondIdx].CreatedAt)
	})
	return shortLinks
}

// ArrangeUsers arranges users based on CreatedAt.
func (c CreatedTime) ArrangeUsers(users []entity.User) []entity.User {
	sort.SliceStable(users, func(firstIdx, secondIdx int) bool {
		return lessTime(users[firstIdx].CreatedAt, users[secondIdx].CreatedAt)
	})
	return users
}

func lessTime(first, second *time.Time) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil {
		return false
	}

	if second == nil {
		return true
	}

	return first.Before(*second)
}

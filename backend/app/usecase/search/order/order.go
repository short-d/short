package order

import "github.com/short-d/short/backend/app/entity"

// By represents the order of resources being sorted in.
type By uint

const (
	// ByUnsorted keeps the original order of search results.
	ByUnsorted By = iota
	// ByCreatedTimeASC sorts search results based on their creation time.
	ByCreatedTimeASC
)

// Order arranges searchable resources based on predefined properties.
type Order interface {
	ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink
	ArrangeUsers(users []entity.User) []entity.User
}

// NewOrder creates Order.
func NewOrder(by By) Order {
	switch by {
	case ByCreatedTimeASC:
		return CreatedTime{}
	default:
		return Unchanged{}
	}
}

package order

import "github.com/short-d/short/backend/app/entity"

// By represents the order of resources being sorted in.
type By uint

// By constants represent supported types of order.
const (
	ByUnsorted By = iota
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

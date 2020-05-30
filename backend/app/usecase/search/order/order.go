package order

import "github.com/short-d/short/backend/app/entity"

type By uint

const (
	ByCreatedTimeASC By = iota
)

type Order interface {
	ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink
	ArrangeUsers(users []entity.User) []entity.User
}

func NewOrder(by By) Order {
	switch by {
	case ByCreatedTimeASC:
		return CreatedTime{}
	default:
		return Unchanged{}
	}
}

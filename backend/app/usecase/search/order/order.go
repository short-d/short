package order

import "github.com/short-d/short/backend/app/entity"

// By represents the order method of resources
type By uint

const (
	ByCreatedTimeASC By = iota
)

// Order interface orders the resources
type Order interface {
	ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink
	ArrangeUsers(users []entity.User) []entity.User
}

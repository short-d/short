package search

import "github.com/short-d/short/backend/app/usecase/search/order"

type Resource uint

const (
	ShortLink Resource = iota
	User
)

type Filter struct {
	maxResults int
	resources  []Resource
	orders     []order.By
}

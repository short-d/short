package search

import "github.com/short-d/short/backend/app/usecase/search/order"

// Resource represents Short's Resources
type Resource uint

const (
	ShortLink Resource = iota
	User
)

// Filter represents the filters for a search request
type Filter struct {
	MaxResults int
	Resources  []Resource
	Orders     []order.By
}

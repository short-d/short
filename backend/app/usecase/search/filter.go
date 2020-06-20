package search

import "github.com/short-d/short/backend/app/usecase/search/order"

// Resource represents a type of searchable objects.
type Resource uint

const (
	ShortLink Resource = iota
	User
)

// Filter represents the filters for a search request.
type Filter struct {
	MaxResults int
	Resources  []Resource
	Orders     []order.By
}

// IsValid checks if the resources and orders have one-to-one relation.
func (f *Filter) IsValid() bool {
	return len(f.Resources) == len(f.Orders)
}

// NewFilter creates Filter.
func NewFilter(maxResults int, resources []Resource, orders []order.By) Filter {
	return Filter{
		MaxResults: maxResults,
		Resources:  resources,
		Orders:     orders,
	}
}

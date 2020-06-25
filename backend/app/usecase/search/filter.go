package search

import (
	"errors"

	"github.com/short-d/short/backend/app/usecase/search/order"
)

// Resource represents a type of searchable objects.
type Resource uint

const (
	// Unknown represents the resource that is not listed or known.
	Unknown Resource = iota
	// ShortLink represents the short link resource.
	ShortLink
	// User represents the user resource.
	User
)

// Filter represents the filters for a search request.
type Filter struct {
	maxResults int
	resources  []Resource
	orders     []order.By
}

// NewFilter creates Filter.
func NewFilter(maxResults int, resources []Resource, orders []order.By) (Filter, error) {
	if len(resources) != len(orders) {
		return Filter{}, errors.New("mismatch between resources and orders")
	}
	return Filter{
		maxResults: maxResults,
		resources:  resources,
		orders:     orders,
	}, nil
}

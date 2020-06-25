package search

import (
	"errors"

	"github.com/short-d/short/backend/app/usecase/search/order"
)

// Resource represents a type of searchable objects.
type Resource uint

const (
	// Unknown implies the resource type is not support by the search module.
	// This is usually used as the fallback resource by the callers of search API.
	Unknown Resource = iota
	// ShortLink represents the short links.
	ShortLink
	// User represents the internal user.
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

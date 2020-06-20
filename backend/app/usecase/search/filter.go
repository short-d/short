package search

import "github.com/short-d/short/backend/app/usecase/search/order"

// Resource represents a type of searchable objects.
type Resource uint

const (
	ShortLink Resource = iota
	User
)

// OrderedResource represents a type of searchable objects together with the order.
type OrderedResource struct {
	Resource Resource
	Order    order.By
}

// Filter represents the filters for a search request.
type Filter struct {
	MaxResults       int
	OrderedResources []OrderedResource
}

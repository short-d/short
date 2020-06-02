package search

import "github.com/short-d/short/backend/app/entity"

// Query represents the query terms of a search request
type Query struct {
	Query string
	User  *entity.User
}

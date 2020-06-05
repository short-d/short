package search

import "github.com/short-d/short/backend/app/entity"

// Query represents a user query.
type Query struct {
	Query string
	User  *entity.User
}

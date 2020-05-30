package search

import "github.com/short-d/short/backend/app/entity"

type Query struct {
	query string
	user  *entity.User
}

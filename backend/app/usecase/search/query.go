package search

import "github.com/short-d/short/backend/app/entity"

type Query struct {
	keywords string
	user     entity.User
}

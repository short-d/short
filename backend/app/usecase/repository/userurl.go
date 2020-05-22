package repository

import "github.com/short-d/short/backend/app/entity"

// UserURLRelation accesses User-ShortLink relationship from storage, such as database.
type UserURLRelation interface {
	CreateRelation(user entity.User, url entity.ShortLink) error
	FindAliasesByUser(user entity.User) ([]string, error)
}

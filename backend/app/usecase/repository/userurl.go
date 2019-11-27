package repository

import "short/app/entity"

// UserURLRelation accesses User-URL relationship from storage, such as database.
type UserURLRelation interface {
	CreateRelation(user entity.User, url entity.URL) error
	FindAliasesByUser(user entity.User) ([]string, error)
}

package repo

import "short/app/entity"

// UserURLRelation accesses User-URL relationship from storage, such as database.
type UserURLRelation interface {
	CreateRelation(user entity.User, urlAlias string) error
}

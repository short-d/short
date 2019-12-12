package repository

import (
	"database/sql"
	"short/app/entity"
)

// UserURLRelation accesses User-URL relationship from storage, such as database.
type UserURLRelation interface {
	CreateRelation(user entity.User, url entity.URL) error
	FindAliasesByUser(user entity.User) ([]string, error)
	NewTransaction() (*sql.Tx, error)
	CreateRelationWithTransaction(tx *sql.Tx, user entity.User, url entity.URL) error
}

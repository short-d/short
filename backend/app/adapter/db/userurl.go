package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/usecase/repo"
)

var _ repo.UserURLRelation = (*UserURLRelationSql)(nil)

// UserURLRelationSql accesses UserURLRelation information in user_url_relation
// table.
type UserURLRelationSql struct {
	db *sql.DB
}

// CreateRelation establishes bi-directional relationship between a user and a
// url in user_url_relation table.
func (u UserURLRelationSql) CreateRelation(userEmail string, urlAlias string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1,$2)
`,
		table.UserURLRelation.TableName,
		table.UserURLRelation.ColumnUserEmail,
		table.UserURLRelation.ColumnURLAlias,
	)

	_, err := u.db.Exec(statement, userEmail, urlAlias)
	return err
}

// NewUserURLSql creates UserURLRelationSql
func NewUserURLSql(db *sql.DB) UserURLRelationSql {
	return UserURLRelationSql{
		db: db,
	}
}

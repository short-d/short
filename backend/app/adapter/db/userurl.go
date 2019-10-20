package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/usecase/repo"
)

var _ repo.UserURLRelation = (*UserURLRelationSql)(nil)

type UserURLRelationSql struct {
	db *sql.DB
}

func (u UserURLRelationSql) CreateRelation(userEmail string, urlAlias string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1,$2)
`,
		table.UserURLRelation.TableName,
		table.UserURLRelation.ColumnUserEmail,
		table.UserURLRelation.ColumnUrlAlias,
	)

	_, err := u.db.Exec(statement, userEmail, urlAlias)
	return err
}

func NewUserURL(db *sql.DB) UserURLRelationSql {
	return UserURLRelationSql{
		db: db,
	}
}

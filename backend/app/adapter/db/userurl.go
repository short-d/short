package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/usecase/repo"
)

var _ repo.UserURL = (*UserURL)(nil)

type UserURL struct {
	db *sql.DB
}

func (u UserURL) CreateRelation(userEmail string, urlAlias string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1,$2)
`,
		table.UserURL.TableName,
		table.UserURL.ColumnUserEmail,
		table.UserURL.ColumnUrlAlias,
	)

	_, err := u.db.Exec(statement, userEmail, urlAlias)
	return err
}

func NewUserURL(db *sql.DB) UserURL {
	return UserURL{
		db: db,
	}
}

package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.UserURLRelation = (*UserURLRelationSQL)(nil)

// UserURLRelationSQL accesses UserURLRelation information in user_url_relation
// table.
type UserURLRelationSQL struct {
	db *sql.DB
}

// CreateRelation establishes bi-directional relationship between a user and a
// url in user_url_relation table.
func (u UserURLRelationSQL) CreateRelation(user entity.User, url entity.URL) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1,$2)
`,
		table.UserURLRelation.TableName,
		table.UserURLRelation.ColumnUserID,
		table.UserURLRelation.ColumnURLAlias,
	)

	_, err := u.db.Exec(statement, user.ID, url.Alias)
	return err
}

// FindAliasesByUser fetches the aliases of all the URLs created by the given user.
// TODO(issue#260): allow API client to filter urls based on visibility.
func (u UserURLRelationSQL) FindAliasesByUser(user entity.User) ([]string, error) {
	statement := fmt.Sprintf(`SELECT "%s" FROM "%s" WHERE "%s"=$1;`,
		table.UserURLRelation.ColumnURLAlias,
		table.UserURLRelation.TableName,
		table.UserURLRelation.ColumnUserID,
	)

	var aliases []string
	rows, err := u.db.Query(statement, user.ID)
	// TODO(issue#711): errors should be checked before using defer
	defer rows.Close()
	if err != nil {
		return aliases, nil
	}

	for rows.Next() {
		var alias string
		err = rows.Scan(&alias)
		if err != nil {
			return aliases, err
		}

		aliases = append(aliases, alias)
	}

	return aliases, nil
}

// NewUserURLRelationSQL creates UserURLRelationSQL
func NewUserURLRelationSQL(db *sql.DB) UserURLRelationSQL {
	return UserURLRelationSQL{
		db: db,
	}
}

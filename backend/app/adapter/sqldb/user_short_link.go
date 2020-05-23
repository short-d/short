package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.UserShortLink = (*UserShortLinkSQL)(nil)

// UserShortLinkSQL accesses UserShortLink information in user_short_link
// table.
type UserShortLinkSQL struct {
	db *sql.DB
}

// CreateRelation establishes bi-directional relationship between a user and a
// short link in user_short_link table.
func (u UserShortLinkSQL) CreateRelation(user entity.User, shortLink entity.ShortLink) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s")
VALUES ($1,$2)
`,
		table.UserShortLink.TableName,
		table.UserShortLink.ColumnUserID,
		table.UserShortLink.ColumnShortLinkAlias,
	)

	_, err := u.db.Exec(statement, user.ID, shortLink.Alias)
	return err
}

// FindAliasesByUser fetches the aliases of all the ShortLinks created by the given user.
// TODO(issue#260): allow API client to filter urls based on visibility.
func (u UserShortLinkSQL) FindAliasesByUser(user entity.User) ([]string, error) {
	statement := fmt.Sprintf(`SELECT "%s" FROM "%s" WHERE "%s"=$1;`,
		table.UserShortLink.ColumnShortLinkAlias,
		table.UserShortLink.TableName,
		table.UserShortLink.ColumnUserID,
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

// NewUserShortLinkSQL creates UserShortLinkSQL
func NewUserShortLinkSQL(db *sql.DB) UserShortLinkSQL {
	return UserShortLinkSQL{
		db: db,
	}
}

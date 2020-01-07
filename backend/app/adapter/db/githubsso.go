package db

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.AccountMapping = (*GithubSSOSql)(nil)

// GithubSSOSql accesses mapping between Github and Short accounts from the SQL
// database.
type GithubSSOSql struct {
	db     *sql.DB
	logger fw.Logger
}

// IsSSOUserExist checks whether mapping for a given Github account exists in
// the database.
func (g GithubSSOSql) IsSSOUserExist(ssoUser entity.SSOUser) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.GithubSSO.ColumnGithubUserID,
		table.GithubSSO.TableName,
		table.GithubSSO.ColumnGithubUserID,
	)
	var id string
	err := g.db.QueryRow(query, ssoUser.ID).Scan(&id)
	if err == nil {
		return true, err
	}
	if err == sql.ErrNoRows {
		return false, nil
	}
	g.logger.Error(err)
	return false, err
}

// CreateMapping creates mapping between user's Github and Short accounts in the
// database.
func (g GithubSSOSql) CreateMapping(ssoUser entity.SSOUser, user entity.User) error {
	panic("implement me")
}

// NewGithubSSOSql creates GithubSSOSql.
func NewGithubSSOSql(db *sql.DB, logger fw.Logger) GithubSSOSql {
	return GithubSSOSql{db: db, logger: logger}
}

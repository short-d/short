package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.SSOMap = (*GithubSSOSql)(nil)

// GithubSSOSql accesses mapping between Github and Short accounts in the
// SQL database.
type GithubSSOSql struct {
	db     *sql.DB
	logger logger.Logger
}

// GetShortUserID retrieves the internal user ID that is linked to the user's
// Github account.
func (g GithubSSOSql) GetShortUserID(ssoUserID string) (string, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.GithubSSO.ColumnShortUserID,
		table.GithubSSO.TableName,
		table.GithubSSO.ColumnGithubUserID,
	)
	var id string
	err := g.db.QueryRow(query, ssoUserID).Scan(&id)
	if err == nil {
		return id, err
	}
	if err == sql.ErrNoRows {
		return "", repository.ErrEntryNotFound(
			fmt.Sprintf("user with Github ID %s not found", ssoUserID),
		)
	}
	g.logger.Error(err)
	return "", err
}

// IsSSOUserExist checks whether mapping for a given Github account exists in
// the database.
func (g GithubSSOSql) IsSSOUserExist(ssoUserID string) (bool, error) {
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
	err := g.db.QueryRow(query, ssoUserID).Scan(&id)
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
func (g GithubSSOSql) CreateMapping(ssoUserID string, userID string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.GithubSSO.TableName,
		table.GithubSSO.ColumnGithubUserID,
		table.GithubSSO.ColumnShortUserID,
	)
	_, err := g.db.Exec(statement, ssoUserID, userID)
	return err
}

// NewGithubSSOSql creates GithubSSOSql.
func NewGithubSSOSql(db *sql.DB, logger logger.Logger) GithubSSOSql {
	return GithubSSOSql{db: db, logger: logger}
}

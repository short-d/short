package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.SSOMap = (*FacebookSSOSql)(nil)

// FacebookSSOSql accesses mapping between Facebook and Short accounts in
// SQL database.
type FacebookSSOSql struct {
	db     *sql.DB
	logger logger.Logger
}

// GetShortUserID retrieves the internal user ID that is linked to the user's
// Facebook account.
func (g FacebookSSOSql) GetShortUserID(ssoUserID string) (string, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.FacebookSSO.ColumnShortUserID,
		table.FacebookSSO.TableName,
		table.FacebookSSO.ColumnFacebookUserID,
	)
	var id string
	err := g.db.QueryRow(query, ssoUserID).Scan(&id)
	if err == nil {
		return id, err
	}
	if err == sql.ErrNoRows {
		return "", repository.ErrEntryNotFound(
			fmt.Sprintf("user with Facebook ID %s not found", ssoUserID),
		)
	}
	g.logger.Error(err)
	return "", err
}

// IsSSOUserExist checks whether mapping for a given Facebook account exists in
// the database.
func (g FacebookSSOSql) IsSSOUserExist(ssoUserID string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.FacebookSSO.ColumnFacebookUserID,
		table.FacebookSSO.TableName,
		table.FacebookSSO.ColumnFacebookUserID,
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

// CreateMapping creates links user's Facebook and Short accounts in the
// database.
func (g FacebookSSOSql) CreateMapping(ssoUserID string, userID string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.FacebookSSO.TableName,
		table.FacebookSSO.ColumnFacebookUserID,
		table.FacebookSSO.ColumnShortUserID,
	)
	_, err := g.db.Exec(statement, ssoUserID, userID)
	return err
}

// NewFacebookSSOSql creates FacebookSSOSql.
func NewFacebookSSOSql(db *sql.DB, logger logger.Logger) FacebookSSOSql {
	return FacebookSSOSql{db: db, logger: logger}
}

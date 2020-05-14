package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.SSOMap = (*GoogleSSOSql)(nil)

// GoogleSSOMapSql accesses mapping between Google and Short accounts from the
// SQL database.
type GoogleSSOSql struct {
	db     *sql.DB
	logger logger.Logger
}

func (g GoogleSSOSql) GetShortUserID(ssoUserID string) (string, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.GoogleSSO.ColumnShortUserID,
		table.GoogleSSO.TableName,
		table.GoogleSSO.ColumnGoogleUserID,
	)
	var id string
	err := g.db.QueryRow(query, ssoUserID).Scan(&id)
	if err == nil {
		return id, err
	}
	if err == sql.ErrNoRows {
		return "", repository.ErrEntryNotFound(
			fmt.Sprintf("user with Google ID %s not found", ssoUserID),
		)
	}
	g.logger.Error(err)
	return "", err
}

// IsSSOUserExist checks whether mapping for a given Google account exists in
// the database.
func (g GoogleSSOSql) IsSSOUserExist(ssoUserID string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s"
FROM "%s"
WHERE "%s"=$1;
`,
		table.GoogleSSO.ColumnGoogleUserID,
		table.GoogleSSO.TableName,
		table.GoogleSSO.ColumnGoogleUserID,
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

// CreateMapping creates mapping between user's Google and Short accounts in the
// database.
func (g GoogleSSOSql) CreateMapping(ssoUserID string, userID string) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);
`,
		table.GoogleSSO.TableName,
		table.GoogleSSO.ColumnGoogleUserID,
		table.GoogleSSO.ColumnShortUserID,
	)
	_, err := g.db.Exec(statement, ssoUserID, userID)
	return err
}

// NewGoogleSSOSql creates GoogleSSOSql.
func NewGoogleSSOSql(db *sql.DB, logger logger.Logger) GoogleSSOSql {
	return GoogleSSOSql{db: db, logger: logger}
}

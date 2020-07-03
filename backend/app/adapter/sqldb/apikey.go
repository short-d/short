package sqldb

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.APIKey = (*APIKeySQL)(nil)

// APIKeySQL accesses APIKey from the database through SQL.
type APIKeySQL struct {
	db *sql.DB
}

// GetAPIKey fetches APIKey from APIKey table using SQL.
func (a APIKeySQL) GetAPIKey(appID string, key string) (entity.APIKey, error) {
	query := fmt.Sprintf(`
SELECT "%s", "%s" 
FROM "%s" WHERE "%s"=$1 AND "%s"=$2;
`,
		table.APIKey.ColumnDisabled,
		table.APIKey.ColumnCreatedAt,
		table.APIKey.TableName,
		table.APIKey.ColumnAppID,
		table.APIKey.ColumnKey,
	)
	apiKey := entity.APIKey{}
	err := a.db.QueryRow(query, appID, key).Scan(&apiKey.IsDisabled, &apiKey.CreatedAt)
	if err == nil {
		apiKey.AppID = appID
		apiKey.Key = key
		return apiKey, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return entity.APIKey{},
			repository.ErrEntryNotFound(
				fmt.Sprintf("appID(%s) and key(%s) not found", appID, key))
	}
	return entity.APIKey{}, err
}

// CreateAPIKey appends a new APIKey entry to APIKey table using SQL.
func (a APIKeySQL) CreateAPIKey(input entity.APIKeyInput) (entity.APIKey, error) {
	stmt := fmt.Sprintf(`
INSERT INTO "%s"("%s", "%s", "%s", "%s")
VALUES ($1, $2, $3, $4);
`,
		table.APIKey.TableName,
		table.APIKey.ColumnAppID,
		table.APIKey.ColumnKey,
		table.APIKey.ColumnDisabled,
		table.APIKey.ColumnCreatedAt,
	)

	isDisabled := input.GetIsDisabled(false)
	_, err := a.db.Exec(
		stmt,
		input.GetAppID(""),
		input.GetKey(""),
		SQLBool(isDisabled),
		input.GetCreatedAt(time.Time{}),
	)

	return entity.APIKey{
		AppID:      input.GetAppID(""),
		Key:        input.GetKey(""),
		IsDisabled: isDisabled,
		CreatedAt:  input.GetCreatedAt(time.Time{}),
	}, err
}

// NewAPIKeySQL creates database access object for APIKey.
func NewAPIKeySQL(db *sql.DB) APIKeySQL {
	return APIKeySQL{db: db}
}

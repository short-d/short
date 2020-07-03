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

var _ repository.ApiKey = (*ApiKeySQL)(nil)

// ApiKeySQL accesses ApiKey from the database through SQL.
type ApiKeySQL struct {
	db *sql.DB
}

// GetApiKey fetches ApiKey from ApiKey table using SQL.
func (a ApiKeySQL) GetApiKey(appID string, key string) (entity.ApiKey, error) {
	query := fmt.Sprintf(`
SELECT "%s", "%s" 
FROM "%s" WHERE "%s"=$1 AND "%s"=$2;
`,
		table.ApiKey.ColumnDisabled,
		table.ApiKey.ColumnCreatedAt,
		table.ApiKey.TableName,
		table.ApiKey.ColumnAppID,
		table.ApiKey.ColumnKey,
	)
	apiKey := entity.ApiKey{}
	err := a.db.QueryRow(query, appID, key).Scan(&apiKey.IsDisabled, &apiKey.CreatedAt)
	if err == nil {
		apiKey.AppID = appID
		apiKey.Key = key
		return apiKey, nil
	}
	if errors.Is(err, sql.ErrNoRows) {
		return entity.ApiKey{},
			repository.ErrEntryNotFound(
				fmt.Sprintf("appID(%s) and key(%s) not found", appID, key))
	}
	return entity.ApiKey{}, err
}

// GetApiKey appends a new ApiKey entry to ApiKey table using SQL.
func (a ApiKeySQL) CreateApiKey(input entity.ApiKeyInput) (entity.ApiKey, error) {
	stmt := fmt.Sprintf(`
INSERT INTO "%s"("%s", "%s", "%s", "%s")
VALUES ($1, $2, $3, $4);
`,
		table.ApiKey.TableName,
		table.ApiKey.ColumnAppID,
		table.ApiKey.ColumnKey,
		table.ApiKey.ColumnDisabled,
		table.ApiKey.ColumnCreatedAt,
	)

	isDisabled := input.GetIsDisabled(false)
	_, err := a.db.Exec(
		stmt,
		input.GetAppID(""),
		input.GetKey(""),
		SQLBool(isDisabled),
		input.GetCreatedAt(time.Time{}),
	)

	return entity.ApiKey{
		AppID:      input.GetAppID(""),
		Key:        input.GetKey(""),
		IsDisabled: isDisabled,
		CreatedAt:  input.GetCreatedAt(time.Time{}),
	}, err
}

// NewApiKeySQL creates database access object for ApiKey.
func NewApiKeySQL(db *sql.DB) ApiKeySQL {
	return ApiKeySQL{db: db}
}

// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
)

var insertAPIKeyRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s, %s) 
VALUES ($1, $2, $3, $4);`,
	table.APIKey.TableName,
	table.APIKey.ColumnAppID,
	table.APIKey.ColumnKey,
	table.APIKey.ColumnDisabled,
	table.APIKey.ColumnCreatedAt,
)

type apiKeyTableRow struct {
	appID      string
	key        string
	isDisabled bool
	createdAt  time.Time
}

func TestAPIKeySQL_GetAPIKey(t *testing.T) {
	testCases := []struct {
		name            string
		appTableRows    []appTableRow
		apiKeyTableRows []apiKeyTableRow
		appID           string
		key             string
		hasErr          bool
		expectedAPIKey  entity.APIKey
	}{
		{
			name:            "app not found",
			appTableRows:    []appTableRow{},
			apiKeyTableRows: []apiKeyTableRow{},
			appID:           "emotic",
			hasErr:          true,
		},
		{
			name: "app found with different key",
			appTableRows: []appTableRow{
				{
					id: "emotic",
				},
			},
			apiKeyTableRows: []apiKeyTableRow{
				{
					appID: "emotic",
					key:   "key1",
				},
			},
			appID:  "emotic",
			key:    "key2",
			hasErr: true,
		},
		{
			name: "app found with the same key",
			appTableRows: []appTableRow{
				{
					id: "emotic",
				},
			},
			apiKeyTableRows: []apiKeyTableRow{
				{
					appID:      "emotic",
					key:        "key",
					isDisabled: false,
				},
			},
			appID:  "emotic",
			key:    "key",
			hasErr: false,
			expectedAPIKey: entity.APIKey{
				AppID:      "emotic",
				Key:        "key",
				IsDisabled: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertAppRows(t, sqlDB, testCase.appTableRows)
					insertAPIKeyRows(t, sqlDB, testCase.apiKeyTableRows)

					apiKeyRepo := sqldb.NewAPIKeySQL(sqlDB)
					gotAPIKey, err := apiKeyRepo.GetAPIKey(testCase.appID, testCase.key)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedAPIKey.AppID, gotAPIKey.AppID)
					assert.Equal(t, testCase.expectedAPIKey.Key, gotAPIKey.Key)
					assert.Equal(t, testCase.expectedAPIKey.IsDisabled, gotAPIKey.IsDisabled)
				})
		})
	}
}

func TestAPIKeySQL_CreateAPIKey(t *testing.T) {
	testCases := []struct {
		name            string
		appTableRows    []appTableRow
		apiKeyTableRows []apiKeyTableRow
		appID           string
		key             string
		isDisabled      bool
		hasErr          bool
		expectedAPIKey  entity.APIKey
	}{
		{
			name:            "app not found",
			appTableRows:    []appTableRow{},
			apiKeyTableRows: []apiKeyTableRow{},
			appID:           "emotic",
			hasErr:          true,
		},
		{
			name: "app found with duplicated key",
			appTableRows: []appTableRow{
				{
					id: "emotic",
				},
			},
			apiKeyTableRows: []apiKeyTableRow{
				{
					appID: "emotic",
					key:   "key",
				},
			},
			appID:  "emotic",
			key:    "key",
			hasErr: true,
		},
		{
			name: "app found with different key",
			appTableRows: []appTableRow{
				{
					id: "emotic",
				},
			},
			apiKeyTableRows: []apiKeyTableRow{
				{
					appID:      "emotic",
					key:        "key1",
					isDisabled: false,
				},
			},
			appID:      "emotic",
			key:        "key2",
			isDisabled: false,
			hasErr:     false,
			expectedAPIKey: entity.APIKey{
				AppID:      "emotic",
				Key:        "key2",
				IsDisabled: false,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			dbtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertAppRows(t, sqlDB, testCase.appTableRows)
					insertAPIKeyRows(t, sqlDB, testCase.apiKeyTableRows)

					apiKeyRepo := sqldb.NewAPIKeySQL(sqlDB)
					input := entity.APIKeyInput{
						AppID:      &testCase.appID,
						Key:        &testCase.key,
						IsDisabled: &testCase.isDisabled,
						CreatedAt:  nil,
					}
					gotAPIKey, err := apiKeyRepo.CreateAPIKey(input)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedAPIKey.AppID, gotAPIKey.AppID)
					assert.Equal(t, testCase.expectedAPIKey.Key, gotAPIKey.Key)
					assert.Equal(t, testCase.expectedAPIKey.IsDisabled, gotAPIKey.IsDisabled)
				})
		})
	}
}

func insertAPIKeyRows(t *testing.T, sqlDB *sql.DB, tableRows []apiKeyTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertAPIKeyRowSQL,
			tableRow.appID,
			tableRow.key,
			sqldb.SQLBool(tableRow.isDisabled),
			tableRow.createdAt,
		)
		assert.Equal(t, nil, err)
	}
}

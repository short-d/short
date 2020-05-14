// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
)

type GoogleSSOTableRow struct {
	googleUserID string
	shortUserID  string
}

func TestGoogleSSOSql_IsSSOUserExist(t *testing.T) {
	testCases := []struct {
		name            string
		tableRows       []GoogleSSOTableRow
		ssoUserID       string
		expectedIsExist bool
	}{
		{
			name:            "sso user not found",
			tableRows:       []GoogleSSOTableRow{},
			ssoUserID:       "220uFicCJj",
			expectedIsExist: false,
		},
		{
			name: "sso user exists",
			tableRows: []GoogleSSOTableRow{
				{
					googleUserID: "220uFicCJj",
					shortUserID:  "alpha",
				},
			},
			ssoUserID:       "220uFicCJj",
			expectedIsExist: true,
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
					insertGoogleSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					GoogleSSORepo := sqldb.NewGoogleSSOSql(sqlDB, lg)
					gotIsExist, err := GoogleSSORepo.IsSSOUserExist(testCase.ssoUserID)

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedIsExist, gotIsExist)
				})
		})
	}
}

func TestGoogleSSOSql_CreateMapping(t *testing.T) {
	testCases := []struct {
		name        string
		tableRows   []GoogleSSOTableRow
		ssoUserID   string
		shortUserID string
		hasErr      bool
	}{
		{
			name: "mapping exists",
			tableRows: []GoogleSSOTableRow{
				{googleUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "only SSO user ID exists",
			tableRows: []GoogleSSOTableRow{
				{googleUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "alpha",
			hasErr:      true,
		},
		{
			name: "only Short user ID exists",
			tableRows: []GoogleSSOTableRow{
				{googleUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "another_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "neither SSO user ID nor Short user ID exists",
			tableRows: []GoogleSSOTableRow{
				{googleUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "another_user_id",
			shortUserID: "alpha",
			hasErr:      false,
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
					insertGoogleSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					GoogleSSORepo := sqldb.NewGoogleSSOSql(sqlDB, lg)
					err = GoogleSSORepo.CreateMapping(testCase.ssoUserID, testCase.shortUserID)
					assert.Equal(t, nil, err)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
				})
		})
	}
}

var insertGoogleSSORowSQL = fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2)`,
	table.GoogleSSO.TableName,
	table.GoogleSSO.ColumnGoogleUserID,
	table.GoogleSSO.ColumnShortUserID,
)

func insertGoogleSSOTableRows(t *testing.T, sqlDB *sql.DB, rows []GoogleSSOTableRow) {
	for _, row := range rows {
		_, err := sqlDB.Exec(
			insertGoogleSSORowSQL,
			row.googleUserID,
			row.shortUserID,
		)
		assert.Equal(t, nil, err)
	}
}

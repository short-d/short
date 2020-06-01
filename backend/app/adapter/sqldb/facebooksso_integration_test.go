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

type FacebookSSOTableRow struct {
	facebookUserID string
	shortUserID    string
}

func TestFacebookSSOSql_IsSSOUserExist(t *testing.T) {
	testCases := []struct {
		name            string
		userTableRows   []userTableRow
		tableRows       []FacebookSSOTableRow
		ssoUserID       string
		expectedIsExist bool
	}{
		{
			name:            "sso user not found",
			userTableRows:   []userTableRow{},
			tableRows:       []FacebookSSOTableRow{},
			ssoUserID:       "220uFicCJj",
			expectedIsExist: false,
		},
		{
			name: "sso user exists",
			userTableRows: []userTableRow{
				{
					id:    "alpha",
					email: "alpha@gmail.com",
					name:  "alpha",
				},
			},
			tableRows: []FacebookSSOTableRow{
				{
					facebookUserID: "220uFicCJj",
					shortUserID:    "alpha",
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertFacebookSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					FacebookSSORepo := sqldb.NewFacebookSSOSql(sqlDB, lg)
					gotIsExist, err := FacebookSSORepo.IsSSOUserExist(testCase.ssoUserID)

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedIsExist, gotIsExist)
				})
		})
	}
}

func TestFacebookSSOSql_CreateMapping(t *testing.T) {
	defaultUserTableRows := []userTableRow{
		{
			id:    "short",
			email: "short@gmail.com",
			name:  "short",
		},
		{
			id:    "alpha",
			email: "alpha@gmail.com",
			name:  "alpha",
		},
	}

	testCases := []struct {
		name          string
		userTableRows []userTableRow
		tableRows     []FacebookSSOTableRow
		ssoUserID     string
		shortUserID   string
		hasErr        bool
	}{
		{
			name: "mapping exists",
			userTableRows: defaultUserTableRows,
			tableRows: []FacebookSSOTableRow{
				{facebookUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "only SSO user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []FacebookSSOTableRow{
				{facebookUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "alpha",
			hasErr:      true,
		},
		{
			name: "only Short user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []FacebookSSOTableRow{
				{facebookUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "another_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "neither SSO user ID nor Short user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []FacebookSSOTableRow{
				{facebookUserID: "long_user_id", shortUserID: "short"},
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertFacebookSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					FacebookSSORepo := sqldb.NewFacebookSSOSql(sqlDB, lg)
					err = FacebookSSORepo.CreateMapping(testCase.ssoUserID, testCase.shortUserID)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
				})
		})
	}
}

var insertFacebookSSORowSQL = fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2)`,
	table.FacebookSSO.TableName,
	table.FacebookSSO.ColumnFacebookUserID,
	table.FacebookSSO.ColumnShortUserID,
)

func insertFacebookSSOTableRows(t *testing.T, sqlDB *sql.DB, rows []FacebookSSOTableRow) {
	for _, row := range rows {
		_, err := sqlDB.Exec(
			insertFacebookSSORowSQL,
			row.facebookUserID,
			row.shortUserID,
		)
		assert.Equal(t, nil, err)
	}
}

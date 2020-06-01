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

type githubSSOTableRow struct {
	githubUserID string
	shortUserID  string
}

func TestGithubSSOSql_IsSSOUserExist(t *testing.T) {
	testCases := []struct {
		name            string
		userTableRows   []userTableRow
		tableRows       []githubSSOTableRow
		ssoUserID       string
		expectedIsExist bool
	}{
		{
			name:            "sso user not found",
			userTableRows:   []userTableRow{},
			tableRows:       []githubSSOTableRow{},
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
			tableRows: []githubSSOTableRow{
				{
					githubUserID: "220uFicCJj",
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertGithubSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					githubSSORepo := sqldb.NewGithubSSOSql(sqlDB, lg)
					gotIsExist, err := githubSSORepo.IsSSOUserExist(testCase.ssoUserID)

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedIsExist, gotIsExist)
				})
		})
	}
}

func TestGithubSSOSql_CreateMapping(t *testing.T) {
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
		tableRows     []githubSSOTableRow
		ssoUserID     string
		shortUserID   string
		hasErr        bool
	}{
		{
			name: "mapping exists",
			userTableRows: defaultUserTableRows,
			tableRows: []githubSSOTableRow{
				{githubUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "only SSO user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []githubSSOTableRow{
				{githubUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "long_user_id",
			shortUserID: "alpha",
			hasErr:      true,
		},
		{
			name: "only Short user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []githubSSOTableRow{
				{githubUserID: "long_user_id", shortUserID: "short"},
			},
			ssoUserID:   "another_user_id",
			shortUserID: "short",
			hasErr:      true,
		},
		{
			name: "neither SSO user ID nor Short user ID exists",
			userTableRows: defaultUserTableRows,
			tableRows: []githubSSOTableRow{
				{githubUserID: "long_user_id", shortUserID: "short"},
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
					insertGithubSSOTableRows(t, sqlDB, testCase.tableRows)

					entryRepo := logger.NewEntryRepoFake()
					lg, err := logger.NewFake(logger.LogOff, &entryRepo)
					assert.Equal(t, nil, err)

					githubSSORepo := sqldb.NewGithubSSOSql(sqlDB, lg)
					err = githubSSORepo.CreateMapping(testCase.ssoUserID, testCase.shortUserID)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
				})
		})
	}
}

var insertGithubSSORowSQL = fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2)`,
	table.GithubSSO.TableName,
	table.GithubSSO.ColumnGithubUserID,
	table.GithubSSO.ColumnShortUserID,
)

func insertGithubSSOTableRows(t *testing.T, sqlDB *sql.DB, rows []githubSSOTableRow) {
	for _, row := range rows {
		_, err := sqlDB.Exec(
			insertGithubSSORowSQL,
			row.githubUserID,
			row.shortUserID,
		)
		assert.Equal(t, nil, err)
	}
}

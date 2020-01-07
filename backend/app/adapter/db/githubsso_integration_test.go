// +build integration all

package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
)

type githubSSOTableRow struct {
	githubUserID string
	shortUserID  string
}

func TestGithubSSOSql_IsSSOUserExist(t *testing.T) {
	testCases := []struct {
		name            string
		tableRows       []githubSSOTableRow
		ssoUser         entity.SSOUser
		expectedIsExist bool
	}{
		{
			name:      "sso user not found",
			tableRows: []githubSSOTableRow{},
			ssoUser: entity.SSOUser{
				ID: "220uFicCJj",
			},
			expectedIsExist: false,
		},
		{
			name: "sso user exists",
			tableRows: []githubSSOTableRow{
				{
					githubUserID: "220uFicCJj",
					shortUserID:  "alpha",
				},
			},
			ssoUser: entity.SSOUser{
				ID: "220uFicCJj",
			},
			expectedIsExist: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertGithubSSOTableRows(t, sqlDB, testCase.tableRows)

					logger := mdtest.NewLoggerFake(mdtest.FakeLoggerArgs{})
					githubSSORepo := db.NewGithubSSOSql(sqlDB, &logger)
					gotIsExist, err := githubSSORepo.IsSSOUserExist(testCase.ssoUser)

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedIsExist, gotIsExist)
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
		mdtest.Equal(t, nil, err)
	}
}

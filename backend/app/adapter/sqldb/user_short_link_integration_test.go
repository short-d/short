// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
)

var insertUserShortLinkRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s)
VALUES ($1, $2)`,
	table.UserShortLink.TableName,
	table.UserShortLink.ColumnShortLinkAlias,
	table.UserShortLink.ColumnUserID,
)

type userShortLinkTableRow struct {
	alias     string
	userEmail string
}

func TestListShortLinkSql_FindAliasesByUser(t *testing.T) {
	now := mustParseTime(t, "2019-05-01T08:02:16Z")

	testCases := []struct {
		name               string
		userTableRows      []userTableRow
		shortLinkTableRows []shortLinkTableRow
		relationTableRows  []userShortLinkTableRow
		user               entity.User
		hasErr             bool
		expectedAliases    []string
	}{
		{
			name:               "no alias found",
			userTableRows:      []userTableRow{},
			shortLinkTableRows: []shortLinkTableRow{},
			relationTableRows:  []userShortLinkTableRow{},
			user: entity.User{
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: &now,
				CreatedAt:      &now,
				UpdatedAt:      &now,
			},
			hasErr:          false,
			expectedAliases: nil,
		},
		{
			name: "aliases found",
			userTableRows: []userTableRow{
				{email: "test@example.com"},
			},
			shortLinkTableRows: []shortLinkTableRow{
				{alias: "abcd-123-xyz"},
			},
			relationTableRows: []userShortLinkTableRow{
				{
					alias:     "abcd-123-xyz",
					userEmail: "test@example.com",
				},
			},
			user: entity.User{
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: &now,
				CreatedAt:      &now,
				UpdatedAt:      &now,
			},
			hasErr: false,
			expectedAliases: []string{
				"abcd-123-xyz",
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertShortLinkTableRows(t, sqlDB, testCase.shortLinkTableRows)
					insertUserShortLinkTableRows(t, sqlDB, testCase.relationTableRows)

					userShortLinkRepo := sqldb.NewUserShortLinkSQL(sqlDB)
					result, err := userShortLinkRepo.FindAliasesByUser(testCase.user)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedAliases, result)
				})
		})
	}
}

func insertUserShortLinkTableRows(
	t *testing.T,
	sqlDB *sql.DB,
	tableRows []userShortLinkTableRow,
) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserShortLinkRowSQL,
			tableRow.alias,
			tableRow.userEmail,
		)
		assert.Equal(t, nil, err)
	}
}

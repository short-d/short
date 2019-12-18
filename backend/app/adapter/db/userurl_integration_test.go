// +build integration

package db_test

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db"
	"short/app/adapter/db/table"
	"short/app/entity"
	"testing"

	"github.com/byliuyang/app/mdtest"
)

var insertUserURLRelationRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s)
VALUES ($1, $2)`,
	table.UserURLRelation.TableName,
	table.UserURLRelation.ColumnURLAlias,
	table.UserURLRelation.ColumnUserEmail,
)

type userURLRelationTableRow struct {
	alias     string
	userEmail string
}

func TestListURLSql_FindAliasesByUser(t *testing.T) {
	now := mustParseTime(t, "2019-05-01T08:02:16Z")

	testCases := []struct {
		name            string
		tableRows       []userURLRelationTableRow
		user            entity.User
		hasErr          bool
		expectedAliases []string
	}{
		{
			name:      "no alias found",
			tableRows: []userURLRelationTableRow{},
			user: entity.User{
				Name:           "mockedUser",
				Email:          "test@example.com",
				LastSignedInAt: &now,
				CreatedAt:      &now,
				UpdatedAt:      &now,
			},
			hasErr:          false,
			expectedAliases: []string{},
		},
		{
			name: "aliases found",
			tableRows: []userURLRelationTableRow{
				{alias: "abcd-123-xyz"},
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
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					userURLRelationRepo := db.NewUserURLRelationSQL(sqlDB)
					result, err := userURLRelationRepo.FindAliasesByUser(testCase.user)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedAliases, result)
				})
		})
	}
}

func insertUserURLRelationTableRows(
	t *testing.T,
	sqlDB *sql.DB,
	tableRows []userURLRelationTableRow,
) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserURLRelationRowSQL,
			tableRow.alias,
			tableRow.userEmail,
		)
		mdtest.Equal(t, nil, err)
	}
}

package db

import (
	"fmt"
	"short/app/adapter/db/table"
	"short/app/entity"
	"testing"

	"github.com/byliuyang/app/mdtest"
)

func TestListURLSql_FindAliasesByUser(t *testing.T) {
	var mockedEmptyAliases []string
	var mockedAliases []string
	mockedUser := entity.User{
		Name:           "mockedUser",
		Email:          "test@example.com",
		LastSignedInAt: mustParseSQLTime("2019-05-01 08:02:16"),
		CreatedAt:      mustParseSQLTime("2019-05-01 08:00:16"),
		UpdatedAt:      mustParseSQLTime("2019-05-01 08:02:16"),
	}

	alias := "abcd-123-xyz"
	mockedAliases = append(mockedAliases, alias)

	testCases := []struct {
		name            string
		tableRows       *mdtest.TableRows
		user            entity.User
		hasErr          bool
		expectedAliases []string
	}{
		{
			name: "no alias found",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
			}),
			user:            mockedUser,
			hasErr:          false,
			expectedAliases: mockedEmptyAliases,
		},
		{
			name: "aliases found",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
			}).AddRow(
				"abcd-123-xyz",
			),
			user:            mockedUser,
			hasErr:          false,
			expectedAliases: mockedAliases,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT "%s" FROM "%s" WHERE "%s"=.+$`,
				table.UserURLRelation.ColumnURLAlias,
				table.UserURLRelation.TableName,
				table.UserURLRelation.ColumnUserEmail)
			stub.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			userURLRelationRepo := NewUserURLRelationSQL(db)
			result, err := userURLRelationRepo.FindAliasesByUser(testCase.user)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedAliases, result)
		})
	}
}

// +build integration

package db

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"short/app/adapter/db/sqltest"
	"short/app/adapter/db/table"
	"short/app/entity"
	"testing"

	"github.com/byliuyang/app/mdtest"
)

func TestURLSql_IsAliasExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  *mdtest.TableRows
		alias      string
		expIsExist bool
	}{
		{
			name:  "alias doesn't exist",
			alias: "gg",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
			}),
			expIsExist: false,
		},
		{
			name:  "alias found",
			alias: "gg",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
			}).
				AddRow("gg"),
			expIsExist: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			expQuery := fmt.Sprintf(`^SELECT ".+" FROM "%s" WHERE "%s"=.+$`, table.URL.TableName, table.URL.ColumnAlias)
			stub.ExpectQuery(expQuery).WillReturnRows(testCase.tableRows)

			urlRepo := NewURLSql(db)
			gotIsExist, err := urlRepo.IsAliasExist(testCase.alias)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expIsExist, gotIsExist)
		})
	}
}

func TestURLSql_GetByAlias(t *testing.T) {
	testCases := []struct {
		name        string
		tableRows   *mdtest.TableRows
		alias       string
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name: "alias not found",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
				table.URL.ColumnOriginalURL,
				table.URL.ColumnExpireAt,
				table.URL.ColumnCreatedAt,
				table.URL.ColumnUpdatedAt,
			}),
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name: "found url",
			tableRows: mdtest.NewTableRows([]string{
				table.URL.ColumnAlias,
				table.URL.ColumnOriginalURL,
				table.URL.ColumnExpireAt,
				table.URL.ColumnCreatedAt,
				table.URL.ColumnUpdatedAt,
			}).AddRow(
				"220uFicCJj",
				"http://www.google.com",
				sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
				sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				nil,
			).AddRow(
				"yDOBcj5HIPbUAsw",
				"http://www.facebook.com",
				sqltest.MustParseSQLTime("2018-04-02 08:02:16"),
				sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				nil,
			),
			alias:  "220uFicCJj",
			hasErr: false,
			expectedURL: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
				CreatedAt:   sqltest.MustParseSQLTime("2017-05-01 08:02:16"),
				UpdatedAt:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()
			mdtest.Equal(t, nil, err)
			defer db.Close()

			statement := fmt.Sprintf(`^SELECT .+ FROM "%s" WHERE "%s"=.+$`, table.URL.TableName, table.URL.ColumnAlias)
			stub.ExpectQuery(statement).WillReturnRows(testCase.tableRows)

			urlRepo := NewURLSql(db)
			url, err := urlRepo.GetByAlias("220uFicCJj")

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedURL, url)
		})
	}
}

func TestURLSql_Create(t *testing.T) {
	testCases := []struct {
		name     string
		url      entity.URL
		rowExist bool
		hasErr   bool
	}{
		{
			name: "url exists",
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
			},
			rowExist: true,
			hasErr:   true,
		},
		{
			name: "successfully create url",
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    sqltest.MustParseSQLTime("2019-05-01 08:02:16"),
			},
			rowExist: false,
			hasErr:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, stub, err := mdtest.NewSQLStub()

			mdtest.Equal(t, nil, err)
			defer db.Close()

			statement := fmt.Sprintf(`INSERT INTO "%s" .+ VALUES .+`, table.URL.TableName)

			if testCase.rowExist {
				stub.ExpectExec(statement).WillReturnError(errors.New("row exists"))
			} else {
				stub.ExpectExec(statement).WillReturnResult(driver.ResultNoRows)
			}

			urlRepo := NewURLSql(db)
			err = urlRepo.Create(testCase.url)

			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
		})
	}
}

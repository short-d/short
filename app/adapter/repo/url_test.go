package repo

import (
	"database/sql/driver"
	"fmt"
	"short/app/adapter/repo/table"
	"short/app/entity"
	"testing"
	"time"

	"github.com/pkg/errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var tableRows = []string{
	table.Url.ColumnAlias,
	table.Url.ColumnOriginalUrl,
	table.Url.ColumnExpireAt,
	table.Url.ColumnCreatedAt,
	table.Url.ColumnUpdatedAt,
}

func TestUrlSql_GetByAlias(t *testing.T) {
	testCases := []struct {
		name        string
		tableRows   *sqlmock.Rows
		alias       string
		hasErr      bool
		expectedUrl entity.Url
	}{
		{
			name:      "alias not found",
			tableRows: sqlmock.NewRows(tableRows),
			alias:     "220uFicCJj",
			hasErr:    true,
		},
		{
			name: "found url",
			tableRows: sqlmock.NewRows(tableRows).AddRow(
				"220uFicCJj",
				"http://www.google.com",
				mustParseSqlTime("2019-05-01 08:02:16"),
				mustParseSqlTime("2017-05-01 08:02:16"),
				nil,
			).AddRow(
				"yDOBcj5HIPbUAsw",
				"http://www.facebook.com",
				mustParseSqlTime("2018-04-02 08:02:16"),
				mustParseSqlTime("2017-05-01 08:02:16"),
				nil,
			),
			alias:  "220uFicCJj",
			hasErr: false,
			expectedUrl: entity.Url{
				Alias:       "220uFicCJj",
				OriginalUrl: "http://www.google.com",
				ExpireAt:    mustParseSqlTime("2019-05-01 08:02:16"),
				CreatedAt:   mustParseSqlTime("2017-05-01 08:02:16"),
				UpdatedAt:   nil,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()

			assert.Nil(t, err)
			defer db.Close()

			statement := fmt.Sprintf(`^SELECT .+ FROM "%s" WHERE "%s"=.+$`, table.Url.TableName, table.Url.ColumnAlias)
			mock.ExpectQuery(statement).WillReturnRows(testCase.tableRows)

			urlRepo := NewUrlSql(db)
			url, err := urlRepo.GetByAlias("220uFicCJj")

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedUrl, url)
			}
		})
	}
}

func TestUrlFake_Create(t *testing.T) {
	testCases := []struct {
		name     string
		url      entity.Url
		rowExist bool
		hasErr   bool
	}{
		{
			name: "url exists",
			url: entity.Url{
				Alias:       "220uFicCJj",
				OriginalUrl: "http://www.google.com",
				ExpireAt:    mustParseSqlTime("2019-05-01 08:02:16"),
				CreatedAt:   mustParseSqlTime("2017-05-01 08:02:16"),
				UpdatedAt:   nil,
			},
			rowExist: true,
			hasErr:   true,
		},
		{
			name: "successfully create url",
			url: entity.Url{
				Alias:       "220uFicCJj",
				OriginalUrl: "http://www.google.com",
				ExpireAt:    mustParseSqlTime("2019-05-01 08:02:16"),
				CreatedAt:   mustParseSqlTime("2017-05-01 08:02:16"),
				UpdatedAt:   nil,
			},
			rowExist: false,
			hasErr:   false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()

			assert.Nil(t, err)
			defer db.Close()

			statement := fmt.Sprintf(`INSERT INTO "%s" .+ VALUES .+`, table.Url.TableName)

			if testCase.rowExist {
				mock.ExpectExec(statement).WillReturnError(errors.New("row exists"))
			} else {
				mock.ExpectExec(statement).WillReturnResult(driver.ResultNoRows)
			}

			urlRepo := NewUrlSql(db)
			err = urlRepo.Create(testCase.url)

			if testCase.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

var dateTimeFmt = "2006-01-02 15:04:05"

func mustParseSqlTime(dateTime string) *time.Time {
	if dateTime == "NULL" {
		return nil
	}

	dt, err := time.Parse(dateTimeFmt, dateTime)
	if err != nil {
		panic(err)
	}
	return &dt
}

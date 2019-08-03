package repo

import (
	"fmt"
	"testing"
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/table"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUrlSql_GetByAlias(t *testing.T) {

	testCases := []struct {
		name        string
		tableRows   *sqlmock.Rows
		alias       string
		hasErr      bool
		expectedUrl entity.Url
	}{
		{
			name: "alias not found",
			tableRows: sqlmock.NewRows([]string{
				table.Url.ColumnAlias,
				table.Url.ColumnOriginalUrl,
				table.Url.ColumnExpireAt,
				table.Url.ColumnCreatedAt,
				table.Url.ColumnUpdatedAt,
			}),
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name: "found url",
			tableRows: sqlmock.NewRows([]string{
				table.Url.ColumnAlias,
				table.Url.ColumnOriginalUrl,
				table.Url.ColumnExpireAt,
				table.Url.ColumnCreatedAt,
				table.Url.ColumnUpdatedAt,
			}).AddRow(
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

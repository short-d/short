package repo

import (
	"testing"
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
				table.Url.Alias,
				table.Url.OriginalUrl,
				table.Url.ExpireAt,
				table.Url.CreatedAt,
				table.Url.UpdatedAt,
			}),
			alias:  "220uFicCJj",
			hasErr: true,
		},
		{
			name: "found url",
			tableRows: sqlmock.NewRows([]string{
				table.Url.Alias,
				table.Url.OriginalUrl,
				table.Url.ExpireAt,
				table.Url.CreatedAt,
				table.Url.UpdatedAt,
			}).AddRow(
				"220uFicCJj",
				"http://www.google.com",
				"2019-05-01 08:02:16",
				"2017-05-01 08:02:16",
				"NULL",
			).AddRow(
				"yDOBcj5HIPbUAsw",
				"http://www.facebook.com",
				"2018-04-02 08:02:16",
				"2017-05-01 08:02:16",
				"NULL",
			),
			alias:  "220uFicCJj",
			hasErr: false,
			expectedUrl: entity.Url{
				Alias:       "220uFicCJj",
				OriginalUrl: "http://www.google.com",
				ExpireAt:    MustParseDatetime("2019-05-01 08:02:16"),
				CreatedAt:   MustParseDatetime("2017-05-01 08:02:16"),
				UpdatedAt:   MustParseDatetime("NULL"),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()

			assert.Nil(t, err)
			defer db.Close()

			mock.ExpectQuery("^SELECT .+ FROM Url WHERE alias=.+$").WillReturnRows(testCase.tableRows)

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

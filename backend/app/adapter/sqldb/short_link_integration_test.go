// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/fw/must"
	"github.com/short-d/short/backend/app/fw/ptr"
)

var insertShortLinkRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
	table.ShortLink.TableName,
	table.ShortLink.ColumnAlias,
	table.ShortLink.ColumnLongLink,
	table.ShortLink.ColumnCreatedAt,
	table.ShortLink.ColumnExpireAt,
	table.ShortLink.ColumnUpdatedAt,
	table.ShortLink.ColumnOpenGraphTitle,
	table.ShortLink.ColumnOpenGraphDescription,
	table.ShortLink.ColumnOpenGraphImageURL,
	table.ShortLink.ColumnTwitterTitle,
	table.ShortLink.ColumnTwitterDescription,
	table.ShortLink.ColumnTwitterImageURL,
)

type shortLinkTableRow struct {
	alias              string
	longLink           string
	createdAt          *time.Time
	expireAt           *time.Time
	updatedAt          *time.Time
	ogTitle            *string
	ogDescription      *string
	ogImageURL         *string
	twitterTitle       *string
	twitterDescription *string
	twitterImageURL    *string
}

func TestShortLinkSql_UpdateOGMetaTags(t *testing.T) {
	testCases := []struct {
		name              string
		tableRows         []shortLinkTableRow
		alias             string
		metaTags          metatag.OpenGraph
		expectedShortLink entity.ShortLink
	}{
		{
			name: "Twitter tags not provided",
			tableRows: []shortLinkTableRow{
				{
					alias:    "220uFicCJj",
					longLink: "http://www.google.com",
				},
			},
			alias: "220uFicCJj",
			metaTags: metatag.OpenGraph{
				Title:       ptr.String("title1"),
				Description: ptr.String("description1"),
				ImageURL:    ptr.String("url1"),
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
			},
		},
		{
			name: "Twitter tags provided",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
			},
			alias: "220uFicCJj",
			metaTags: metatag.OpenGraph{
				Title:       ptr.String("title2"),
				Description: ptr.String("description2"),
				ImageURL:    ptr.String("url2"),
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title2"),
					Description: ptr.String("description2"),
					ImageURL:    ptr.String("url2"),
				},
				TwitterTags: metatag.Twitter{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)

					shortLink, err := shortLinkRepo.UpdateOpenGraphTags(testCase.alias, testCase.metaTags)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedShortLink, shortLink)
				})
		})
	}
}

func TestShortLinkSql_UpdateTwitterTags(t *testing.T) {
	testCases := []struct {
		name              string
		tableRows         []shortLinkTableRow
		alias             string
		metaTags          metatag.Twitter
		expectedShortLink entity.ShortLink
	}{
		{
			name: "Twitter tags not provided",
			tableRows: []shortLinkTableRow{
				{
					alias:         "220uFicCJj",
					longLink:      "http://www.google.com",
					ogTitle:       ptr.String("title1"),
					ogDescription: ptr.String("description1"),
					ogImageURL:    ptr.String("url1"),
				},
			},
			alias: "220uFicCJj",
			metaTags: metatag.Twitter{
				Title:       ptr.String("title2"),
				Description: ptr.String("description2"),
				ImageURL:    ptr.String("url2"),
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
				TwitterTags: metatag.Twitter{
					Title:       ptr.String("title2"),
					Description: ptr.String("description2"),
					ImageURL:    ptr.String("url2"),
				},
			},
		},
		{
			name: "Twitter tags provided",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
			},
			alias: "220uFicCJj",
			metaTags: metatag.Twitter{
				Title:       ptr.String("title2"),
				Description: ptr.String("description2"),
				ImageURL:    ptr.String("url2"),
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
				TwitterTags: metatag.Twitter{
					Title:       ptr.String("title2"),
					Description: ptr.String("description2"),
					ImageURL:    ptr.String("url2"),
				},
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)

					shortLink, err := shortLinkRepo.UpdateTwitterTags(testCase.alias, testCase.metaTags)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedShortLink, shortLink)
				})
		})
	}
}

func TestShortLinkSql_IsAliasExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  []shortLinkTableRow
		alias      string
		expIsExist bool
	}{
		{
			name:       "alias doesn't exist",
			alias:      "gg",
			tableRows:  []shortLinkTableRow{},
			expIsExist: false,
		},
		{
			name:  "alias found",
			alias: "gg",
			tableRows: []shortLinkTableRow{
				{alias: "gg"},
			},
			expIsExist: true,
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					gotIsExist, err := shortLinkRepo.IsAliasExist(testCase.alias)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expIsExist, gotIsExist)
				})
		})
	}
}

func TestShortLinkSql_GetShortLinkByAlias(t *testing.T) {
	testCases := []struct {
		name              string
		tableRows         []shortLinkTableRow
		alias             string
		hasErr            bool
		expectedShortLink entity.ShortLink
	}{
		{
			name:      "alias not found",
			tableRows: []shortLinkTableRow{},
			alias:     "220uFicCJj",
			hasErr:    true,
		},
		{
			name: "found short link",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					createdAt:          ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					updatedAt:          ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					updatedAt:          ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title2"),
					ogDescription:      ptr.String("description2"),
					ogImageURL:         ptr.String("url2"),
					twitterTitle:       ptr.String("title2"),
					twitterDescription: ptr.String("description2"),
					twitterImageURL:    ptr.String("url2"),
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "220uFicCJj",
				LongLink:  "http://www.google.com",
				CreatedAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
				ExpireAt:  ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
				UpdatedAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
				TwitterTags: metatag.Twitter{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
			},
		},
		{
			name: "nil time",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					createdAt:          nil,
					expireAt:           nil,
					updatedAt:          nil,
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					updatedAt:          ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title2"),
					ogDescription:      ptr.String("description2"),
					ogImageURL:         ptr.String("url2"),
					twitterTitle:       ptr.String("title2"),
					twitterDescription: ptr.String("description2"),
					twitterImageURL:    ptr.String("url2"),
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "220uFicCJj",
				LongLink:  "http://www.google.com",
				CreatedAt: nil,
				ExpireAt:  nil,
				UpdatedAt: nil,
				OpenGraphTags: metatag.OpenGraph{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
				TwitterTags: metatag.Twitter{
					Title:       ptr.String("title1"),
					Description: ptr.String("description1"),
					ImageURL:    ptr.String("url1"),
				},
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					shortLink, err := shortLinkRepo.GetShortLinkByAlias(testCase.alias)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedShortLink, shortLink)
				},
			)
		})
	}
}

func TestShortLinkSql_CreateShortLink(t *testing.T) {
	testCases := []struct {
		name           string
		tableRows      []shortLinkTableRow
		shortLinkInput entity.ShortLinkInput
		hasErr         bool
	}{
		{
			name: "alias exists",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.facebook.com",
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
			},
			shortLinkInput: entity.ShortLinkInput{
				CustomAlias: ptr.String("220uFicCJj"),
				LongLink:    ptr.String("http://www.google.com"),
				ExpireAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
				CreatedAt:   ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
			},
			hasErr: true,
		},
		{
			name: "successfully create short link",
			tableRows: []shortLinkTableRow{
				{
					alias:              "abc",
					longLink:           "http://www.google.com",
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
			},
			shortLinkInput: entity.ShortLinkInput{
				CustomAlias: ptr.String("220uFicCJj"),
				LongLink:    ptr.String("http://www.google.com"),
				ExpireAt:    ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
				CreatedAt:   ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
			},
			hasErr: false,
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					err := shortLinkRepo.CreateShortLink(testCase.shortLinkInput)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)

					shortLink, err := shortLinkRepo.GetShortLinkByAlias(testCase.shortLinkInput.GetCustomAlias(""))
					assert.Equal(t, nil, err)
					assert.Equal(t, *testCase.shortLinkInput.CustomAlias, shortLink.Alias)
					assert.Equal(t, *testCase.shortLinkInput.LongLink, shortLink.LongLink)
					assert.Equal(t, testCase.shortLinkInput.ExpireAt, shortLink.ExpireAt)
					assert.Equal(t, testCase.shortLinkInput.CreatedAt, shortLink.CreatedAt)
				},
			)
		})
	}
}

func TestShortLinkSql_UpdateShortLink(t *testing.T) {
	testCases := []struct {
		name              string
		oldAlias          string
		shortLinkInput    entity.ShortLinkInput
		tableRows         []shortLinkTableRow
		hasErr            bool
		expectedShortLink entity.ShortLink
	}{
		{
			name:     "alias not found",
			oldAlias: "does_not_exist",
			tableRows: []shortLinkTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
				},
			},
			hasErr:            true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:     "alias is taken",
			oldAlias: "220uFicCja",
			tableRows: []shortLinkTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
				},
				{
					alias:     "efpIZ4OS",
					longLink:  "https://gmail.com",
					createdAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
				},
			},
			hasErr:            true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:     "valid new alias",
			oldAlias: "220uFicCJj",
			shortLinkInput: entity.ShortLinkInput{
				CustomAlias: ptr.String("GxtKXM9V"),
				LongLink:    ptr.String("https://www.google.com"),
				UpdatedAt:   ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
			},
			tableRows: []shortLinkTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
				},
			},
			hasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "GxtKXM9V",
				LongLink:  "https://www.google.com",
				UpdatedAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)
					expectedShortLink := testCase.expectedShortLink

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					shortLink, err := shortLinkRepo.UpdateShortLink(
						testCase.oldAlias,
						testCase.shortLinkInput,
					)
					assert.Equal(t, nil, err)

					shortLink, err = shortLinkRepo.GetShortLinkByAlias(testCase.shortLinkInput.GetCustomAlias(""))
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, expectedShortLink.Alias, shortLink.Alias)
					assert.Equal(t, expectedShortLink.LongLink, shortLink.LongLink)
					assert.Equal(t, expectedShortLink.ExpireAt, shortLink.ExpireAt)
					assert.Equal(t, expectedShortLink.UpdatedAt, shortLink.UpdatedAt)
				},
			)
		})
	}
}

func TestShortLinkSql_GetShortLinkByAliases(t *testing.T) {
	testCases := []struct {
		name               string
		tableRows          []shortLinkTableRow
		aliases            []string
		hasErr             bool
		expectedShortLinks []entity.ShortLink
	}{
		{
			name:      "alias not found",
			tableRows: []shortLinkTableRow{},
			aliases:   []string{"220uFicCJj"},
			hasErr:    false,
		},
		{
			name: "found short link",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					createdAt:          ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					updatedAt:          ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title1"),
					ogDescription:      ptr.String("description1"),
					ogImageURL:         ptr.String("url1"),
					twitterTitle:       ptr.String("title1"),
					twitterDescription: ptr.String("description1"),
					twitterImageURL:    ptr.String("url1"),
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					expireAt:           ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					updatedAt:          ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					ogTitle:            ptr.String("title2"),
					ogDescription:      ptr.String("description2"),
					ogImageURL:         ptr.String("url2"),
					twitterTitle:       ptr.String("title2"),
					twitterDescription: ptr.String("description2"),
					twitterImageURL:    ptr.String("url2"),
				},
			},
			aliases: []string{"220uFicCJj", "yDOBcj5HIPbUAsw"},
			hasErr:  false,
			expectedShortLinks: []entity.ShortLink{
				{
					Alias:     "220uFicCJj",
					LongLink:  "http://www.google.com",
					CreatedAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					ExpireAt:  ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					UpdatedAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("url1"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("url1"),
					},
				},
				{
					Alias:     "yDOBcj5HIPbUAsw",
					LongLink:  "http://www.facebook.com",
					CreatedAt: ptr.Time(must.Time(t, "2017-05-01T08:02:16-07:00")),
					ExpireAt:  ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					UpdatedAt: ptr.Time(must.Time(t, "2019-05-01T08:02:16-07:00")),
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("url2"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("url2"),
					},
				},
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					shortLink, err := shortLinkRepo.GetShortLinksByAliases(testCase.aliases)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedShortLinks, shortLink)
				},
			)
		})
	}
}

func TestShortLinkSql_DeleteShortLink(t *testing.T) {
	var (
		twoYearsAgo = mustParseTime(t, "2018-05-01T08:02:16-07:00")
		now         = mustParseTime(t, "2020-05-01T08:02:16-07:00")

		longLink    = "https://short-d.com/"
		customAlias = "short_is_great"
	)

	testCases := []struct {
		name      string
		tableRows []shortLinkTableRow
		input     entity.ShortLinkInput
		hasErr    bool
	}{
		{
			name: "delete exisiting shortlink",
			tableRows: []shortLinkTableRow{
				{
					alias:              "short_is_great",
					longLink:           "https://short-d.com",
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					ogTitle:            nil,
					ogDescription:      nil,
					ogImageURL:         nil,
					twitterTitle:       nil,
					twitterDescription: nil,
					twitterImageURL:    nil,
				},
			},
			input: entity.ShortLinkInput{
				LongLink:    &longLink,
				CustomAlias: &customAlias,
				CreatedAt:   nil,
			},
			hasErr: false,
		},
		{
			name: "shortlink does not exist",
			tableRows: []shortLinkTableRow{
				{
					alias:              "i_luv_short",
					longLink:           "https://short-d.com",
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					ogTitle:            nil,
					ogDescription:      nil,
					ogImageURL:         nil,
					twitterTitle:       nil,
					twitterDescription: nil,
					twitterImageURL:    nil,
				},
			},
			input: entity.ShortLinkInput{
				LongLink:    &longLink,
				CustomAlias: &customAlias,
				CreatedAt:   nil,
			},
			hasErr: true,
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
					insertShortLinkTableRows(t, sqlDB, testCase.tableRows)

					shortLinkRepo := sqldb.NewShortLinkSQL(sqlDB)
					err := shortLinkRepo.DeleteShortLink(testCase.input)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
					}
					assert.Equal(t, nil, err)
				},
			)
		})
	}
}

func insertShortLinkTableRows(t *testing.T, sqlDB *sql.DB, tableRows []shortLinkTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertShortLinkRowSQL,
			tableRow.alias,
			tableRow.longLink,
			tableRow.createdAt,
			tableRow.expireAt,
			tableRow.updatedAt,
			tableRow.ogTitle,
			tableRow.ogDescription,
			tableRow.ogImageURL,
			tableRow.twitterTitle,
			tableRow.twitterDescription,
			tableRow.twitterImageURL,
		)
		assert.Equal(t, nil, err)
	}
}

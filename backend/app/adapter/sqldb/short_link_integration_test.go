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
	title1 := "title1"
	description1 := "description1"
	imageURL1 := "url1"
	title2 := "title2"
	description2 := "description2"
	imageURL2 := "url2"

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
				Title:       &title1,
				Description: &description1,
				ImageURL:    &imageURL1,
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
				},
			},
		},
		{
			name: "Twitter tags provided",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.google.com",
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
			},
			alias: "220uFicCJj",
			metaTags: metatag.OpenGraph{
				Title:       &title2,
				Description: &description2,
				ImageURL:    &imageURL2,
			},
			expectedShortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				OpenGraphTags: metatag.OpenGraph{
					Title:       &title2,
					Description: &description2,
					ImageURL:    &imageURL2,
				},
				TwitterTags: metatag.Twitter{
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)

					shortLink, err := shortLinkRepo.UpdateOpenGraphTags(testCase.alias, testCase.metaTags)
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)
					gotIsExist, err := shortLinkRepo.IsAliasExist(testCase.alias)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expIsExist, gotIsExist)
				})
		})
	}
}

func TestShortLinkSql_GetShortLinkByAlias(t *testing.T) {
	twoYearsAgo := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")
	title1 := "title1"
	description1 := "description1"
	imageURL1 := "url1"
	title2 := "title2"
	description2 := "description2"
	imageURL2 := "url2"

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
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					updatedAt:          &now,
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					updatedAt:          &now,
					ogTitle:            &title2,
					ogDescription:      &description2,
					ogImageURL:         &imageURL2,
					twitterTitle:       &title2,
					twitterDescription: &description2,
					twitterImageURL:    &imageURL2,
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "220uFicCJj",
				LongLink:  "http://www.google.com",
				CreatedAt: &twoYearsAgo,
				ExpireAt:  &now,
				UpdatedAt: &now,
				OpenGraphTags: metatag.OpenGraph{
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
				},
				TwitterTags: metatag.Twitter{
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
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
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					updatedAt:          &now,
					ogTitle:            &title2,
					ogDescription:      &description2,
					ogImageURL:         &imageURL2,
					twitterTitle:       &title2,
					twitterDescription: &description2,
					twitterImageURL:    &imageURL2,
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
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
				},
				TwitterTags: metatag.Twitter{
					Title:       &title1,
					Description: &description1,
					ImageURL:    &imageURL1,
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)
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
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")
	title1 := "title1"
	description1 := "description1"
	imageURL1 := "url1"
	title2 := "title2"
	description2 := "description2"
	imageURL2 := "url2"

	testCases := []struct {
		name      string
		tableRows []shortLinkTableRow
		shortLink entity.ShortLink
		hasErr    bool
	}{
		{
			name: "alias exists",
			tableRows: []shortLinkTableRow{
				{
					alias:              "220uFicCJj",
					longLink:           "http://www.facebook.com",
					expireAt:           &now,
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
			},
			shortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				ExpireAt: &now,
				OpenGraphTags: metatag.OpenGraph{
					Title:       &title2,
					Description: &description2,
					ImageURL:    &imageURL2,
				},
				TwitterTags: metatag.Twitter{
					Title:       &title2,
					Description: &description2,
					ImageURL:    &imageURL2,
				},
			},
			hasErr: true,
		},
		{
			name: "successfully create short link",
			tableRows: []shortLinkTableRow{
				{
					alias:              "abc",
					longLink:           "http://www.google.com",
					expireAt:           &now,
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
			},
			shortLink: entity.ShortLink{
				Alias:    "220uFicCJj",
				LongLink: "http://www.google.com",
				ExpireAt: &now,
				OpenGraphTags: metatag.OpenGraph{
					Title:       &title2,
					Description: &description2,
					ImageURL:    &imageURL2,
				},
				TwitterTags: metatag.Twitter{
					Title:       &title2,
					Description: &description2,
					ImageURL:    &imageURL2,
				},
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)
					err := shortLinkRepo.CreateShortLink(testCase.shortLink)

					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)

					shortLink, err := shortLinkRepo.GetShortLinkByAlias(testCase.shortLink.Alias)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.shortLink, shortLink)
				},
			)
		})
	}
}

func TestShortLinkSql_UpdateShortLink(t *testing.T) {
	createdAt := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := mustParseTime(t, "2020-05-01T08:02:16-07:00")

	testCases := []struct {
		name              string
		oldAlias          string
		newShortLink      entity.ShortLink
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
					createdAt: &createdAt,
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
					createdAt: &createdAt,
				},
				{
					alias:     "efpIZ4OS",
					longLink:  "https://gmail.com",
					createdAt: &createdAt,
				},
			},
			hasErr:            true,
			expectedShortLink: entity.ShortLink{},
		},
		{
			name:     "valid new alias",
			oldAlias: "220uFicCJj",
			newShortLink: entity.ShortLink{
				Alias:     "GxtKXM9V",
				LongLink:  "https://www.google.com",
				UpdatedAt: &now,
			},
			tableRows: []shortLinkTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: &createdAt,
				},
			},
			hasErr: false,
			expectedShortLink: entity.ShortLink{
				Alias:     "GxtKXM9V",
				LongLink:  "https://www.google.com",
				UpdatedAt: &now,
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)
					shortLink, err := shortLinkRepo.UpdateShortLink(
						testCase.oldAlias,
						testCase.newShortLink,
					)
					assert.Equal(t, nil, err)

					shortLink, err = shortLinkRepo.GetShortLinkByAlias(testCase.newShortLink.Alias)
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
	twoYearsAgo := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")
	title1 := "title1"
	description1 := "description1"
	imageURL1 := "url1"
	title2 := "title2"
	description2 := "description2"
	imageURL2 := "url2"

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
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					updatedAt:          &now,
					ogTitle:            &title1,
					ogDescription:      &description1,
					ogImageURL:         &imageURL1,
					twitterTitle:       &title1,
					twitterDescription: &description1,
					twitterImageURL:    &imageURL1,
				},
				{
					alias:              "yDOBcj5HIPbUAsw",
					longLink:           "http://www.facebook.com",
					createdAt:          &twoYearsAgo,
					expireAt:           &now,
					updatedAt:          &now,
					ogTitle:            &title2,
					ogDescription:      &description2,
					ogImageURL:         &imageURL2,
					twitterTitle:       &title2,
					twitterDescription: &description2,
					twitterImageURL:    &imageURL2,
				},
			},
			aliases: []string{"220uFicCJj", "yDOBcj5HIPbUAsw"},
			hasErr:  false,
			expectedShortLinks: []entity.ShortLink{
				{
					Alias:     "220uFicCJj",
					LongLink:  "http://www.google.com",
					CreatedAt: &twoYearsAgo,
					ExpireAt:  &now,
					UpdatedAt: &now,
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
				},
				{
					Alias:     "yDOBcj5HIPbUAsw",
					LongLink:  "http://www.facebook.com",
					CreatedAt: &twoYearsAgo,
					ExpireAt:  &now,
					UpdatedAt: &now,
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
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

					shortLinkRepo := sqldb.NewShortLinkSql(sqlDB)
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

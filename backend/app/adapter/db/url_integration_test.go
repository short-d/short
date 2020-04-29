// +build integration all

package db_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
)

var insertURLRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s, %s, %s)
VALUES ($1, $2, $3, $4, $5)`,
	table.URL.TableName,
	table.URL.ColumnAlias,
	table.URL.ColumnOriginalURL,
	table.URL.ColumnCreatedAt,
	table.URL.ColumnExpireAt,
	table.URL.ColumnUpdatedAt,
)

type urlTableRow struct {
	alias     string
	longLink  string
	createdAt *time.Time
	expireAt  *time.Time
	updatedAt *time.Time
}

func TestURLSql_IsAliasExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  []urlTableRow
		alias      string
		expIsExist bool
	}{
		{
			name:       "alias doesn't exist",
			alias:      "gg",
			tableRows:  []urlTableRow{},
			expIsExist: false,
		},
		{
			name:  "alias found",
			alias: "gg",
			tableRows: []urlTableRow{
				{alias: "gg"},
			},
			expIsExist: true,
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
					insertURLTableRows(t, sqlDB, testCase.tableRows)

					urlRepo := db.NewURLSql(sqlDB)
					gotIsExist, err := urlRepo.IsAliasExist(testCase.alias)
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expIsExist, gotIsExist)
				})
		})
	}
}

func TestURLSql_GetByAlias(t *testing.T) {
	twoYearsAgo := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")

	testCases := []struct {
		name        string
		tableRows   []urlTableRow
		alias       string
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name:      "alias not found",
			tableRows: []urlTableRow{},
			alias:     "220uFicCJj",
			hasErr:    true,
		},
		{
			name: "found url",
			tableRows: []urlTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "http://www.google.com",
					createdAt: &twoYearsAgo,
					expireAt:  &now,
					updatedAt: &now,
				},
				{
					alias:     "yDOBcj5HIPbUAsw",
					longLink:  "http://www.facebook.com",
					createdAt: &twoYearsAgo,
					expireAt:  &now,
					updatedAt: &now,
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedURL: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				CreatedAt:   &twoYearsAgo,
				ExpireAt:    &now,
				UpdatedAt:   &now,
			},
		},
		{
			name: "nil time",
			tableRows: []urlTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "http://www.google.com",
					createdAt: nil,
					expireAt:  nil,
					updatedAt: nil,
				},
				{
					alias:     "yDOBcj5HIPbUAsw",
					longLink:  "http://www.facebook.com",
					createdAt: &twoYearsAgo,
					expireAt:  &now,
					updatedAt: &now,
				},
			},
			alias:  "220uFicCJj",
			hasErr: false,
			expectedURL: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				CreatedAt:   nil,
				ExpireAt:    nil,
				UpdatedAt:   nil,
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
					insertURLTableRows(t, sqlDB, testCase.tableRows)

					urlRepo := db.NewURLSql(sqlDB)
					url, err := urlRepo.GetByAlias(testCase.alias)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedURL, url)
				},
			)
		})
	}
}

// TODO(issue#698): change to TestURLSql_CreateURL
func TestURLSql_Create(t *testing.T) {
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")

	testCases := []struct {
		name      string
		tableRows []urlTableRow
		url       entity.URL
		hasErr    bool
	}{
		{
			name: "alias exists",
			tableRows: []urlTableRow{
				{
					alias:    "220uFicCJj",
					longLink: "http://www.facebook.com",
					expireAt: &now,
				},
			},
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    &now,
			},
			hasErr: true,
		},
		{
			name: "successfully create url",
			tableRows: []urlTableRow{
				{
					alias:    "abc",
					longLink: "http://www.google.com",
					expireAt: &now,
				},
			},
			url: entity.URL{
				Alias:       "220uFicCJj",
				OriginalURL: "http://www.google.com",
				ExpireAt:    &now,
			},
			hasErr: false,
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
					insertURLTableRows(t, sqlDB, testCase.tableRows)

					urlRepo := db.NewURLSql(sqlDB)
					err := urlRepo.Create(testCase.url)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
				},
			)
		})
	}
}

func TestURLSql_UpdateURL(t *testing.T) {
	createdAt := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := time.Now()

	testCases := []struct {
		name        string
		oldAlias    string
		newURL      entity.URL
		tableRows   []urlTableRow
		hasErr      bool
		expectedURL entity.URL
	}{
		{
			name:     "alias not found",
			oldAlias: "does_not_exist",
			tableRows: []urlTableRow{
				urlTableRow{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: &createdAt,
				},
			},
			hasErr:      true,
			expectedURL: entity.URL{},
		},
		{
			name:     "alias is taken",
			oldAlias: "220uFicCja",
			tableRows: []urlTableRow{
				urlTableRow{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: &createdAt,
				},
				urlTableRow{
					alias:     "efpIZ4OS",
					longLink:  "https://gmail.com",
					createdAt: &createdAt,
				},
			},
			hasErr:      true,
			expectedURL: entity.URL{},
		},
		{
			name:     "valid new alias",
			oldAlias: "220uFicCJj",
			newURL: entity.URL{
				Alias:       "GxtKXM9V",
				OriginalURL: "https://www.google.com",
				UpdatedAt:   &now,
			},
			tableRows: []urlTableRow{
				urlTableRow{
					alias:     "220uFicCJj",
					longLink:  "https://www.google.com",
					createdAt: &createdAt,
				},
			},
			hasErr: false,
			expectedURL: entity.URL{
				Alias:       "GxtKXM9V",
				OriginalURL: "https://www.google.com",
				UpdatedAt:   &now,
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
					insertURLTableRows(t, sqlDB, testCase.tableRows)
					expectedURL := testCase.expectedURL

					urlRepo := db.NewURLSql(sqlDB)
					url, err := urlRepo.UpdateURL(
						testCase.oldAlias,
						testCase.newURL,
					)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, expectedURL.Alias, url.Alias)
					mdtest.Equal(t, expectedURL.OriginalURL, url.OriginalURL)
					mdtest.Equal(t, expectedURL.ExpireAt, url.ExpireAt)
					mdtest.Equal(t, expectedURL.UpdatedAt, url.UpdatedAt)
				},
			)
		})
	}
}

func TestURLSql_GetByAliases(t *testing.T) {
	twoYearsAgo := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	now := mustParseTime(t, "2019-05-01T08:02:16-07:00")

	testCases := []struct {
		name         string
		tableRows    []urlTableRow
		aliases      []string
		hasErr       bool
		expectedURLs []entity.URL
	}{
		{
			name:      "alias not found",
			tableRows: []urlTableRow{},
			aliases:   []string{"220uFicCJj"},
			hasErr:    false,
		},
		{
			name: "found url",
			tableRows: []urlTableRow{
				{
					alias:     "220uFicCJj",
					longLink:  "http://www.google.com",
					createdAt: &twoYearsAgo,
					expireAt:  &now,
					updatedAt: &now,
				},
				{
					alias:     "yDOBcj5HIPbUAsw",
					longLink:  "http://www.facebook.com",
					createdAt: &twoYearsAgo,
					expireAt:  &now,
					updatedAt: &now,
				},
			},
			aliases: []string{"220uFicCJj", "yDOBcj5HIPbUAsw"},
			hasErr:  false,
			expectedURLs: []entity.URL{
				{
					Alias:       "220uFicCJj",
					OriginalURL: "http://www.google.com",
					CreatedAt:   &twoYearsAgo,
					ExpireAt:    &now,
					UpdatedAt:   &now,
				},
				{
					Alias:       "yDOBcj5HIPbUAsw",
					OriginalURL: "http://www.facebook.com",
					CreatedAt:   &twoYearsAgo,
					ExpireAt:    &now,
					UpdatedAt:   &now,
				},
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
					insertURLTableRows(t, sqlDB, testCase.tableRows)

					urlRepo := db.NewURLSql(sqlDB)
					urls, err := urlRepo.GetByAliases(testCase.aliases)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedURLs, urls)
				},
			)
		})
	}
}

func insertURLTableRows(t *testing.T, sqlDB *sql.DB, tableRows []urlTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertURLRowSQL,
			tableRow.alias,
			tableRow.longLink,
			tableRow.createdAt,
			tableRow.expireAt,
			tableRow.updatedAt,
		)
		mdtest.Equal(t, nil, err)
	}
}

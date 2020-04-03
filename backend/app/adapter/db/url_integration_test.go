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

func TestURLSql_FindURLsByUser(t *testing.T) {
	testCases := []struct {
		name                     string
		urlTableRows             []urlTableRow
		userURLRelationTableRows []userURLRelationTableRow
		userTableRows            []userTableRow
		user                     entity.User
		expectedURLs             []entity.URL
	}{
		{
			name: "URL's present for given user",
			urlTableRows: []urlTableRow{
				{
					alias:    "google",
					longLink: "https://www.google.com/",
				},
				{
					alias:    "short",
					longLink: "https://github.com/short-d/short/",
				},
				{
					alias:    "mozilla",
					longLink: "https://www.mozilla.org/",
				},
			},
			userURLRelationTableRows: []userURLRelationTableRow{
				{
					alias:     "google",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "short",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "mozilla",
					userEmail: "test2@gmail.com",
				},
			},
			userTableRows: []userTableRow{
				{
					id:    "12345",
					email: "test@gmail.com",
					name:  "Test User",
				},
				{
					id:    "12346",
					email: "test2@gmail.com",
					name:  "Test User",
				},
			},
			user: entity.User{
				ID:    "test",
				Name:  "test user",
				Email: "test@gmail.com",
			},
			expectedURLs: []entity.URL{
				{
					Alias:       "google",
					OriginalURL: "https://www.google.com/",
				},
				{
					Alias:       "short",
					OriginalURL: "https://github.com/short-d/short/",
				},
			},
		}, {
			name: "No URL present for given user",
			urlTableRows: []urlTableRow{
				{
					alias:    "google",
					longLink: "https://www.google.com/",
				},
				{
					alias:    "short",
					longLink: "https://github.com/short-d/short/",
				},
				{
					alias:    "mozilla",
					longLink: "https://www.mozilla.org/",
				},
			},
			userURLRelationTableRows: []userURLRelationTableRow{
				{
					alias:     "google",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "short",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "mozilla",
					userEmail: "test@gmail.com",
				},
			},
			userTableRows: []userTableRow{
				{
					id:    "12345",
					email: "test@gmail.com",
					name:  "Test User",
				},
				{
					id:    "12346",
					email: "test2@gmail.com",
					name:  "Test User",
				},
			},
			user: entity.User{
				ID:    "test",
				Name:  "test user",
				Email: "test2@gmail.com",
			},
			expectedURLs: []entity.URL{},
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertURLTableRows(t, sqlDB, testCase.urlTableRows)
					insertUserURLRelationTableRows(t, sqlDB, testCase.userURLRelationTableRows)

					urlRepo := db.NewURLSql(sqlDB)
					urls, err := urlRepo.FindURLsByUser(testCase.user)

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedURLs, urls)
				},
			)
		})
	}
}

func TestURLSql_DeleteURLByUser(t *testing.T) {
	testCases := []struct {
		name                     string
		urlTableRows             []urlTableRow
		userURLRelationTableRows []userURLRelationTableRow
		userTableRows            []userTableRow
		user                     entity.User
		alias                    string
		hasErr                   bool
		expectedURLs             []entity.URL
		deletedURL               entity.URL
	}{
		{
			name: "Delete alias created by given user",
			urlTableRows: []urlTableRow{
				{
					alias:    "google",
					longLink: "https://www.google.com/",
				},
				{
					alias:    "short",
					longLink: "https://github.com/short-d/short/",
				},
				{
					alias:    "mozilla",
					longLink: "https://www.mozilla.org/",
				},
			},
			userURLRelationTableRows: []userURLRelationTableRow{
				{
					alias:     "google",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "short",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "mozilla",
					userEmail: "test2@gmail.com",
				},
			},
			userTableRows: []userTableRow{
				{
					id:    "12345",
					email: "test@gmail.com",
					name:  "Test User",
				},
				{
					id:    "12346",
					email: "test2@gmail.com",
					name:  "Test User",
				},
			},
			user: entity.User{
				ID:    "test",
				Name:  "test user",
				Email: "test@gmail.com",
			},
			alias:  "google",
			hasErr: false,
			expectedURLs: []entity.URL{
				{
					Alias:       "short",
					OriginalURL: "https://github.com/short-d/short/",
				},
			},
			deletedURL: entity.URL{
				Alias:       "google",
				OriginalURL: "https://www.google.com/",
			},
		},
		{
			name: "Cannot delete alias created by different user",
			urlTableRows: []urlTableRow{
				{
					alias:    "google",
					longLink: "https://www.google.com/",
				},
				{
					alias:    "short",
					longLink: "https://github.com/short-d/short/",
				},
				{
					alias:    "mozilla",
					longLink: "https://www.mozilla.org/",
				},
			},
			userURLRelationTableRows: []userURLRelationTableRow{
				{
					alias:     "google",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "short",
					userEmail: "test@gmail.com",
				},
				{
					alias:     "mozilla",
					userEmail: "test2@gmail.com",
				},
			},
			userTableRows: []userTableRow{
				{
					id:    "12345",
					email: "test@gmail.com",
					name:  "Test User",
				},
				{
					id:    "12346",
					email: "test2@gmail.com",
					name:  "Test User",
				},
			},
			user: entity.User{
				ID:    "test",
				Name:  "test user",
				Email: "test2@gmail.com",
			},
			alias:  "google",
			hasErr: true,
			expectedURLs: []entity.URL{
				{
					Alias:       "google",
					OriginalURL: "https://www.google.com/",
				},
				{
					Alias:       "short",
					OriginalURL: "https://github.com/short-d/short/",
				},
				{
					Alias:       "mozilla",
					OriginalURL: "https://www.mozilla.org/",
				},
			},
			deletedURL: entity.URL{},
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertURLTableRows(t, sqlDB, testCase.urlTableRows)
					insertUserURLRelationTableRows(t, sqlDB, testCase.userURLRelationTableRows)

					urlRepo := db.NewURLSql(sqlDB)
					url, err := urlRepo.DeleteURLByUser(testCase.user, testCase.alias)

					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					} else {
						mdtest.Equal(t, nil, err)
						mdtest.Equal(t, testCase.deletedURL, url)
					}

					urls, err := urlRepo.FindURLsByUser(testCase.user)

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

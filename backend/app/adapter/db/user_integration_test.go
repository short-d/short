// +build integration

package db_test

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db"
	"short/app/adapter/db/table"
	"short/app/entity"
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"
)

var insertUserRowSQL = fmt.Sprintf(`
INSERT INTO "%s" (%s, %s, %s, %s, %s, %s)
VALUES ($1, $2, $3, $4, $5, $6)`,
	table.User.TableName,
	table.User.ColumnID,
	table.User.ColumnEmail,
	table.User.ColumnName,
	table.User.ColumnLastSignedInAt,
	table.User.ColumnCreatedAt,
	table.User.ColumnUpdatedAt,
)

type userTableRow struct {
	id           string
	email        string
	name         string
	lastSignedIn *time.Time
	createdAt    *time.Time
	updatedAt    *time.Time
}

func TestUserSql_IsEmailExist(t *testing.T) {
	testCases := []struct {
		name       string
		tableRows  []userTableRow
		email      string
		expIsExist bool
	}{
		{
			name:       "email doesn't exist",
			email:      "user@example.com",
			tableRows:  []userTableRow{},
			expIsExist: false,
		},
		{
			name:       "email found",
			email:      "user@example.com",
			tableRows:  []userTableRow{{email: "user@example.com"}},
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
					insertUserTableRows(t, sqlDB, testCase.tableRows)

					userRepo := db.NewUserSQL(sqlDB)
					gotIsExist, err := userRepo.IsEmailExist(testCase.email)
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expIsExist, gotIsExist)
				})
		})
	}
}

func TestUserSql_GetUserByEmail(t *testing.T) {
	twoYearsAgo := mustParseTime(t, "2017-05-01T08:02:16-07:00")

	testCases := []struct {
		name      string
		tableRows []userTableRow
		email     string
		hasErr    bool
		expUser   entity.User
	}{
		{
			name:      "email doesn't exist",
			email:     "alpha@example.com",
			tableRows: []userTableRow{},
			hasErr:    true,
		},
		{
			name:  "email found",
			email: "alpha@example.com",
			tableRows: []userTableRow{
				{
					id:           "alpha",
					email:        "alpha@example.com",
					name:         "Alpha",
					lastSignedIn: &twoYearsAgo,
					createdAt:    &twoYearsAgo,
					updatedAt:    &twoYearsAgo,
				},
			},
			hasErr: false,
			expUser: entity.User{
				ID:             "alpha",
				Name:           "Alpha",
				Email:          "alpha@example.com",
				LastSignedInAt: &twoYearsAgo,
				CreatedAt:      &twoYearsAgo,
				UpdatedAt:      &twoYearsAgo,
			},
		},
		{
			name:  "nil times",
			email: "alpha@example.com",
			tableRows: []userTableRow{
				{
					id:           "alpha",
					email:        "alpha@example.com",
					name:         "Alpha",
					lastSignedIn: nil,
					createdAt:    nil,
					updatedAt:    nil,
				},
			},
			hasErr: false,
			expUser: entity.User{
				ID:             "alpha",
				Name:           "Alpha",
				Email:          "alpha@example.com",
				LastSignedInAt: nil,
				CreatedAt:      nil,
				UpdatedAt:      nil,
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
					insertUserTableRows(t, sqlDB, testCase.tableRows)

					userRepo := db.NewUserSQL(sqlDB)
					gotUser, err := userRepo.GetUserByEmail(testCase.email)
					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expUser, gotUser)
				})
		})
	}
}

func TestUserSql_CreateUser(t *testing.T) {
	testCases := []struct {
		name      string
		user      entity.User
		tableRows []userTableRow
		hasErr    bool
	}{
		{
			name: "email exists",
			user: entity.User{
				Email: "alpha@example.com",
			},
			tableRows: []userTableRow{
				{email: "alpha@example.com"},
			},
			hasErr: true,
		},
		{
			name: "no given email",
			user: entity.User{
				Email: "user@example.com",
				Name:  "Test User",
			},
			tableRows: []userTableRow{},
			hasErr:    false,
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
					insertUserTableRows(t, sqlDB, testCase.tableRows)

					userRepo := db.NewUserSQL(sqlDB)

					err := userRepo.CreateUser(testCase.user)
					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
				})
		})
	}
}

func TestUserSQL_UpdateUserID(t *testing.T) {
	testCases := []struct {
		name      string
		email     string
		tableRows []userTableRow
		hasErr    bool
	}{
		{
			name:      "user not found",
			email:     "alpha@example.com",
			tableRows: []userTableRow{},
			hasErr:    true,
		},
		{
			name:  "update user ID successfully",
			email: "alpha@example.com",
			tableRows: []userTableRow{
				{email: "alpha@example.com"},
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
					insertUserTableRows(t, sqlDB, testCase.tableRows)

					userRepo := db.NewUserSQL(sqlDB)

					err := userRepo.UpdateUserID(testCase.email, "alpha")
					if testCase.hasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}
					mdtest.Equal(t, nil, err)
				})
		})
	}
}

func insertUserTableRows(t *testing.T, sqlDB *sql.DB, tableRows []userTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserRowSQL,
			tableRow.id,
			tableRow.email,
			tableRow.name,
			tableRow.lastSignedIn,
			tableRow.createdAt,
			tableRow.updatedAt,
		)
		mdtest.Equal(t, nil, err)
	}
}

// +build integration all

package db_test

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/short-d/short/app/adapter/db"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
)

var insertUserChangeLogRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s)
VALUES ($1, $2, $3)`,
	table.UserChangeLog.TableName,
	table.UserChangeLog.ColumnUserID,
	table.UserChangeLog.ColumnEmail,
	table.UserChangeLog.ColumnLastViewedAt,
)

type userChangeLogTableRow struct {
	userID       string
	email        string
	lastViewedAt time.Time
}

func TestUserChangeLogSQL_GetLastViewedAt(t *testing.T) {
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)

	testCases := []struct {
		name                   string
		userChangeLogTableRows []userChangeLogTableRow
		user                   entity.User
		expectedLastViewedAt   time.Time
	}{
		{
			name: "user last viewed at time is one month ago",
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					email:        "test2@gmail.com",
					lastViewedAt: now,
				},
				{
					userID:       "12345",
					email:        "test@gmail.com",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: monthAgo,
		},
		{
			name: "user does not have last viewed time",
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					email:        "test2@gmail.com",
					lastViewedAt: now,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: time.Time{},
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
					insertUserChangeLogTableRows(t, sqlDB, testCase.userChangeLogTableRows)

					timerFake := mdtest.NewTimerFake(now)
					userChangeLogRepo := db.NewUserChangeLogSql(sqlDB, timerFake)

					lastViewedAt, err := userChangeLogRepo.GetLastViewedAt(testCase.user)
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedLastViewedAt.Unix(), lastViewedAt.Unix())
				})
		})
	}
}

func TestUserChangeLogSQL_UpdateLastViewedAt(t *testing.T) {
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)

	testCases := []struct {
		name                   string
		userChangeLogTableRows []userChangeLogTableRow
		user                   entity.User
		expectedLastViewedAt   time.Time
	}{
		{
			name: "user does not have last viewed time",
			userChangeLogTableRows: []userChangeLogTableRow{},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: now,
		},
		{
			name: "user has last viewed time",
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12345",
					email:        "test@gmail.com",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: now,
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
					insertUserChangeLogTableRows(t, sqlDB, testCase.userChangeLogTableRows)

					timerFake := mdtest.NewTimerFake(now)
					userChangeLogRepo := db.NewUserChangeLogSql(sqlDB, timerFake)

					_, err := userChangeLogRepo.UpdateLastViewedAt(testCase.user)
					mdtest.Equal(t, nil, err)

					lastViewedAt, err := userChangeLogRepo.GetLastViewedAt(testCase.user)
					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedLastViewedAt.Unix(), lastViewedAt.Unix())
				})
		})
	}
}

func insertUserChangeLogTableRows(
	t *testing.T,
	sqlDB *sql.DB,
	tableRows []userChangeLogTableRow,
) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertUserChangeLogRowSQL,
			tableRow.userID,
			tableRow.email,
			tableRow.lastViewedAt,
		)
		mdtest.Equal(t, nil, err)
	}
}

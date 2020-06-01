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
)

var insertUserChangeLogRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s)
VALUES ($1, $2);`,
	table.UserChangeLog.TableName,
	table.UserChangeLog.ColumnUserID,
	table.UserChangeLog.ColumnLastViewedAt,
)

type userChangeLogTableRow struct {
	userID       string
	lastViewedAt time.Time
}

func TestUserChangeLogSQL_GetLastViewedAt(t *testing.T) {
	now := mustParseTime(t, "2020-04-04T08:02:16-07:00")
	monthAgo := now.AddDate(0, -1, 0)

	testCases := []struct {
		name                   string
		userTableRows          []userTableRow
		userChangeLogTableRows []userChangeLogTableRow
		user                   entity.User
		expectedLastViewedAt   time.Time
		hasErr                 bool
	}{
		{
			name: "user last viewed at time is one month ago",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					lastViewedAt: now,
				},
				{
					userID:       "12345",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: monthAgo,
			hasErr:               false,
		},
		{
			name: "user does not exist",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					lastViewedAt: now,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: time.Time{},
			hasErr:               true,
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertUserChangeLogTableRows(t, sqlDB, testCase.userChangeLogTableRows)

					userChangeLogRepo := sqldb.NewUserChangeLogSQL(sqlDB)

					lastViewedAt, err := userChangeLogRepo.GetLastViewedAt(testCase.user)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedLastViewedAt.UTC(), lastViewedAt)
				})
		})
	}
}

func TestUserChangeLogSQL_UpdateLastViewedAt(t *testing.T) {
	now := mustParseTime(t, "2020-04-04T08:02:16-07:00")
	monthAgo := now.AddDate(0, -1, 0)

	testCases := []struct {
		name                   string
		userTableRows          []userTableRow
		userChangeLogTableRows []userChangeLogTableRow
		user                   entity.User
		expectedLastViewedAt   time.Time
		hasErr                 bool
	}{
		{
			name: "user does not exist",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: time.Time{},
			hasErr:               true,
		},
		{
			name: "user has last viewed time",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12345",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: now,
			hasErr:               false,
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertUserChangeLogTableRows(t, sqlDB, testCase.userChangeLogTableRows)

					userChangeLogRepo := sqldb.NewUserChangeLogSQL(sqlDB)

					_, err := userChangeLogRepo.UpdateLastViewedAt(testCase.user, now)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)

					lastViewedAt, err := userChangeLogRepo.GetLastViewedAt(testCase.user)

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedLastViewedAt.UTC(), lastViewedAt)
				})
		})
	}
}

func TestUserChangeLogSQL_CreateRelation(t *testing.T) {
	now := mustParseTime(t, "2020-04-04T08:02:16-07:00")
	monthAgo := now.AddDate(0, -1, 0)
	twoMonthsAgo := monthAgo.AddDate(0, -1, 0)

	testCases := []struct {
		name                   string
		userTableRows          []userTableRow
		userChangeLogTableRows []userChangeLogTableRow
		user                   entity.User
		expectedLastViewedAt   time.Time
		hasErr                 bool
	}{
		{
			name: "user does not exist",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: now,
			hasErr:               false,
		},
		{
			name: "user already exists",
			userTableRows: []userTableRow{
				{
					id:    "12345",
					name:  "Test User",
					email: "test@gmail.com",
				},
				{
					id:    "12346",
					name:  "Test User 2",
					email: "test2@gmail.com",
				},
			},
			userChangeLogTableRows: []userChangeLogTableRow{
				{
					userID:       "12346",
					lastViewedAt: twoMonthsAgo,
				},
				{
					userID:       "12345",
					lastViewedAt: monthAgo,
				},
			},
			user: entity.User{
				ID:    "12345",
				Name:  "Test User",
				Email: "test@gmail.com",
			},
			expectedLastViewedAt: time.Time{},
			hasErr:               true,
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
					insertUserTableRows(t, sqlDB, testCase.userTableRows)
					insertUserChangeLogTableRows(t, sqlDB, testCase.userChangeLogTableRows)

					userChangeLogRepo := sqldb.NewUserChangeLogSQL(sqlDB)
					err := userChangeLogRepo.CreateRelation(testCase.user, now)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)

					lastViewedAt, err := userChangeLogRepo.GetLastViewedAt(testCase.user)
					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedLastViewedAt.UTC(), lastViewedAt)
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
			tableRow.lastViewedAt,
		)
		assert.Equal(t, nil, err)
	}
}

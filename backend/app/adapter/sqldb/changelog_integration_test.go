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
	"github.com/short-d/short/backend/app/fw/ptr"
)

var insertChangeLogRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s, %s) 
VALUES ($1, $2, $3, $4)`,
	table.ChangeLog.TableName,
	table.ChangeLog.ColumnID,
	table.ChangeLog.ColumnTitle,
	table.ChangeLog.ColumnSummaryMarkdown,
	table.ChangeLog.ColumnReleasedAt,
)

type changeLogTableRow struct {
	id              string
	title           string
	summaryMarkdown string
	releasedAt      time.Time
}

func TestChangeLogSql_GetChangeLog(t *testing.T) {
	testCases := []struct {
		name              string
		tableRows         []changeLogTableRow
		expectedChangeLog []entity.Change
	}{
		{
			name: "get full changelog",
			tableRows: []changeLogTableRow{
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				}, {
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
			},
			expectedChangeLog: []entity.Change{
				{
					ID:              "12346",
					Title:           "title 2",
					SummaryMarkdown: ptr.String("summary 2"),
				},
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: ptr.String("summary 1"),
				},
			},
		},
		{
			name:              "get empty changelog",
			tableRows:         []changeLogTableRow{},
			expectedChangeLog: []entity.Change{},
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
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := sqldb.NewChangeLogSQL(sqlDB)
					changeLog, err := changeLogRepo.GetChangeLog()

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedChangeLog, changeLog)
				})
		})
	}
}

func TestChangeLogSql_CreateChange(t *testing.T) {
	testCases := []struct {
		name                  string
		tableRows             []changeLogTableRow
		change                entity.Change
		expectedChangeLogSize int
		expectedChange        entity.Change
	}{
		{
			name: "create a change",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			change: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: ptr.String("summary 3"),
			},
			expectedChangeLogSize: 3,
			expectedChange: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: ptr.String("summary 3"),
			},
		}, {
			name: "create a change with nil summary",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			change: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: nil,
			},
			expectedChangeLogSize: 3,
			expectedChange: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: nil,
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
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := sqldb.NewChangeLogSQL(sqlDB)

					change, err := changeLogRepo.CreateChange(testCase.change)
					changeLog, _ := changeLogRepo.GetChangeLog()

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedChangeLogSize, len(changeLog))
					assert.Equal(t, testCase.expectedChange, change)
				})
		})
	}
}

func TestChangeLogSql_DeleteChange(t *testing.T) {

	testCases := []struct {
		name                  string
		tableRows             []changeLogTableRow
		deleteChangeId        string
		expectedChangeLogSize int
		expectedChangeLog     []entity.Change
	}{
		{
			name: "delete an existing change",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "67890",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			deleteChangeId:        "67890",
			expectedChangeLogSize: 1,
			expectedChangeLog: []entity.Change{
				{
					ID:              "12346",
					Title:           "title 1",
					SummaryMarkdown: ptr.String("summary 1"),
				},
			},
		}, {
			name: "delete a non existent change",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "67890",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			deleteChangeId:        "34567",
			expectedChangeLogSize: 2,
			expectedChangeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: ptr.String("summary 1"),
				},
				{
					ID:              "67890",
					Title:           "title 2",
					SummaryMarkdown: ptr.String("summary 2"),
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
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := sqldb.NewChangeLogSQL(sqlDB)

					err := changeLogRepo.DeleteChange(testCase.deleteChangeId)
					changeLog, _ := changeLogRepo.GetChangeLog()

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedChangeLogSize, len(changeLog))
				})
		})
	}
}

func TestChangeLogSql_UpdateChange(t *testing.T) {
	testCases := []struct {
		name              string
		tableRows         []changeLogTableRow
		change            entity.Change
		expectedChangeLog []entity.Change
	}{
		{
			name: "update an existing change",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			change: entity.Change{
				ID:              "12345",
				Title:           "title 3",
				SummaryMarkdown: ptr.String("summary 3"),
			},
			expectedChangeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 3",
					SummaryMarkdown: ptr.String("summary 3"),
				},
				{
					ID:              "12346",
					Title:           "title 2",
					SummaryMarkdown: ptr.String("summary 2"),
				},
			},
		}, {
			name: "update a non existing change",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: "summary 1",
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: "summary 2",
				},
			},
			change: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: ptr.String("summary 3"),
			},
			expectedChangeLog: []entity.Change{
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: ptr.String("summary 1"),
				},
				{
					ID:              "12346",
					Title:           "title 2",
					SummaryMarkdown: ptr.String("summary 2"),
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
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := sqldb.NewChangeLogSQL(sqlDB)

					change, err := changeLogRepo.UpdateChange(testCase.change)
					changeLog, _ := changeLogRepo.GetChangeLog()

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.change, change)
					assert.SameElements(t, testCase.expectedChangeLog, changeLog)
				})
		})
	}
}

func insertChangeLogTableRows(t *testing.T, sqlDB *sql.DB, tableRows []changeLogTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertChangeLogRowSQL,
			tableRow.id,
			tableRow.title,
			tableRow.summaryMarkdown,
			tableRow.releasedAt,
		)
		assert.Equal(t, nil, err)
	}
}

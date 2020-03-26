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
	releasedAt      *time.Time
}

func TestChangeLogSql_GetChangeLog(t *testing.T) {
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
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
					summaryMarkdown: summaryMarkdown2,
				}, {
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: summaryMarkdown1,
				},
			},
			expectedChangeLog: []entity.Change{
				{
					ID:              "12346",
					Title:           "title 2",
					SummaryMarkdown: &summaryMarkdown2,
				},
				{
					ID:              "12345",
					Title:           "title 1",
					SummaryMarkdown: &summaryMarkdown1,
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
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := db.NewChangeLogSQL(sqlDB)
					changeLog, _ := changeLogRepo.GetChangeLog()

					mdtest.Equal(t, testCase.expectedChangeLog, changeLog)
				})
		})
	}
}

func TestChangeLogSql_CreateChange(t *testing.T) {
	summaryMarkdown1 := "summary 1"
	summaryMarkdown2 := "summary 2"
	summaryMarkdown3 := "summary 3"
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
					summaryMarkdown: summaryMarkdown1,
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: summaryMarkdown2,
				},
			},
			change: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
			expectedChangeLogSize: 3,
			expectedChange: entity.Change{
				ID:              "23456",
				Title:           "title 3",
				SummaryMarkdown: &summaryMarkdown3,
			},
		},{
			name: "create a change with nil summary",
			tableRows: []changeLogTableRow{
				{
					id:              "12345",
					title:           "title 1",
					summaryMarkdown: summaryMarkdown1,
				},
				{
					id:              "12346",
					title:           "title 2",
					summaryMarkdown: summaryMarkdown2,
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
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertChangeLogTableRows(t, sqlDB, testCase.tableRows)

					changeLogRepo := db.NewChangeLogSQL(sqlDB)

					change, err := changeLogRepo.CreateChange(testCase.change)
					changeLog, _ := changeLogRepo.GetChangeLog()

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectedChangeLogSize, len(changeLog))
					mdtest.Equal(t, testCase.expectedChange, change)
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
		mdtest.Equal(t, nil, err)
	}
}

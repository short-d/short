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

var insertAppRowSQL = fmt.Sprintf(`
INSERT INTO %s (%s, %s, %s) 
VALUES ($1, $2, $3);`,
	table.App.TableName,
	table.App.ColumnID,
	table.App.ColumnName,
	table.App.ColumnCreatedAt,
)

type appTableRow struct {
	id        string
	name      string
	createdAt time.Time
}

func TestAppSQL_FindAppByID(t *testing.T) {
	createdAt := mustParseTime(t, "2017-05-01T08:02:16-07:00")
	testCases := []struct {
		name        string
		tableRows   []appTableRow
		appID       string
		hasErr      bool
		expectedApp entity.App
	}{
		{
			name:      "app not found",
			tableRows: []appTableRow{},
			appID:     "emotic",
			hasErr:    true,
		},
		{
			name: "app found",
			tableRows: []appTableRow{
				{
					id:        "emotic",
					name:      "Feedback Widget",
					createdAt: createdAt,
				},
			},
			appID:  "emotic",
			hasErr: false,
			expectedApp: entity.App{
				ID:        "emotic",
				Name:      "Feedback Widget",
				CreatedAt: createdAt,
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
					insertAppRows(t, sqlDB, testCase.tableRows)

					appRepo := sqldb.NewAppSQL(sqlDB)
					gotApp, err := appRepo.GetAppByID(testCase.appID)
					if testCase.hasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectedApp.ID, gotApp.ID)
					assert.Equal(t, testCase.expectedApp.Name, gotApp.Name)
					assert.Equal(t, testCase.expectedApp.CreatedAt, gotApp.CreatedAt.UTC())
				})
		})
	}
}

func insertAppRows(t *testing.T, sqlDB *sql.DB, tableRows []appTableRow) {
	for _, tableRow := range tableRows {
		_, err := sqlDB.Exec(
			insertAppRowSQL,
			tableRow.id,
			tableRow.name,
			tableRow.createdAt,
		)
		assert.Equal(t, nil, err)
	}
}

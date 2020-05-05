// +build integration all

package sqldb_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/db/dbtest"
	"github.com/short-d/short/backend/app/adapter/sqldb"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
)

type featureToggleTableRow struct {
	toggleID  string
	isEnabled bool
}

func TestFeatureToggleSQL_FindToggleByID(t *testing.T) {
	testCases := []struct {
		name         string
		tableRows    []featureToggleTableRow
		toggleID     string
		expectHasErr bool
		expectToggle entity.Toggle
	}{
		{
			name:         "toggle not found",
			tableRows:    []featureToggleTableRow{},
			toggleID:     "example-feature",
			expectHasErr: true,
		},
		{
			name: "toggle enabled",
			tableRows: []featureToggleTableRow{
				{
					toggleID:  "example-feature",
					isEnabled: true,
				},
			},
			toggleID:     "example-feature",
			expectHasErr: false,
			expectToggle: entity.Toggle{
				ID:        "example-feature",
				IsEnabled: true,
			},
		},
		{
			name: "toggle disabled",
			tableRows: []featureToggleTableRow{
				{
					toggleID:  "example-feature",
					isEnabled: false,
				},
			},
			toggleID:     "example-feature",
			expectHasErr: false,
			expectToggle: entity.Toggle{
				ID:        "example-feature",
				IsEnabled: false,
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
					insertFeatureToggleTableRows(t, sqlDB, testCase.tableRows)

					featureToggleRepo := sqldb.NewFeatureToggleSQL(sqlDB)
					gotToggle, err := featureToggleRepo.FindToggleByID(testCase.toggleID)
					if testCase.expectHasErr {
						assert.NotEqual(t, nil, err)
						return
					}

					assert.Equal(t, nil, err)
					assert.Equal(t, testCase.expectToggle, gotToggle)
				})
		})
	}
}

var insertFeatureToggleRowSQL = fmt.Sprintf(`
INSERT INTO "%s" ("%s", "%s")
VALUES ($1, $2);`,
	table.FeatureToggle.TableName,
	table.FeatureToggle.ColumnToggleID,
	table.FeatureToggle.ColumnIsEnabled,
)

func insertFeatureToggleTableRows(t *testing.T, sqlDB *sql.DB, rows []featureToggleTableRow) {
	for _, row := range rows {
		_, err := sqlDB.Exec(
			insertFeatureToggleRowSQL,
			row.toggleID,
			sqldb.SQLBool(row.isEnabled),
		)
		assert.Equal(t, nil, err)
	}
}

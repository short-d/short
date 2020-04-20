// +build integration all

package db_test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/adapter/db"
	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
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
			mdtest.AccessTestDB(
				dbConnector,
				dbMigrationTool,
				dbMigrationRoot,
				dbConfig,
				func(sqlDB *sql.DB) {
					insertFeatureToggleTableRows(t, sqlDB, testCase.tableRows)

					featureToggleRepo := db.NewFeatureToggleSQL(sqlDB)
					gotToggle, err := featureToggleRepo.FindToggleByID(testCase.toggleID)
					if testCase.expectHasErr {
						mdtest.NotEqual(t, nil, err)
						return
					}

					mdtest.Equal(t, nil, err)
					mdtest.Equal(t, testCase.expectToggle, gotToggle)
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
			db.SQLBool(row.isEnabled),
		)
		mdtest.Equal(t, nil, err)
	}
}

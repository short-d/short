package db

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.FeatureToggle = (*FeatureToggleSQL)(nil)

// FeatureToggleSQL accesses feature toggle information in feature_toggle table through SQL.
type FeatureToggleSQL struct {
	db *sql.DB
}

// FindToggleByID fetches feature toggle from the database given toggle id.
func (f FeatureToggleSQL) FindToggleByID(id string) (entity.Toggle, error) {
	query := fmt.Sprintf(`
SELECT "%s","%s" 
FROM "%s"
WHERE "%s"=$1;`,
		table.FeatureToggle.ColumnToggleID,
		table.FeatureToggle.ColumnIsEnabled,
		table.FeatureToggle.TableName,
		table.FeatureToggle.ColumnToggleID,
	)

	fmt.Println(query)

	toggle := entity.Toggle{}
	err := f.db.QueryRow(query, id).Scan(&toggle.ID, &toggle.IsEnabled)
	return toggle, err
}

// NewFeatureToggleSQL create FeatureToggleSQL repository.
func NewFeatureToggleSQL(db *sql.DB) FeatureToggleSQL {
	return FeatureToggleSQL{db: db}
}

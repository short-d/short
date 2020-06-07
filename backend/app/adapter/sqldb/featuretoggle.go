package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.FeatureToggle = (*FeatureToggleSQL)(nil)

// FeatureToggleSQL accesses feature toggle information in feature_toggle table through SQL.
type FeatureToggleSQL struct {
	db *sql.DB
}

// FindToggleByID fetches feature toggle from the database given toggle id.
func (f FeatureToggleSQL) FindToggleByID(id string) (entity.Toggle, error) {
	query := fmt.Sprintf(`
SELECT "%s","%s", "%s"
FROM "%s"
WHERE "%s"=$1;`,
		table.FeatureToggle.ColumnToggleID,
		table.FeatureToggle.ColumnIsEnabled,
		table.FeatureToggle.ColumnType,
		table.FeatureToggle.TableName,
		table.FeatureToggle.ColumnToggleID,
	)

	toggle := entity.Toggle{}
	err := f.db.QueryRow(query, id).Scan(&toggle.ID, &toggle.IsEnabled, &toggle.Type)
	return toggle, err
}

// NewFeatureToggleSQL create FeatureToggleSQL repository.
func NewFeatureToggleSQL(db *sql.DB) FeatureToggleSQL {
	return FeatureToggleSQL{db: db}
}

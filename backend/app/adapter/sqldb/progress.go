package sqldb

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.Progress = (*IncidentSQL)(nil)

type ProgressSQL struct {
	db *sql.DB
}

func (p ProgressSQL) GetProgress(progressID string) (entity.Progress, error) {
	query := fmt.Sprintf(
		`SELECT "%s", "%s", "%s", "%s" FROM "%s" WHERE "%s"=$1`,
		table.Progress.ColumnIncidentID,
		table.Progress.ColumnStatus,
		table.Progress.ColumnInfo,
		table.Progress.ColumnCreatedAt,
		table.Progress.TableName,
		table.Progress.ColumnIncidentID,
	)
	progress := entity.Progress{}
	err := p.db.QueryRow(query, progress.ID).Scan(&progress.Incident, &progress.Status, &progress.Info, &progress.CreatedAt)
	if err != nil {
		return entity.Progress{}, err
	}
	progress.CreatedAt = utc(progress.createdAt)
	return progress, nil
}

package sqldb

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.Incident = (*IncidentSQL)(nil)

type IncidentSQL struct {
	db *sql.DB
}

func (i IncidentSQL) GetIncident(incidentID string) (entity.Incident, error) {
	query := fmt.Sprintf(
		`SELECT "%s", "%s", "%s" FROM "%s"WHERE "%s"=$1;`,
		table.Incident.ColumnID,
		table.Incident.ColumnTitle,
		table.Incident.ColumnCreatedAt,
		table.Incident.TableName,
		table.Incident.ColumnID,
	)
	incident := entity.Incident{}
	err := i.db.QueryRow(query, incidentID).Scan(&incident.ID, &incident.Title, &incident.CreatedAt)
	if err != nil {
		return entity.Incident{}, err
	}
	incident.CreatedAt = utc(incident.CreatedAt)
	return incident, nil
}

func (i IncidentSQL) GetIncidents(after time.Time) ([]entity.Incident, error) {
	statement := fmt.Sprintf(
		`SELECT "%s", "%s", "%s" FROM "%s" WHERE "%s">$1;`,
		table.Incident.ColumnID,
		table.Incident.ColumnTitle,
		table.Incident.ColumnCreatedAt,
		table.Incident.TableName,
		table.Incident.ColumnCreatedAt,
	)
	var incidents []entity.Incident

	rows, err := i.db.Query(statement, after)

	defer rows.Close()
	if err != nil {
		return incidents, nil
	}

	for rows.Next() {
		incident := entity.Incident{}
		err := rows.Scan(
			&incident.ID,
			&incident.Title,
			&incident.CreatedAt,
		)
		if err != nil {
			return incidents, err
		}
		incident.CreatedAt = utc(incident.CreatedAt)
		incidents = append(incidents, incident)
	}
	return incidents, nil
}

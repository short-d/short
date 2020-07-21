package repository

import (
	"time"

	"github.com/short-d/short/backend/app/entity"
)

// Incident accesses incidents from storage, such as a database.
type Incident interface {
	GetIncident(incidentID string) (entity.Incident, error)
	GetIncidents(after time.Time) ([]entity.Incident, error)
}

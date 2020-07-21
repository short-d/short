package repository

import (
	"errors"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ Incident = (*IncidentFake)(nil)

type IncidentFake struct {
	incidents []entity.Incident
}

// GetIncident finds an Incident in the incident table given an Incident ID
func (i IncidentFake) GetIncident(incidentID string) (entity.Incident, error) {
	for _, incident := range i.incidents {
		if incident.ID == incidentID {
			return incident, nil
		}
	}
	return entity.Incident{}, errors.New("incident not found")
}

// GetIncidents gets all incidents after a given time.
func (i IncidentFake) GetIncidents(after time.Time) ([]entity.Incident, error) {
	var incidents []entity.Incident
	for _, incident := range i.incidents {
		if incident.CreatedAt.After(after) {
			incidents = append(incidents, incident)
		}
	}
	return incidents, nil
}

func NewIncidentFake(incidents []entity.Incident) IncidentFake {
	return IncidentFake{incidents: incidents}
}

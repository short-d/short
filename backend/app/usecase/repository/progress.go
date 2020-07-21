package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// Progress access progress from a storage, such as a database
type Progress interface {
	GetProgress(incidentID string) ([]entity.Progress, error)
}

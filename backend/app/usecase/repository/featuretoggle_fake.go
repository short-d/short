package repository

import (
	"fmt"

	"github.com/short-d/short/backend/app/entity"
)

var _ FeatureToggle = (*FeatureToggleFake)(nil)

// FeatureToggleFake represents in-memory implementation of FeatureToggle repository.
type FeatureToggleFake struct {
	toggles map[string]entity.Toggle
}

// FindToggleByID fetches Toggle with given ID.
func (f FeatureToggleFake) FindToggleByID(id string) (entity.Toggle, error) {
	for _, toggle := range f.toggles {
		if toggle.ID == id {
			return toggle, nil
		}
	}
	return entity.Toggle{}, fmt.Errorf("failed to find toggle with id %s", id)
}

// NewFeatureToggleFake creates fake feature toggle repository.
func NewFeatureToggleFake(toggles map[string]entity.Toggle) FeatureToggleFake {
	return FeatureToggleFake{toggles: toggles}
}

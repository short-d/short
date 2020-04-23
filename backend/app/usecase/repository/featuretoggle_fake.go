package repository

import (
	"fmt"

	"github.com/short-d/short/app/entity"
)

var _ FeatureToggle = (*FeatureToggleFake)(nil)

type FeatureToggleFake struct {
	toggles map[string]entity.Toggle
}

func (f FeatureToggleFake) FindToggleByID(id string) (entity.Toggle, error) {
	for _, toggle := range f.toggles {
		if toggle.ID == id {
			return toggle, nil
		}
	}
	return entity.Toggle{}, fmt.Errorf("failed to find toggle with id %s", id)
}

func NewFeatureToggleFake(toggles map[string]entity.Toggle) FeatureToggleFake {
	return FeatureToggleFake{toggles: toggles}
}

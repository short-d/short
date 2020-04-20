package repository

import "github.com/short-d/short/app/entity"

// FeatureToggle accesses feature toggle from storage, such as database.
type FeatureToggle interface {
	FindToggleByID(id string) (entity.Toggle, error)
}

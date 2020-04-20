package repository

import "github.com/short-d/short/app/entity"

type FeatureToggle interface {
	FindToggleByID(id string) (entity.Toggle, error)
}

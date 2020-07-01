package repository

import (
	"fmt"

	"github.com/short-d/short/backend/app/entity"
)

var _ App = (*AppFake)(nil)

// AppFake represents in memory implementation of App repository
type AppFake struct {
	apps []entity.App
}

// FindAppByID fetches an app with given ID from memory.
func (a AppFake) FindAppByID(id string) (entity.App, error) {
	for _, app := range a.apps {
		if app.ID == appID {
			return app, nil
		}
	}
	return entity.App{}, ErrEntryNotFound(fmt.Sprintf("ID(%s)", appID))
}

// NewAppFake create in-memory implementation of App repository.
func NewAppFake(apps []entity.App) AppFake {
	return AppFake{
		apps: apps,
	}
}

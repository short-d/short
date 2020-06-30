package repository

import "github.com/short-d/short/backend/app/entity"

type App interface {
	FindAppByID(appID string) (entity.App, error)
}

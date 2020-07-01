package repository

import "github.com/short-d/short/backend/app/entity"

// App accesses third party app info from persistent storage, such as database.
type App interface {
	FindAppByID(appID string) (entity.App, error)
}

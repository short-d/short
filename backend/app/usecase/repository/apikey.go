package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// APIKey accesses API keys for third party apps from persistent storage, such as database.
type APIKey interface {
	GetAPIKey(appID string, key string) (entity.APIKey, error)
	CreateAPIKey(input entity.APIKeyInput) (entity.APIKey, error)
}

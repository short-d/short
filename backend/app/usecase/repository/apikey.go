package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

// ApiKey accesses API keys for third party apps from persistent storage, such as database.
type ApiKey interface {
	GetApiKey(appID string, key string) (entity.ApiKey, error)
	CreateApiKey(input entity.ApiKeyInput) (entity.ApiKey, error)
}

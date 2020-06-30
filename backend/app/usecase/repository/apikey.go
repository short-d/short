package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

type ApiKey interface {
	FindApiKey(appID string, key string) (entity.ApiKey, error)
	CreateApiKey(input entity.ApiKeyInput) (entity.ApiKey, error)
}

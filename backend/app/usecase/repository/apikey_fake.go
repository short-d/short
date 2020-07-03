package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ ApiKey = (*ApiKeyFake)(nil)

// ApiKeyFake represents in memory implementation of ApiKey repository
type ApiKeyFake struct {
	apiKeys []entity.ApiKey
}

// GetApiKey fetches an api key for a given app.
func (a ApiKeyFake) GetApiKey(appID string, key string) (entity.ApiKey, error) {
	for _, apiKey := range a.apiKeys {
		if apiKey.AppID == appID && apiKey.Key == key {
			return apiKey, nil
		}
	}
	return entity.ApiKey{}, ErrEntryNotFound(fmt.Sprintf("appID(%s),key(%s)", appID, key))
}

// CreateApiKey creates an api key for a given app.
func (a *ApiKeyFake) CreateApiKey(input entity.ApiKeyInput) (entity.ApiKey, error) {
	if input.AppID == nil {
		return entity.ApiKey{}, errors.New("appID can't be nil")
	}

	if input.Key == nil {
		return entity.ApiKey{}, errors.New("key can't be nil")
	}

	apiKey, err := a.GetApiKey(input.GetAppID(""), input.GetKey(""))
	if err == nil {
		return entity.ApiKey{}, ErrEntryExists(
			fmt.Sprintf("appID(%s),key(%s)", apiKey.AppID, apiKey.Key),
		)
	}
	apiKey = entity.ApiKey{
		AppID:      input.GetAppID(""),
		Key:        input.GetKey(""),
		IsDisabled: input.GetIsDisabled(false),
		CreatedAt:  input.GetCreatedAt(time.Time{}),
	}
	a.apiKeys = append(a.apiKeys, apiKey)
	return apiKey, nil
}

// NewApiKeyFake creates in memory implementation of ApiKey repository.
func NewApiKeyFake(apiKeys []entity.ApiKey) ApiKeyFake {
	return ApiKeyFake{apiKeys: apiKeys}
}

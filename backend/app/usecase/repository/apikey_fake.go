package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ APIKey = (*APIKeyFake)(nil)

// APIKeyFake represents in memory implementation of APIKey repository
type APIKeyFake struct {
	apiKeys []entity.APIKey
}

// GetAPIKey fetches an api key for a given app.
func (a APIKeyFake) GetAPIKey(appID string, key string) (entity.APIKey, error) {
	for _, apiKey := range a.apiKeys {
		if apiKey.AppID == appID && apiKey.Key == key {
			return apiKey, nil
		}
	}
	return entity.APIKey{}, ErrEntryNotFound(fmt.Sprintf("appID(%s),key(%s)", appID, key))
}

// CreateAPIKey creates an api key for a given app.
func (a *APIKeyFake) CreateAPIKey(input entity.APIKeyInput) (entity.APIKey, error) {
	if input.AppID == nil {
		return entity.APIKey{}, errors.New("appID can't be nil")
	}

	if input.Key == nil {
		return entity.APIKey{}, errors.New("key can't be nil")
	}

	apiKey, err := a.GetAPIKey(input.GetAppID(""), input.GetKey(""))
	if err == nil {
		return entity.APIKey{}, ErrEntryExists(
			fmt.Sprintf("appID(%s),key(%s)", apiKey.AppID, apiKey.Key),
		)
	}
	apiKey = entity.APIKey{
		AppID:      input.GetAppID(""),
		Key:        input.GetKey(""),
		IsDisabled: input.GetIsDisabled(false),
		CreatedAt:  input.GetCreatedAt(time.Time{}),
	}
	a.apiKeys = append(a.apiKeys, apiKey)
	return apiKey, nil
}

// NewAPIKeyFake creates in memory implementation of APIKey repository.
func NewAPIKeyFake(apiKeys []entity.APIKey) APIKeyFake {
	return APIKeyFake{apiKeys: apiKeys}
}

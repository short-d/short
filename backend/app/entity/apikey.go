package entity

import "time"

// ApiKey allows external systems to access contain API enabled for the given
// third party app.
type ApiKey struct {
	AppID      string
	Key        string
	IsDisabled bool
	CreatedAt  time.Time
}

// ApiKeyInput represents ApiKey with all fields optional and default values
// for certain fields.
type ApiKeyInput struct {
	AppID      *string
	Key        *string
	IsDisabled *bool
	CreatedAt  *time.Time
}

// GetAppID fetches AppID for ApiKeyInput with default value.
func (a ApiKeyInput) GetAppID(defaultVal string) string {
	if a.AppID == nil {
		return defaultVal
	}
	return *a.AppID
}

// GetKey fetches Key for ApiKeyInput with default value.
func (a ApiKeyInput) GetKey(defaultVal string) string {
	if a.Key == nil {
		return defaultVal
	}
	return *a.Key
}

// GetIsDisabled fetches isDisabled for ApiKeyInput with default value.
func (a ApiKeyInput) GetIsDisabled(defaultVal bool) bool {
	if a.IsDisabled == nil {
		return defaultVal
	}
	return *a.IsDisabled
}

// GetCreatedAt fetches createdAt for ApiKeyInput with default value.
func (a ApiKeyInput) GetCreatedAt(defaultVal time.Time) time.Time {
	if a.CreatedAt == nil {
		return defaultVal
	}
	return *a.CreatedAt
}

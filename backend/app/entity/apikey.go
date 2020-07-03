package entity

import "time"

// APIKey allows external systems to access contain API enabled for the given
// third party app.
type APIKey struct {
	AppID      string
	Key        string
	IsDisabled bool
	CreatedAt  time.Time
}

// APIKeyInput represents APIKey with all fields optional and default values
// for certain fields.
type APIKeyInput struct {
	AppID      *string
	Key        *string
	IsDisabled *bool
	CreatedAt  *time.Time
}

// GetAppID fetches AppID for APIKeyInput with default value.
func (a APIKeyInput) GetAppID(defaultVal string) string {
	if a.AppID == nil {
		return defaultVal
	}
	return *a.AppID
}

// GetKey fetches Key for APIKeyInput with default value.
func (a APIKeyInput) GetKey(defaultVal string) string {
	if a.Key == nil {
		return defaultVal
	}
	return *a.Key
}

// GetIsDisabled fetches isDisabled for APIKeyInput with default value.
func (a APIKeyInput) GetIsDisabled(defaultVal bool) bool {
	if a.IsDisabled == nil {
		return defaultVal
	}
	return *a.IsDisabled
}

// GetCreatedAt fetches createdAt for APIKeyInput with default value.
func (a APIKeyInput) GetCreatedAt(defaultVal time.Time) time.Time {
	if a.CreatedAt == nil {
		return defaultVal
	}
	return *a.CreatedAt
}

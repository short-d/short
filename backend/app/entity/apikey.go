package entity

import "time"

type ApiKey struct {
	AppID string
	Key string
	IsDisabled bool
	CreatedAt time.Time
}

type ApiKeyInput struct {
	AppID *string
	Key *string
	IsDisabled *bool
	CreatedAt *time.Time
}

func (a ApiKeyInput) GetAppID(defaultVal string) string {
	if a.AppID == nil {
		return defaultVal
	}
	return *a.AppID
}

func (a ApiKeyInput) GetKey(defaultVal string) string {
	if a.Key == nil {
		return defaultVal
	}
	return *a.Key
}

func (a ApiKeyInput) GetIsDisabled(defaultVal bool) bool {
	if a.IsDisabled == nil {
		return defaultVal
	}
	return *a.IsDisabled
}

func (a ApiKeyInput) GetCreatedAt(defaultVal time.Time) time.Time {
	if a.CreatedAt == nil {
		return defaultVal
	}
	return *a.CreatedAt
}
package payload

import (
	"errors"

	"github.com/short-d/app/fw/crypto"
)

// APIKey represents the payload of an API key.
type APIKey struct {
	AppID string
	Key   string
}

// NewAPIKey parses APIKey from token payload.
func NewAPIKey(payload crypto.TokenPayload) (APIKey, error) {
	appID, ok := payload["app_id"]
	if !ok {
		return APIKey{}, errors.New("missing app id")
	}
	key, ok := payload["key"]
	if !ok {
		return APIKey{}, errors.New("missing key")
	}
	return APIKey{
		AppID: appID.(string),
		Key:   key.(string),
	}, nil
}

// NewTokenPayload converts API key into token payload.
func (a *APIKey) NewTokenPayload() crypto.TokenPayload {
	payload := crypto.TokenPayload{}
	payload["app_id"] = a.AppID
	payload["key"] = a.Key
	return payload
}

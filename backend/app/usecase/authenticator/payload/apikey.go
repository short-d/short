package payload

import (
	"errors"

	"github.com/short-d/app/fw/crypto"
)

var (
	// ErrMissingAppID implies that the payload of API key does NOT contain app ID.
	ErrMissingAppID = errors.New("missing app id")
	// ErrMissingKey implies that the payload of API key does NOT contain secret key.
	ErrMissingKey = errors.New("missing key")
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
		return APIKey{}, ErrMissingAppID
	}
	key, ok := payload["key"]
	if !ok {
		return APIKey{}, ErrMissingKey
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

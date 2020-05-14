package authenticator

import (
	"errors"
	"time"

	"github.com/short-d/app/fw/crypto"
)

// Payload represents the metadata encoded in the authentication token.
type Payload struct {
	id       string
	issuedAt time.Time
}

// TokenPayload retrieves key-value pairs representation of the payload.
func (p Payload) TokenPayload() crypto.TokenPayload {
	return map[string]interface{}{
		"id":        p.id,
		"issued_at": p.issuedAt,
	}
}

func newPayload(id string, issuedAt time.Time) Payload {
	return Payload{
		id:       id,
		issuedAt: issuedAt,
	}
}

func fromTokenPayload(tokenPayload crypto.TokenPayload) (Payload, error) {
	payload := Payload{}
	var ok bool

	id := tokenPayload["id"]
	if payload.id, ok = id.(string); !ok {
		return payload, errors.New("expect payload to contain id")
	}

	issuedAtJSON := tokenPayload["issued_at"]
	var issuedAtStr string
	if issuedAtStr, ok = issuedAtJSON.(string); !ok {
		return payload, errors.New("expect payload to contain issued_at")
	}

	issuedAt, err := time.Parse(time.RFC3339, issuedAtStr)
	if err != nil {
		return payload, err
	}
	payload.issuedAt = issuedAt

	return payload, nil
}

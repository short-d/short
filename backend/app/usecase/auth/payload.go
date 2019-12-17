package auth

import (
	"errors"
	"time"

	"github.com/byliuyang/app/fw"
)

// Payload represents the metadata encoded in the authentication token.
type Payload struct {
	email    string
	issuedAt time.Time
}

// TokenPayload retrieves key-value pairs representation of the payload.
func (p Payload) TokenPayload() fw.TokenPayload {
	return map[string]interface{}{
		"email":     p.email,
		"issued_at": p.issuedAt,
	}
}

func newPayload(email string, issuedAt time.Time) Payload {
	return Payload{
		email:    email,
		issuedAt: issuedAt,
	}
}

func fromTokenPayload(tokenPayload fw.TokenPayload) (Payload, error) {
	payload := Payload{}
	var ok bool

	email := tokenPayload["email"]
	if payload.email, ok = email.(string); !ok {
		return payload, errors.New("expect payload to contain email")
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

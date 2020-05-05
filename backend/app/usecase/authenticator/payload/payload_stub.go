package payload

import (
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/short/backend/app/entity"
)

var _ Payload = (*Stub)(nil)

// Stub represents a payload with preset data.
type Stub struct {
	TokenPayload crypto.TokenPayload
	User         entity.User
}

// GetTokenPayload generates token payload based on preset data.
func (s Stub) GetTokenPayload() crypto.TokenPayload {
	return s.TokenPayload
}

// GetUser retrieves the user based on preset data.
func (s Stub) GetUser() entity.User {
	return s.User
}

package payloadtest

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth/payload"
)

var _ payload.Payload = (*Stub)(nil)

// Stub represents a payload with preset data.
type Stub struct {
	TokenPayload fw.TokenPayload
	User         entity.User
}

// GetTokenPayload generates token payload based on preset data.
func (s Stub) GetTokenPayload() fw.TokenPayload {
	return s.TokenPayload
}

// GetUser retrieves the user based on preset data.
func (s Stub) GetUser() entity.User {
	return s.User
}

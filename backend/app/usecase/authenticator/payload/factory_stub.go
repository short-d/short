package payload

import (
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/short/app/entity"
)

var _ Factory = (*FactoryStub)(nil)

// FactoryStub creates payloads based on preset data.
type FactoryStub struct {
	Payload  Payload
	TokenErr error
	UserErr  error
}

// FromTokenPayload creates payload based on preset payload and error.
func (f FactoryStub) FromTokenPayload(tokenPayload crypto.TokenPayload) (Payload, error) {
	return f.Payload, f.TokenErr
}

// FromUser creates payload based on preset payload and error.
func (f FactoryStub) FromUser(user entity.User) (Payload, error) {
	return f.Payload, f.UserErr
}

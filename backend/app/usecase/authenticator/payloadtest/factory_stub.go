package payloadtest

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/authenticator/payload"
)

var _ payload.Factory = (*FactoryStub)(nil)

// FactoryStub creates payloads based on preset data.
type FactoryStub struct {
	Payload  payload.Payload
	TokenErr error
	UserErr  error
}

// FromTokenPayload creates payload based on preset payload and error.
func (f FactoryStub) FromTokenPayload(tokenPayload fw.TokenPayload) (payload.Payload, error) {
	return f.Payload, f.TokenErr
}

// FromUser creates payload based on preset payload and error.
func (f FactoryStub) FromUser(user entity.User) (payload.Payload, error) {
	return f.Payload, f.UserErr
}

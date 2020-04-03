package payloadtest

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth/payload"
)

var _ Factory = (*FactoryStub)(nil)

type FactoryStub struct {
	Payload  Payload
	TokenErr error
	UserErr  error
}

func (f FactoryStub) FromTokenPayload(tokenPayload fw.TokenPayload) (payload.Payload, error) {
	return f.Payload, f.TokenErr
}

func (f FactoryStub) FromUser(user entity.User) (payload.Payload, error) {
	return f.Payload, f.UserErr
}

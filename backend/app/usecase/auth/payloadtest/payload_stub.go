package payloadtest

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/auth/payload"
)

var _ payload.Payload = (*Stub)(nil)

type Stub struct {
	TokenPayload fw.TokenPayload
	User         entity.User
}

func (s Stub) GetTokenPayload() fw.TokenPayload {
	return s.TokenPayload
}

func (s Stub) GetUser() entity.User {
	return s.User
}

package payloadtest

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

var _ Payload = (*Stub)(nil)

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

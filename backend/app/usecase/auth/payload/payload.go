package payload

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

type Payload interface {
	GetTokenPayload() fw.TokenPayload
	GetUser() entity.User
}

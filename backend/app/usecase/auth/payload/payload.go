package payload

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

// Payload represents the metadata encoded in a message.
type Payload interface {
	GetTokenPayload() fw.TokenPayload
	GetUser() entity.User
}

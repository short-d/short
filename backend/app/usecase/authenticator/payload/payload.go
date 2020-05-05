package payload

import (
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/short/app/entity"
)

// Payload represents a message with encoded metadata.
type Payload interface {
	GetTokenPayload() crypto.TokenPayload
	GetUser() entity.User
}

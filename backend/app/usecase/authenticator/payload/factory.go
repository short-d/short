package payload

import (
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/short/app/entity"
)

// Factory creates payload based on metadata provided.
type Factory interface {
	FromTokenPayload(tokenPayload crypto.TokenPayload) (Payload, error)
	FromUser(user entity.User) (Payload, error)
}

package payload

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/entity"
)

// Factory creates payload based on metadata provided.
type Factory interface {
	FromTokenPayload(tokenPayload fw.TokenPayload) (Payload, error)
	FromUser(user entity.User) (Payload, error)
}

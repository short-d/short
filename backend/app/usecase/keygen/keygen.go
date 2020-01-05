package keygen

import "github.com/short-d/kgs/app/entity"

// KeyGenerator generates unique keys.
type KeyGenerator interface {
	NewKey() (entity.Key, error)
}

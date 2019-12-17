package keygen

import "github.com/byliuyang/kgs/app/entity"

// KeyGenerator generates unique keys.
type KeyGenerator interface {
	NewKey() (entity.Key, error)
}

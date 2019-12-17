package keygen

import "github.com/byliuyang/kgs/app/entity"

// KeyGenerator generates unique key.
type KeyGenerator interface {
	NewKey() (entity.Key, error)
}

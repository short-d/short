package keygen

import "github.com/byliuyang/kgs/app/entity"

type KeyGenerator interface {
	NewKey() (entity.Key, error)
}

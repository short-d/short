package keygen

import (
	"errors"

	"github.com/byliuyang/kgs/app/entity"
)

var _ KeyGenerator = (*Fake)(nil)

// Fake represents in-memory implementation of key generator
type Fake struct {
	keys       []string
	currKeyIdx int
}

// NewKey generates a new unique key
func (k *Fake) NewKey() (entity.Key, error) {
	if k.currKeyIdx >= len(k.keys) {
		return "", errors.New("no available key")
	}

	key := k.keys[k.currKeyIdx]
	k.currKeyIdx++
	return entity.Key(key), nil
}

// NewFake initialize Fake key generator
func NewFake(keys []string) Fake {
	return Fake{
		keys: keys,
	}
}

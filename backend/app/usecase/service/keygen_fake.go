package service

import (
	"errors"
	"math"

	"github.com/byliuyang/kgs/app/entity"
)

var _ KeyGen = (*KeyGenFake)(nil)

// KeyGenFake represents an in memory key generator
type KeyGenFake struct {
	availableKeys []entity.Key
}

// FetchKeys returns keys from the buffer
func (k *KeyGenFake) FetchKeys(maxCount int) ([]entity.Key, error) {
	if len(k.availableKeys) < 1 {
		return nil, errors.New("no available key")
	}
	keyCount := int(math.Min(float64(len(k.availableKeys)), float64(maxCount)))
	allocatedKeys := k.availableKeys[:keyCount]
	k.availableKeys = k.availableKeys[keyCount:]
	return allocatedKeys, nil
}

// NewKeyGenFake creates fake key generator
func NewKeyGenFake(availableKeys []entity.Key) KeyGenFake {
	return KeyGenFake{
		availableKeys: availableKeys,
	}
}

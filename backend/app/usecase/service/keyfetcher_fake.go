package service

import (
	"errors"
	"math"

	"github.com/byliuyang/kgs/app/entity"
)

var _ KeyFetcher = (*KeyFetcherFake)(nil)

// KeyFetcherFake represents an in memory key generator
type KeyFetcherFake struct {
	availableKeys []entity.Key
}

// FetchKeys returns keys from the buffer
func (k *KeyFetcherFake) FetchKeys(maxCount int) ([]entity.Key, error) {
	if len(k.availableKeys) < 1 {
		return nil, errors.New("no available key")
	}
	keyCount := int(math.Min(float64(len(k.availableKeys)), float64(maxCount)))
	allocatedKeys := k.availableKeys[:keyCount]
	k.availableKeys = k.availableKeys[keyCount:]
	return allocatedKeys, nil
}

// NewKeyFetcherFake creates fake key generator
func NewKeyFetcherFake(availableKeys []entity.Key) KeyFetcherFake {
	return KeyFetcherFake{
		availableKeys: availableKeys,
	}
}

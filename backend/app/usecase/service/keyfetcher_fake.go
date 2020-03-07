package service

import (
	"errors"
	"math"
)

var _ KeyFetcher = (*KeyFetcherFake)(nil)

// KeyFetcherFake represents an in memory key generator
type KeyFetcherFake struct {
	availableKeys []Key
}

// FetchKeys returns keys from the buffer
func (k *KeyFetcherFake) FetchKeys(maxCount int) ([]Key, error) {
	if len(k.availableKeys) < 1 {
		return nil, errors.New("no available key")
	}
	keyCount := int(math.Min(float64(len(k.availableKeys)), float64(maxCount)))
	allocatedKeys := k.availableKeys[:keyCount]
	k.availableKeys = k.availableKeys[keyCount:]
	return allocatedKeys, nil
}

// NewKeyFetcherFake creates fake key generator
func NewKeyFetcherFake(availableKeys []Key) KeyFetcherFake {
	return KeyFetcherFake{
		availableKeys: availableKeys,
	}
}

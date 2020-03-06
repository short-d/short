package keygen

import (
	"errors"

	"github.com/short-d/short/app/usecase/service"
)

type bufferEntry struct {
	key service.Key
	err error
}

// KeyGenerator fetches unique keys in batch from key generation service
// and buffer them in memory for fast response.
type KeyGenerator struct {
	bufferSize int
	buffer     chan bufferEntry
	keyFetcher service.KeyFetcher
}

// NewKey produces a unique key
func (r KeyGenerator) NewKey() (service.Key, error) {
	if len(r.buffer) == 0 {
		go func() {
			r.fetchKeys()
		}()
	}

	entry := <-r.buffer
	return entry.key, entry.err
}

func (r KeyGenerator) fetchKeys() {
	keys, err := r.keyFetcher.FetchKeys(r.bufferSize)
	if err != nil {
		r.buffer <- bufferEntry{
			key: "",
			err: err,
		}
		return
	}

	for _, key := range keys {
		r.buffer <- bufferEntry{
			key: key,
			err: nil,
		}
	}
}

// NewKeyGenerator creates KeyGenerator
func NewKeyGenerator(bufferSize int, keyFetcher service.KeyFetcher) (KeyGenerator, error) {
	if bufferSize < 1 {
		return KeyGenerator{}, errors.New("buffer size can't be less than 1")
	}
	return KeyGenerator{
		bufferSize: bufferSize,
		buffer:     make(chan bufferEntry, bufferSize),
		keyFetcher: keyFetcher,
	}, nil
}

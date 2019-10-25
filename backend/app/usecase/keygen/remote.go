package keygen

import (
	"errors"
	"short/app/usecase/service"

	"github.com/byliuyang/kgs/app/entity"
)

var _ KeyGenerator = (*Remote)(nil)

type bufferEntry struct {
	key entity.Key
	err error
}

// Remote represents a KeyGenerator which fetch unique keys from remote service
// and buffer them in memory for fast response.
type Remote struct {
	bufferSize    int
	buffer        chan bufferEntry
	keyGenService service.KeyGen
}

// NewKey produces a key
func (r Remote) NewKey() (entity.Key, error) {
	if len(r.buffer) == 0 {
		go func() {
			r.fetchKeys()
		}()
	}

	entry := <-r.buffer
	return entry.key, entry.err
}

func (r Remote) fetchKeys() {
	keys, err := r.keyGenService.FetchKeys(r.bufferSize)
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

// NewRemote creates Remote keygen generator
func NewRemote(bufferSize int, keyGenService service.KeyGen) (Remote, error) {
	if bufferSize < 1 {
		return Remote{}, errors.New("buffer size can't be less than 1")
	}
	return Remote{
		bufferSize:    bufferSize,
		buffer:        make(chan bufferEntry, bufferSize),
		keyGenService: keyGenService,
	}, nil
}

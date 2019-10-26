package provider

import (
	"short/app/usecase/keygen"
	"short/app/usecase/service"
)

// KeyGenBufferSize specifies the size of the local cache for fetched keys
type KeyGenBufferSize int

// NewRemote creates Remote with KeyGenBufferSize to uniquely identify
// bufferSize
func NewRemote(
	bufferSize KeyGenBufferSize,
	keyFetcher service.KeyFetcher,
) (keygen.Remote, error) {
	return keygen.NewRemote(int(bufferSize), keyFetcher)
}

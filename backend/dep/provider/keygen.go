package provider

import (
	"github.com/short-d/short/app/usecase/external"
	"github.com/short-d/short/app/usecase/keygen"
)

// KeyGenBufferSize specifies the size of the local cache for fetched keys
type KeyGenBufferSize int

// NewKeyGenerator creates KeyGenerator with KeyGenBufferSize to uniquely identify
// bufferSize
func NewKeyGenerator(
	bufferSize KeyGenBufferSize,
	keyFetcher external.KeyFetcher,
) (keygen.KeyGenerator, error) {
	return keygen.NewKeyGenerator(int(bufferSize), keyFetcher)
}

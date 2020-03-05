package provider

import (
	"github.com/short-d/short/app/usecase/keygen"
	"github.com/short-d/short/app/usecase/service"
)

// KeyGenBufferSize specifies the size of the local cache for fetched keys
type KeyGenBufferSize int

// NewKeyGenerator creates remote NewKeyGenerator with KeyGenBufferSize to uniquely identify
// bufferSize
func NewKeyGenerator(
	bufferSize KeyGenBufferSize,
	keyFetcher service.KeyFetcher,
) (keygen.KeyGenerator, error) {
	return keygen.NewKeyGenerator(int(bufferSize), keyFetcher)
}

package provider

import (
	"github.com/short-d/short/backend/app/usecase/keygen"
)

// KeyGenBufferSize specifies the size of the local cache for fetched keys
type KeyGenBufferSize int

// NewKeyGenerator creates KeyGenerator with KeyGenBufferSize to uniquely identify
// bufferSize
func NewKeyGenerator(
	bufferSize KeyGenBufferSize,
	keyFetcher keygen.KeyFetcher,
) (keygen.KeyGenerator, error) {
	return keygen.NewKeyGenerator(int(bufferSize), keyFetcher)
}

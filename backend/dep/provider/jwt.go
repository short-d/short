package provider

import (
	"github.com/byliuyang/app/modern/mdcrypto"

	"github.com/byliuyang/app/fw"
)

// JwtSecret represents the secret used to encode and decode JWT token.
type JwtSecret string

// NewJwtGo creates Crypto Tokenizer with JwtSecret to uniquely identify secret during dependency injection.
func NewJwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

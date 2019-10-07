package provider

import (
	"github.com/byliuyang/app/modern/mdcrypto"

	"github.com/byliuyang/app/fw"
)

// JwtSecret secret to encode and decode JWT.
type JwtSecret string

// JwtGo initializes Crypto Tokenizer for encode and decode tokens using the JWT secret.
func JwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

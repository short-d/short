package provider

import (
	"short/app/usecase/auth"
	"time"

	"github.com/byliuyang/app/fw"
)

// TokenValidDuration duration of a valid token.
type TokenValidDuration time.Duration

// Authenticator initializes Authenticator to generete and get information of tokens.
func Authenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer, duration TokenValidDuration) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, time.Duration(duration))
}

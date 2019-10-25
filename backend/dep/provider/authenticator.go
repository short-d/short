package provider

import (
	"short/app/usecase/auth"
	"time"

	"github.com/byliuyang/app/fw"
)

// TokenValidDuration represents the duration of a valid token.
type TokenValidDuration time.Duration

// Authenticator creates Authenticator with TokenValidDuration to uniquely identify duration during dependency injection.
func NewAuthenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer, duration TokenValidDuration) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, time.Duration(duration))
}

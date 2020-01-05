package provider

import (
	"time"

	"github.com/short-d/short/app/usecase/auth"

	"github.com/short-d/app/fw"
)

// TokenValidDuration represents the duration of a valid token.
type TokenValidDuration time.Duration

// NewAuthenticator creates Authenticator with TokenValidDuration to uniquely identify duration during dependency injection.
func NewAuthenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer, duration TokenValidDuration) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, time.Duration(duration))
}

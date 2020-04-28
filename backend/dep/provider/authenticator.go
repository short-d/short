package provider

import (
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/authenticator"
)

// TokenValidDuration represents the duration of a valid token.
type TokenValidDuration time.Duration

// NewAuthenticator creates Authenticator with TokenValidDuration to uniquely identify duration during dependency injection.
func NewAuthenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer, duration TokenValidDuration) authenticator.Authenticator {
	return authenticator.NewAuthenticator(tokenizer, timer, time.Duration(duration))
}

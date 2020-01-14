package auth

import (
	"time"

	"github.com/short-d/app/mdtest"
)

// NewAuthenticatorFake creates fake authenticator for easy testing.
func NewAuthenticatorFake(current time.Time, validPeriod time.Duration) Authenticator {
	tokenizer := mdtest.NewCryptoTokenizerFake()
	timer := mdtest.NewTimerFake(current)
	return NewAuthenticator(tokenizer, timer, validPeriod)
}

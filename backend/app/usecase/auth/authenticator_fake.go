package auth

import (
	"time"

	"github.com/byliuyang/app/mdtest"
)

// NewAuthenticatorFake creates authenticator.
func NewAuthenticatorFake(current time.Time, validPeriod time.Duration) Authenticator {
	tokenizer := mdtest.NewCryptoTokenizerFake()
	timer := mdtest.NewTimerFake(current)
	return NewAuthenticator(tokenizer, timer, validPeriod)
}

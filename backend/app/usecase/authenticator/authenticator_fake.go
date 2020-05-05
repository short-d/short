package authenticator

import (
	"time"

	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
)

// NewAuthenticatorFake creates fake authenticator for easy testing.
func NewAuthenticatorFake(current time.Time, validPeriod time.Duration) Authenticator {
	tokenizer := crypto.NewTokenizerFake()
	tm := timer.NewStub(current)
	return NewAuthenticator(tokenizer, tm, validPeriod)
}

package auth

import (
	"time"

	"short/mdtest"
)

func NewAuthenticatorFake(current time.Time, validPeriod time.Duration) Authenticator {
	tokenizer := mdtest.FakeCryptoTokenizer
	timer := mdtest.NewFakeTimer(current)
	return NewAuthenticator(tokenizer, timer, validPeriod)
}

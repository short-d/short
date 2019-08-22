package auth

import (
	"short/mdtest"
	"time"
)

func NewAuthenticatorFake(current time.Time, validPeriod time.Duration) Authenticator {
	tokenizer := mdtest.FakeCryptoTokenizer
	timer := mdtest.NewFakeTimer(current)
	return NewAuthenticator(tokenizer, timer, validPeriod)
}

package provider

import (
	"short/app/usecase/auth"
	"time"

	"github.com/byliuyang/app/fw"
)

type TokenValidDuration time.Duration

func Authenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer, duration TokenValidDuration) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, time.Duration(duration))
}

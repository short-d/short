package inject

import (
	"short/app/usecase/auth"
	"time"

	"short/fw"
)

const oneDay = 24 * time.Hour
const oneWeek = 7 * oneDay

func Authenticator(tokenizer fw.CryptoTokenizer, timer fw.Timer) auth.Authenticator {
	return auth.NewAuthenticator(tokenizer, timer, oneWeek)
}

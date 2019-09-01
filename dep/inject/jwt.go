package inject

import (
	"short/modern/mdcrypto"

	"short/fw"
)

type JwtSecret string

func JwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

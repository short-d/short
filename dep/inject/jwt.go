package inject

import (
	"short/fw"
	"short/modern/mdcrypto"
)

type JwtSecret string

func JwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

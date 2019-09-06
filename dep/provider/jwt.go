package provider

import (
	"github.com/byliuyang/app/modern/mdcrypto"

	"github.com/byliuyang/app/fw"
)

type JwtSecret string

func JwtGo(secret JwtSecret) fw.CryptoTokenizer {
	return mdcrypto.NewJwtGo(string(secret))
}

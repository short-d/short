package fw

type TokenPayload = map[string]interface{}

type CryptoTokenizer interface {
	Encode(payload TokenPayload) (string, error)
	Decode(tokenStr string) (TokenPayload, error)
}

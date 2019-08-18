package mdtest

import (
	"encoding/json"
	"short/fw"
)

type Tokenizer struct {
}

func (t Tokenizer) Encode(payload fw.TokenPayload) (string, error) {
	buf, err := json.Marshal(payload)
	return string(buf), err
}

func (t Tokenizer) Decode(tokenStr string) (fw.TokenPayload, error) {
	payload := map[string]interface{}{}
	err := json.Unmarshal([]byte(tokenStr), payload)
	return payload, err
}

var FakeCryptoTokenizer fw.CryptoTokenizer = Tokenizer{}

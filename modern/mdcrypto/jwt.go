package mdcrypto

import (
	"errors"
	"short/fw"

	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JwtGo struct {
	algorithm jwt.SigningMethod
	secret    []byte
}

func (j JwtGo) Encode(payload fw.TokenPayload) (string, error) {
	token := jwt.NewWithClaims(j.algorithm, jwt.MapClaims(payload))
	return token.SignedString(j.secret)
}

func (j JwtGo) Decode(tokenStr string) (fw.TokenPayload, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	expAlgorithm := j.algorithm.Alg()
	gotAlgorithm := token.Method.Alg()
	if gotAlgorithm != expAlgorithm {
		return nil, errors.New(fmt.Sprintf("unexpected signing method: exp=%v, got=%v", expAlgorithm, gotAlgorithm))
	}

	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = token.Claims.(jwt.MapClaims); !ok {
		return nil, errors.New("token payload is not map")
	}

	return claims, nil
}

func NewJwtGo(secret string) fw.CryptoTokenizer {
	return JwtGo{
		algorithm: jwt.SigningMethodHS256,
		secret:    []byte(secret),
	}
}

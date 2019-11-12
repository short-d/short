package auth

import (
	"errors"
	"short/app/entity"
	"time"

	"github.com/byliuyang/app/fw"
)

type Authenticator struct {
	tokenizer          fw.CryptoTokenizer
	timer              fw.Timer
	tokenValidDuration time.Duration
}

func (a Authenticator) isTokenValid(payload Payload, validDuring time.Duration) bool {
	now := a.timer.Now()
	tokenExpireAt := payload.issuedAt.Add(validDuring)
	return !tokenExpireAt.Before(now)
}

func (a Authenticator) getPayload(token string) (Payload, error) {
	tokenPayload, err := a.tokenizer.Decode(token)
	if err != nil {
		return Payload{}, err
	}

	payload, err := fromTokenPayload(tokenPayload)
	if err != nil {
		return Payload{}, err
	}
	return payload, nil
}

func (a Authenticator) IsSignedIn(token string) bool {
	payload, err := a.getPayload(token)
	if err != nil {
		return false
	}

	if !a.isTokenValid(payload, a.tokenValidDuration) {
		return false
	}

	return true
}

func (a Authenticator) GetUser(token string) (entity.User, error) {
	payload, err := a.getPayload(token)
	if err != nil {
		return entity.User{}, err
	}

	if !a.isTokenValid(payload, a.tokenValidDuration) {
		return entity.User{}, errors.New("token expired")
	}

	if len(payload.email) < 1 {
		return entity.User{}, errors.New("email can't be empty")
	}
	return entity.User{
		Email: payload.email,
	}, nil
}

func (a Authenticator) GenerateToken(user entity.User) (string, error) {
	issuedAt := a.timer.Now()
	payload := newPayload(user.Email, issuedAt)
	tokenPayload := payload.TokenPayload()
	return a.tokenizer.Encode(tokenPayload)
}

func NewAuthenticator(
	tokenizer fw.CryptoTokenizer,
	timer fw.Timer,
	tokenValidDuration time.Duration,
) Authenticator {
	return Authenticator{
		tokenizer:          tokenizer,
		timer:              timer,
		tokenValidDuration: tokenValidDuration,
	}
}

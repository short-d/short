// +build !integration all

package authenticator

import (
	"testing"
	"time"

	"github.com/short-d/app/fw"
	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
)

func TestAuthenticator_GenerateToken(t *testing.T) {
	tokenizer := mdtest.NewCryptoTokenizerFake()
	expIssuedAt := time.Now()
	timer := mdtest.NewTimerFake(expIssuedAt)
	authenticator := NewAuthenticator(tokenizer, timer, 2*time.Millisecond)

	expUser := entity.User{
		Email: "test@s.time4hacks.com",
	}
	token, err := authenticator.GenerateToken(expUser)
	mdtest.Equal(t, nil, err)

	tokenPayload, err := tokenizer.Decode(token)
	mdtest.Equal(t, nil, err)

	mdtest.Equal(t, expUser.Email, tokenPayload["email"])

	expIssuedAtStr := expIssuedAt.Format(time.RFC3339Nano)
	mdtest.Equal(t, expIssuedAtStr, tokenPayload["issued_at"])
}

func TestAuthenticator_IsSignedIn(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name               string
		expIssuedAt        time.Time
		tokenValidDuration time.Duration
		currentTime        time.Time
		tokenPayload       fw.TokenPayload
		expIsSignIn        bool
	}{
		{
			name:               "Token payload empty",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload:       map[string]interface{}{},
			expIsSignIn:        false,
		},
		{
			name:               "Token payload without email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"issued_at": now.Format(time.RFC3339Nano),
			},
			expIsSignIn: false,
		},
		{
			name:               "Token payload has empty email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email":     "",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			expIsSignIn: false,
		},
		{
			name:               "Token payload without issue_at",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email": "test@s.time4hacks.com",
			},
			expIsSignIn: false,
		},
		{
			name:               "Token expired",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(2 * time.Hour),
			tokenPayload: map[string]interface{}{
				"email":     "test@s.time4hacks.com",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			expIsSignIn: false,
		},
		{
			name:               "Token valid",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email":     "test@s.time4hacks.com",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			expIsSignIn: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tokenizer := mdtest.NewCryptoTokenizerFake()
			timer := mdtest.NewTimerFake(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, timer, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			mdtest.Equal(t, nil, err)
			gotIsSignIn := authenticator.IsSignedIn(token)
			mdtest.Equal(t, testCase.expIsSignIn, gotIsSignIn)
		})
	}
}

func TestAuthenticator_GetUser(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name               string
		expIssuedAt        time.Time
		tokenValidDuration time.Duration
		currentTime        time.Time
		tokenPayload       fw.TokenPayload
		hasErr             bool
		expUser            entity.User
	}{
		{
			name:               "Token payload empty",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload:       map[string]interface{}{},
			hasErr:             true,
			expUser:            entity.User{},
		},
		{
			name:               "Token payload without email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"issued_at": now.Format(time.RFC3339Nano),
			},
			hasErr:  true,
			expUser: entity.User{},
		},
		{
			name:               "Token payload without issue_at",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email": "test@s.time4hacks.com",
			},
			hasErr:  true,
			expUser: entity.User{},
		},
		{
			name:               "Token expired",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(2 * time.Hour),
			tokenPayload: map[string]interface{}{
				"email":     "test@s.time4hacks.com",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			hasErr:  true,
			expUser: entity.User{},
		},
		{
			name:               "Valid token with empty email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email":     "",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			hasErr:  true,
			expUser: entity.User{},
		},
		{
			name:               "Token valid with correct email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email":     "test@s.time4hacks.com",
				"issued_at": now.Format(time.RFC3339Nano),
			},
			hasErr: false,
			expUser: entity.User{
				Email: "test@s.time4hacks.com",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tokenizer := mdtest.NewCryptoTokenizerFake()
			timer := mdtest.NewTimerFake(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, timer, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			mdtest.Equal(t, nil, err)
			gotUser, err := authenticator.GetUser(token)
			if testCase.hasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, testCase.expUser, gotUser)
		})
	}
}

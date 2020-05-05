// +build !integration all

package authenticator

import (
	"testing"
	"time"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/crypto"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
)

func TestAuthenticator_GenerateToken(t *testing.T) {
	tokenizer := crypto.NewTokenizerFake()
	expIssuedAt := time.Now()
	tm := timer.NewStub(expIssuedAt)
	authenticator := NewAuthenticator(tokenizer, tm, 2*time.Millisecond)

	expUser := entity.User{
		Email: "test@s.time4hacks.com",
	}
	token, err := authenticator.GenerateToken(expUser)
	assert.Equal(t, nil, err)

	tokenPayload, err := tokenizer.Decode(token)
	assert.Equal(t, nil, err)

	assert.Equal(t, expUser.Email, tokenPayload["email"])

	expIssuedAtStr := expIssuedAt.Format(time.RFC3339Nano)
	assert.Equal(t, expIssuedAtStr, tokenPayload["issued_at"])
}

func TestAuthenticator_IsSignedIn(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name               string
		expIssuedAt        time.Time
		tokenValidDuration time.Duration
		currentTime        time.Time
		tokenPayload       crypto.TokenPayload
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
			tokenizer := crypto.NewTokenizerFake()
			tm := timer.NewStub(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, tm, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			assert.Equal(t, nil, err)
			gotIsSignIn := authenticator.IsSignedIn(token)
			assert.Equal(t, testCase.expIsSignIn, gotIsSignIn)
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
		tokenPayload       crypto.TokenPayload
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
			tokenizer := crypto.NewTokenizerFake()
			tm := timer.NewStub(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, tm, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			assert.Equal(t, nil, err)
			gotUser, err := authenticator.GetUser(token)
			if testCase.hasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, testCase.expUser, gotUser)
		})
	}
}

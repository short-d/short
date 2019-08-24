package auth

import (
	"testing"
	"time"

	"github.com/byliuyang/app/mdtest"

	"github.com/byliuyang/app/fw"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticator_GenerateToken(t *testing.T) {
	tokenizer := mdtest.FakeCryptoTokenizer
	expIssuedAt := time.Now()
	timer := mdtest.NewFakeTimer(expIssuedAt)
	authenticator := NewAuthenticator(tokenizer, timer, 2*time.Millisecond)

	expEmail := "test@s.time4hacks.com"
	token, err := authenticator.GenerateToken(expEmail)
	assert.Nil(t, err)

	tokenPayload, err := tokenizer.Decode(token)
	assert.Nil(t, err)

	assert.Equal(t, expEmail, tokenPayload["email"])

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
			tokenizer := mdtest.FakeCryptoTokenizer
			timer := mdtest.NewFakeTimer(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, timer, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			assert.Nil(t, err)
			gotIsSignIn := authenticator.IsSignedIn(token)
			assert.Equal(t, testCase.expIsSignIn, gotIsSignIn)
		})
	}
}

func TestAuthenticator_GetUserEmail(t *testing.T) {
	now := time.Now()

	testCases := []struct {
		name               string
		expIssuedAt        time.Time
		tokenValidDuration time.Duration
		currentTime        time.Time
		tokenPayload       fw.TokenPayload
		hasErr             bool
		expEmail           string
	}{
		{
			name:               "Token payload empty",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload:       map[string]interface{}{},
			hasErr:             true,
			expEmail:           "",
		},
		{
			name:               "Token payload without email",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"issued_at": now.Format(time.RFC3339Nano),
			},
			hasErr:   true,
			expEmail: "",
		},
		{
			name:               "Token payload without issue_at",
			expIssuedAt:        now,
			tokenValidDuration: time.Hour,
			currentTime:        now.Add(30 * time.Minute),
			tokenPayload: map[string]interface{}{
				"email": "test@s.time4hacks.com",
			},
			hasErr:   true,
			expEmail: "",
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
			hasErr:   true,
			expEmail: "",
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
			hasErr:   true,
			expEmail: "",
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
			hasErr:   false,
			expEmail: "test@s.time4hacks.com",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			tokenizer := mdtest.FakeCryptoTokenizer
			timer := mdtest.NewFakeTimer(testCase.currentTime)
			authenticator := NewAuthenticator(tokenizer, timer, testCase.tokenValidDuration)

			token, err := tokenizer.Encode(testCase.tokenPayload)
			assert.Nil(t, err)
			gotEmail, err := authenticator.GetUserEmail(token)
			if testCase.hasErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, testCase.expEmail, gotEmail)
		})
	}
}

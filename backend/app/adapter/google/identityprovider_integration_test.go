// +build integration all

package google

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/short-d/app/fw/webreq"

	"github.com/short-d/app/fw/assert"
)

func TestIdentityProvider_GetAuthorizationURL(t *testing.T) {
	t.Parallel()
	httpRequest := webreq.NewHTTPFake(
		func(req *http.Request) (response *http.Response, e error) {
			return nil, nil
		})
	clientID := "id_12345"
	clientSecret := "client_secret"
	redirectURI := "http://localhost/oauth/google/sign-in/callback"
	identityProvider := NewIdentityProvider(httpRequest, clientID, clientSecret, redirectURI)

	urlResponse := identityProvider.GetAuthorizationURL()

	parsedUrl, err := url.Parse(urlResponse)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https", parsedUrl.Scheme)
	assert.Equal(t, "accounts.google.com", parsedUrl.Host)
	assert.Equal(t, "/o/oauth2/v2/auth", parsedUrl.Path)
	assert.Equal(t, "true", parsedUrl.Query().Get("include_granted_scopes"))
	assert.Equal(t, "email profile", parsedUrl.Query().Get("scope"))
	assert.Equal(t, "code", parsedUrl.Query().Get("response_type"))
	assert.Equal(t, clientID, parsedUrl.Query().Get("client_id"))
	assert.Equal(t, redirectURI, parsedUrl.Query().Get("redirect_uri"))
}

func TestIdentityProvider_RequestAccessToken(t *testing.T) {
	testCases := []struct {
		name                string
		httpResponse        *http.Response
		httpErr             error
		clientID            string
		clientSecret        string
		redirectURI         string
		authorizationCode   string
		expectHasErr        bool
		expectedAccessToken string
	}{
		{
			name:                "invalid authorization code",
			httpResponse:        nil,
			httpErr:             errors.New("invalid authorization code"),
			clientID:            "id_12345",
			clientSecret:        "client_secret",
			redirectURI:         "http://localhost/oauth/google/sign-in/callback",
			authorizationCode:   "invalidCode",
			expectHasErr:        true,
			expectedAccessToken: "",
		},
		{
			name:                "invalid clientID",
			httpResponse:        nil,
			httpErr:             errors.New("invalid clientID"),
			clientID:            "invalidID",
			clientSecret:        "client_secret",
			redirectURI:         "http://localhost/oauth/google/sign-in/callback",
			authorizationCode:   "authorizationCode_1",
			expectHasErr:        true,
			expectedAccessToken: "",
		},
		{
			name:                "invalid client secret",
			httpResponse:        nil,
			httpErr:             errors.New("invalid clientSecret"),
			clientID:            "id_12345",
			clientSecret:        "invalidSecret",
			redirectURI:         "http://localhost/oauth/google/sign-in/callback",
			authorizationCode:   "authorizationCode_1",
			expectHasErr:        true,
			expectedAccessToken: "",
		},
		{
			name:                "invalid redirect URI",
			httpResponse:        nil,
			httpErr:             errors.New("invalid redirectURI"),
			clientID:            "id_12345",
			clientSecret:        "client_secret",
			redirectURI:         "http://localhost/invalidRedirectURI",
			authorizationCode:   "authorizationCode_1",
			expectHasErr:        true,
			expectedAccessToken: "",
		},
		{
			name: "success",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
      "access_token": "bcBi3AMeOV3Zg3AlOPyn",
      "token_type": "bearer"
}
`,
				)))},
			httpErr:             nil,
			clientID:            "id_12345",
			clientSecret:        "client_secret",
			redirectURI:         "http://localhost/oauth/google/sign-in/callback",
			authorizationCode:   "authorizationCode_1",
			expectHasErr:        false,
			expectedAccessToken: "bcBi3AMeOV3Zg3AlOPyn",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			httpRequest := webreq.NewHTTPFake(
				func(req *http.Request) (response *http.Response, e error) {
					assert.Equal(t, "https", req.URL.Scheme)
					assert.Equal(t, "www.googleapis.com", req.URL.Host)
					assert.Equal(t, "/oauth2/v4/token", req.URL.Path)
					assert.Equal(t, testCase.clientID, req.URL.Query().Get("client_id"))
					assert.Equal(t, testCase.clientSecret, req.URL.Query().Get("client_secret"))
					assert.Equal(t, testCase.authorizationCode, req.URL.Query().Get("code"))
					assert.Equal(t, testCase.redirectURI, req.URL.Query().Get("redirect_uri"))
					assert.Equal(t, "authorization_code", req.URL.Query().Get("grant_type"))
					assert.Equal(t, "POST", req.Method)
					assert.Equal(t, "application/json", req.Header.Get("Accept"))
					assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

					return testCase.httpResponse, testCase.httpErr
				})
			identityProvider := NewIdentityProvider(httpRequest, testCase.clientID, testCase.clientSecret, testCase.redirectURI)

			actualAccessToken, err := identityProvider.RequestAccessToken(testCase.authorizationCode)

			if testCase.expectHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedAccessToken, actualAccessToken)
		})
	}
}

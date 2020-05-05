// +build integration all

package github

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/app/mdtest"
)

func TestIdentityProvider_GetAuthorizationURL(t *testing.T) {
	t.Parallel()
	httpRequest := webreq.NewHTTPFake(
		func(req *http.Request) (response *http.Response, e error) {
			return nil, nil
		})
	clientID := "id_12345"
	clientSecret := "client_secret"
	identityProvider := NewIdentityProvider(httpRequest, clientID, clientSecret)

	urlResponse := identityProvider.GetAuthorizationURL()

	parsedUrl, err := url.Parse(urlResponse)
	assert.Equal(t, nil, err)
	assert.Equal(t, "https", parsedUrl.Scheme)
	assert.Equal(t, "github.com", parsedUrl.Host)
	assert.Equal(t, "/login/oauth/authorize", parsedUrl.Path)
	assert.Equal(t, clientID, parsedUrl.Query().Get("client_id"))
	assert.Equal(t, "read:user", parsedUrl.Query().Get("scope"))
}

func TestIdentityProvider_RequestAccessToken(t *testing.T) {
	testCases := []struct {
		name                string
		httpResponse        *http.Response
		httpErr             error
		clientID            string
		clientSecret        string
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
      "scope": "read:user",
      "token_type": "bearer"
}
`,
				)))},
			httpErr:             nil,
			clientID:            "id_12345",
			clientSecret:        "client_secret",
			authorizationCode:   "authorizationCode_1",
			expectHasErr:        false,
			expectedAccessToken: "bcBi3AMeOV3Zg3AlOPyn",
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			httpRequest := mdtest.NewHTTPRequestFake(
				func(req *http.Request) (response *http.Response, e error) {
					assert.Equal(t, "https", req.URL.Scheme)
					assert.Equal(t, "github.com", req.URL.Host)
					assert.Equal(t, "/login/oauth/access_token", req.URL.Path)
					assert.Equal(t, "POST", req.Method)
					assert.Equal(t, "application/json", req.Header.Get("Accept"))
					assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

					return testCase.httpResponse, testCase.httpErr
				})
			identityProvider := NewIdentityProvider(httpRequest, testCase.clientID, testCase.clientSecret)

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

// +build integration all

package google

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/short-d/app/mdtest"
)

func TestIdentityProvider_GetAuthorizationURL(t *testing.T) {
	t.Parallel()
	httpRequest := mdtest.NewHTTPRequestFake(
		func(req *http.Request) (response *http.Response, e error) {
			return nil, nil
		})
	clientID := "id_12345"
	clientSecret := "client_secret"
	redirectURI := "http://localhost/oauth/google/sign-in/callback"
	identityProvider := NewIdentityProvider(httpRequest, clientID, clientSecret, redirectURI)

	url := identityProvider.GetAuthorizationURL()

	mdtest.Equal(t, "https://accounts.google.com/o/oauth2/v2/auth?client_id=id_12345&include_granted_scopes=true&"+
		"redirect_uri=http%3A%2F%2Flocalhost%2Foauth%2Fgoogle%2Fsign-in%2Fcallback&response_type=code&scope=email+profile", url)
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
		t.Run(testCase.name, func(t *testing.T) {
			httpRequest := mdtest.NewHTTPRequestFake(
				func(req *http.Request) (response *http.Response, e error) {
					mdtest.Equal(t, "https://www.googleapis.com/oauth2/v4/token?"+
						"client_id="+testCase.clientID+"&client_secret="+testCase.clientSecret+"&code="+testCase.authorizationCode+
						"&grant_type=authorization_code&redirect_uri="+url.QueryEscape(testCase.redirectURI), req.URL.String())
					mdtest.Equal(t, "POST", req.Method)
					mdtest.Equal(t, "application/json", req.Header.Get("Accept"))
					mdtest.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

					return testCase.httpResponse, testCase.httpErr
				})
			identityProvider := NewIdentityProvider(httpRequest, testCase.clientID, testCase.clientSecret, testCase.redirectURI)

			actualAccessToken, err := identityProvider.RequestAccessToken(testCase.authorizationCode)

			if testCase.expectHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedAccessToken, actualAccessToken)
		})
	}
}

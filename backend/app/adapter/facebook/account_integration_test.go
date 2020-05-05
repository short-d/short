// +build integration all

package facebook

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/app/entity"
)

func TestAccount_GetSingleSignOnUser(t *testing.T) {
	testCases := []struct {
		name            string
		httpResponse    *http.Response
		httpErr         error
		expectHasErr    bool
		expectedSSOUser entity.SSOUser
	}{
		{
			name:         "invalid access token",
			httpResponse: nil,
			httpErr:      errors.New("invalid access token"),
			expectHasErr: true,
		},
		{
			name: "user has id, email and name",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
      "id": "12321321312312",
      "name": "Facebook User",
      "email": "facebookUser@gmail.com"
}
`,
				)))},
			expectHasErr: false,
			expectedSSOUser: entity.SSOUser{
				ID:    "12321321312312",
				Name:  "Facebook User",
				Email: "facebookUser@gmail.com",
			},
		},
		{
			name: "user doesn't have email",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
      "name": "Facebook User",
      "email": ""
}
`,
				)))},
			expectHasErr: false,
			expectedSSOUser: entity.SSOUser{
				Name:  "Facebook User",
				Email: "",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			httpRequest := webreq.NewHTTPFake(
				func(req *http.Request) (response *http.Response, e error) {
					assert.Equal(t, "https", req.URL.Scheme)
					assert.Equal(t, "graph.facebook.com", req.URL.Host)
					assert.Equal(t, "/me", req.URL.Path)
					assert.Equal(t, "access_token", req.URL.Query().Get("access_token"))
					assert.Equal(t, "id,name,email", req.URL.Query().Get("fields"))
					assert.Equal(t, "GET", req.Method)

					return testCase.httpResponse, testCase.httpErr
				})
			facebookAccount := NewAccount(httpRequest)

			gotSSOUser, err := facebookAccount.GetSingleSignOnUser("access_token")

			if testCase.expectHasErr {
				assert.NotEqual(t, nil, err)
				return
			}
			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedSSOUser, gotSSOUser)
		})
	}
}

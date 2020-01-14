// +build integration all

package github

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/short-d/app/mdtest"
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
			name: "user has id, email, and name",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
  "data": {
    "viewer": {
      "id": "pwBi3AMeOV3Zg3AlOPyn",
      "name": "Github User",
      "email": "github-user@gmail.com"
    }
  }
}
`,
				)))},
			expectHasErr: false,
			expectedSSOUser: entity.SSOUser{
				ID:    "pwBi3AMeOV3Zg3AlOPyn",
				Name:  "Github User",
				Email: "github-user@gmail.com",
			},
		},
		{
			name: "user doesn't have email",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
  "data": {
    "viewer": {
      "id": "pwBi3AMeOV3Zg3AlOPyn",
      "name": "Github User",
      "email": ""
    }
  }
}
`,
				)))},
			expectHasErr: false,
			expectedSSOUser: entity.SSOUser{
				ID:    "pwBi3AMeOV3Zg3AlOPyn",
				Name:  "Github User",
				Email: "",
			},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			graphQLRequest := mdtest.NewGraphQLRequestFake(
				func(req *http.Request) (response *http.Response, e error) {
					mdtest.Equal(t, "https://api.github.com/graphql", req.URL.String())
					mdtest.Equal(t, "POST", req.Method)
					mdtest.Equal(t, "application/json", req.Header.Get("Content-Type"))
					mdtest.Equal(t, "application/json", req.Header.Get("Accept"))

					return testCase.httpResponse, testCase.httpErr
				})
			githubAccount := NewAccount(graphQLRequest)

			gotSSOUser, err := githubAccount.GetSingleSignOnUser("access_token")
			if testCase.expectHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expectedSSOUser, gotSSOUser)
		})
	}
}

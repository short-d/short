package oauth

import (
	"errors"
	"net/http"
	"testing"

	"github.com/byliuyang/app/mdtest"
	"github.com/byliuyang/app/modern/mdrequest"
)

func TestFacebook_GetAuthorizationURL(t *testing.T) {
	facebook := Facebook{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		redirectURI:  "https://some.uri/",
	}

	mdtest.Equal(
		t,
		facebook.GetAuthorizationURL(),
		"https://www.facebook.com/v4.0/dialog/oauth?client_id=clientID&"+
			"redirect_uri=https%3A%2F%2Fsome.uri%2F"+
			"&scope=public_profile%2Cemail&response_type=code")
}

func TestFacebook_RequestAccessToken(t *testing.T) {
	transport := mdtest.NewTransportMock(func(req *http.Request) (response *http.Response, e error) {
		expectedURL := "https://graph.facebook.com/v4.0/oauth/access_token" +
			"?client_id=clientID&client_secret=clientSecret&code=code" +
			"&redirect_uri=https%3A%2F%2Fsome.uri%2F"
		mdtest.Equal(t, expectedURL, req.URL.String())
		mdtest.Equal(t, "GET", req.Method)
		mdtest.Equal(t, "application/json", req.Header.Get("Accept"))

		return mdtest.JSONResponse(map[string]interface{}{
			"access_token": "token",
			"token_type":   "type",
			"expires_in":   123,
		})
	})

	client := http.Client{
		Transport: transport,
	}
	req := mdrequest.NewHTTP(client)

	facebook := Facebook{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		redirectURI:  "https://some.uri/",
		http:         req,
	}

	accessToken, _ := facebook.RequestAccessToken("code")

	mdtest.Equal(t, "token", accessToken)
}

type fakeHTTP struct{}

func (h fakeHTTP) JSON(string, string, map[string]string, string, interface{}) error {
	return errors.New("oops")
}

func TestFacebook_RequestAccessTokenErr(t *testing.T) {
	facebook := Facebook{
		clientID:     "clientID",
		clientSecret: "clientSecret",
		redirectURI:  "https://some.uri/",
		http:         fakeHTTP{},
	}

	_, err := facebook.RequestAccessToken("code")

	mdtest.Equal(t, errors.New("oops"), err)
}

func TestFacebook_RedirectURI(t *testing.T) {
	facebook := Facebook{redirectURI: "https://some.uri/"}
	redirectURI := facebook.RedirectURI()
	expectedURI := "https%3A%2F%2Fsome.uri%2F"

	mdtest.Equal(t, expectedURI, redirectURI)
}

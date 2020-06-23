package handle

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/usecase/sso"
)

// SSOSignIn redirects user to the sign in page.
func SSOSignIn(
	singleSignOn sso.SingleSignOn,
	webFrontendURL string,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		token := getToken(params)
		if singleSignOn.IsSignedIn(token) {
			http.Redirect(w, r, webFrontendURL, http.StatusSeeOther)
			return
		}
		signInLink := singleSignOn.GetSignInLink()
		http.Redirect(w, r, signInLink, http.StatusSeeOther)
	}
}

// SSOSignInCallback generates Short's authentication token given identity provider's authorization code.
func SSOSignInCallback(
	singleSignOn sso.SingleSignOn,
	webFrontendURL url.URL,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		code := params["code"]

		authToken, err := singleSignOn.SignIn(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		webFrontendURL = setToken(webFrontendURL, authToken)
		http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
	}
}

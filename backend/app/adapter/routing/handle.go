package routing

import (
	"net/http"
	netURL "net/url"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase"
	"github.com/short-d/short/app/usecase/auth"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/sso"
)

// NewOriginalURL translates alias to the original long link.
func NewOriginalURL(
	tracer fw.Tracer,
	useCase usecase.UseCase,
	webFrontendURL netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		trace := tracer.BeginTrace("OriginalURL")
		alias := params["alias"]

		useCase.ViewLongLink(alias, func() {
			serve404(w, r, webFrontendURL)
			trace.End()
		}, func(longLink string) {
			http.Redirect(w, r, longLink, http.StatusSeeOther)
			trace.End()
		})
	}
}

func serve404(w http.ResponseWriter, r *http.Request, webFrontendURL netURL.URL) {
	webFrontendURL.Path = "/404"
	http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
}

// NewSSOSignIn redirects user to the sign in page.
func NewSSOSignIn(
	logger fw.Logger,
	tracer fw.Tracer,
	identityProvider service.IdentityProvider,
	authenticator auth.Authenticator,
	webFrontendURL string,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		token := getToken(params)
		if authenticator.IsSignedIn(token) {
			http.Redirect(w, r, webFrontendURL, http.StatusSeeOther)
			return
		}
		signInLink := identityProvider.GetAuthorizationURL()
		http.Redirect(w, r, signInLink, http.StatusSeeOther)
	}
}

// NewSSOSignInCallback generates Short's authentication token given identity provider's authorization code.
func NewSSOSignInCallback(
	logger fw.Logger,
	tracer fw.Tracer,
	singleSignOn sso.SingleSignOn,
	webFrontendURL netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
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

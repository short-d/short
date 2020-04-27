package routing

import (
	"encoding/json"
	"net/http"
	netURL "net/url"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/request"

	"github.com/short-d/short/app/usecase/authenticator"
	"github.com/short-d/short/app/usecase/feature"
	"github.com/short-d/short/app/usecase/service"
	"github.com/short-d/short/app/usecase/sso"
	"github.com/short-d/short/app/usecase/url"
)

// NewOriginalURL translates alias to the original long link.
func NewOriginalURL(
	instrumentationFactory request.InstrumentationFactory,
	urlRetriever url.Retriever,
	timer fw.Timer,
	webFrontendURL netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		i := instrumentationFactory.NewHTTP(r)
		i.RedirectingAliasToLongLink(nil)

		alias := params["alias"]
		now := timer.Now()
		u, err := urlRetriever.GetURL(alias, &now)
		if err != nil {
			i.LongLinkRetrievalFailed(err)
			serve404(w, r, webFrontendURL)
			return
		}
		i.LongLinkRetrievalSucceed()

		originURL := u.OriginalURL
		http.Redirect(w, r, originURL, http.StatusSeeOther)
		i.RedirectedAliasToLongLink(nil)
	}
}

func serve404(w http.ResponseWriter, r *http.Request, webFrontendURL netURL.URL) {
	webFrontendURL.Path = "/404"
	http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
}

// NewSSOSignIn redirects user to the sign in page.
func NewSSOSignIn(
	identityProvider service.IdentityProvider,
	auth authenticator.Authenticator,
	webFrontendURL string,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		token := getToken(params)
		if auth.IsSignedIn(token) {
			http.Redirect(w, r, webFrontendURL, http.StatusSeeOther)
			return
		}
		signInLink := identityProvider.GetAuthorizationURL()
		http.Redirect(w, r, signInLink, http.StatusSeeOther)
	}
}

// NewSSOSignInCallback generates Short's authentication token given identity provider's authorization code.
func NewSSOSignInCallback(
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

// FeatureHandle retrieves the status of feature toggle.
func FeatureHandle(
	instrumentationFactory request.InstrumentationFactory,
	featureDecisionFactory feature.DecisionFactory,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		i := instrumentationFactory.NewHTTP(r)
		featureID := params["featureID"]

		decision := featureDecisionFactory.NewDecision(i)
		isEnable := decision.IsFeatureEnable(featureID)

		body, err := json.Marshal(isEnable)
		if err != nil {
			return
		}

		w.Write(body)
	}
}

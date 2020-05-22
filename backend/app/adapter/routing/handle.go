package routing

import (
	"encoding/json"
	"net/http"
	netURL "net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/sso"
	"github.com/short-d/short/backend/app/usecase/url"
)

// NewOriginalURL translates alias to the original long link.
func NewOriginalURL(
	instrumentationFactory request.InstrumentationFactory,
	urlRetriever url.Retriever,
	timer timer.Timer,
	webFrontendURL netURL.URL,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		alias := params["alias"]

		i := instrumentationFactory.NewHTTP(r)
		i.RedirectingAliasToLongLink(alias)

		now := timer.Now()
		u, err := urlRetriever.GetURL(alias, &now)
		if err != nil {
			i.LongLinkRetrievalFailed(err)
			serve404(w, r, webFrontendURL)
			return
		}
		i.LongLinkRetrievalSucceed()

		originURL := u.LongLink
		http.Redirect(w, r, originURL, http.StatusSeeOther)
		i.RedirectedAliasToLongLink(u)
	}
}

func serve404(w http.ResponseWriter, r *http.Request, webFrontendURL netURL.URL) {
	webFrontendURL.Path = "/404"
	http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
}

// NewSSOSignIn redirects user to the sign in page.
func NewSSOSignIn(
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

// NewSSOSignInCallback generates Short's authentication token given identity provider's authorization code.
func NewSSOSignInCallback(
	singleSignOn sso.SingleSignOn,
	webFrontendURL netURL.URL,
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

// FeatureHandle retrieves the status of feature toggle.
func FeatureHandle(
	instrumentationFactory request.InstrumentationFactory,
	featureDecisionMakerFactory feature.DecisionMakerFactory,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		i := instrumentationFactory.NewRequest()
		featureID := params["featureID"]

		decision := featureDecisionMakerFactory.NewDecision(i)
		isEnable := decision.IsFeatureEnable(featureID)

		body, err := json.Marshal(isEnable)
		if err != nil {
			return
		}

		w.Write(body)
	}
}

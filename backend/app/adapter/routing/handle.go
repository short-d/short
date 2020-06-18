package routing

import (
	"encoding/json"
	"net/http"
	netURL "net/url"
	"strings"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/feature"
	"github.com/short-d/short/backend/app/usecase/search"
	"github.com/short-d/short/backend/app/usecase/shortlink"
	"github.com/short-d/short/backend/app/usecase/sso"
)

// NewLongLink translates alias to the original long link.
func NewLongLink(
	instrumentationFactory request.InstrumentationFactory,
	shortLinkRetriever shortlink.Retriever,
	timer timer.Timer,
	webFrontendURL netURL.URL,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		alias := params["alias"]

		i := instrumentationFactory.NewHTTP(r)
		i.RedirectingAliasToLongLink(alias)

		now := timer.Now()
		s, err := shortLinkRetriever.GetShortLink(alias, &now)
		if err != nil {
			i.LongLinkRetrievalFailed(err)
			serve404(w, r, webFrontendURL)
			return
		}
		i.LongLinkRetrievalSucceed()

		longLink := s.LongLink
		http.Redirect(w, r, longLink, http.StatusSeeOther)
		i.RedirectedAliasToLongLink(s)
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
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		i := instrumentationFactory.NewRequest()
		featureID := params["featureID"]
		user := getUser(r, authenticator)

		decision := featureDecisionMakerFactory.NewDecision(i)
		isEnable := decision.IsFeatureEnable(featureID, user)

		body, err := json.Marshal(isEnable)
		if err != nil {
			return
		}

		w.Write(body)
	}
}

// SearchHandle fetches resources under certain criterias.
func SearchHandle(
	search search.Search,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		w.Write([]byte("not implemented"))
	}
}

func getUser(r *http.Request, authenticator authenticator.Authenticator) *entity.User {
	authToken := getBearerToken(r)
	user, err := authenticator.GetUser(authToken)
	if err != nil {
		return nil
	}
	return &user
}

// getBearerToken parses Authorization token with format "Bearer <token>"
func getBearerToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if len(authHeader) < 1 {
		return ""
	}
	words := strings.Split(authHeader, " ")
	if len(words) != 2 {
		return ""
	}
	if words[0] != "Bearer" {
		return ""
	}
	return words[1]
}

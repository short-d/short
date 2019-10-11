package routing

import (
	"net/http"
	netURL "net/url"
	"short/app/adapter/oauth"
	"short/app/usecase/auth"
	"short/app/usecase/signin"
	"short/app/usecase/url"

	"github.com/byliuyang/app/fw"
)

func NewOriginalURL(
	logger fw.Logger,
	tracer fw.Tracer,
	urlRetriever url.Retriever,
	timer fw.Timer,
	webFrontendURL *netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		trace := tracer.BeginTrace("OriginalURL")

		alias := params["alias"]

		trace1 := trace.Next("GetUrlAfter")
		u, err := urlRetriever.GetAfter(trace1, alias, timer.Now())
		trace1.End()

		if err != nil {
			logger.Error(err)
			serve404(w, r, *webFrontendURL)
			return
		}

		originURL := u.OriginalURL
		http.Redirect(w, r, originURL, http.StatusSeeOther)
		trace.End()
	}
}

func serve404(w http.ResponseWriter, r *http.Request, webFrontendURL netURL.URL) {
	webFrontendURL.Path = "/404"
	http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
}

func NewGithubSignIn(
	logger fw.Logger,
	tracer fw.Tracer,
	githubOAuth oauth.Github,
	authenticator auth.Authenticator,
	webFrontendURL *netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		token := getToken(r, params)
		if authenticator.IsSignedIn(token) {
			http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
			return
		}
		signInLink := githubOAuth.GetAuthorizationURL()
		http.Redirect(w, r, signInLink, http.StatusSeeOther)
	}
}

func NewGithubSignInCallback(
	logger fw.Logger,
	tracer fw.Tracer,
	oauthSignIn signin.OAuth,
	webFrontendURL *netURL.URL,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		code := params["code"]

		authToken, err := oauthSignIn.SignIn(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		setToken(w, webFrontendURL.Hostname(), authToken)
		http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
	}
}

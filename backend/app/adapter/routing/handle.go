package routing

import (
	"fmt"
	"net/http"
	netURL "net/url"
	"os"
	"path/filepath"
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
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		trace := tracer.BeginTrace("OriginalURL")

		alias := params["alias"]

		trace1 := trace.Next("GetUrlAfter")
		u, err := urlRetriever.GetAfter(trace1, alias, timer.Now())
		trace1.End()

		if err != nil {
			logger.Error(err)
			serve404(w, r)
			return
		}

		originURL := u.OriginalURL
		http.Redirect(w, r, originURL, http.StatusSeeOther)
		trace.End()
	}
}

func serve404(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404", http.StatusSeeOther)
}

func NewServeFile(logger fw.Logger, tracer fw.Tracer, wwwRoot string) fw.Handle {
	filePath := filePathBuilder(wwwRoot, "index.html")
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		path := filePath(r)
		logger.Info(fmt.Sprintf("serving %s from %s", path, wwwRoot))
		http.ServeFile(w, r, path)
	}
}

func filePathBuilder(rootDir string, indexPath string) func(r *http.Request) string {
	return func(r *http.Request) string {
		path := filepath.Join(rootDir, r.URL.Path)
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return filepath.Join(rootDir, indexPath)
		}
		return path
	}
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

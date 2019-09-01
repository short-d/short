package routing

import (
	"fmt"
	"net/http"
	"short/app/adapter/oauth"
	"short/app/usecase/auth"
	"short/app/usecase/signin"
	"short/app/usecase/url"
	"strings"

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

func getFilenameFromPath(path string, indexFile string) string {
	filePath := strings.Trim(path, "/")
	if filePath == "" {
		return indexFile
	}
	return filePath
}

func NewServeFile(logger fw.Logger, tracer fw.Tracer, wwwRoot string) fw.Handle {
	rootDir := http.Dir(wwwRoot)
	fs := http.FileServer(rootDir)

	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		fileName := getFilenameFromPath(r.URL.Path, "index.html")

		_, err := rootDir.Open(fileName)
		if err != nil {
			logger.Error(err)
			serve404(w, r)
			return
		}

		logger.Info(fmt.Sprintf("serving %s from %s", fileName, wwwRoot))
		fs.ServeHTTP(w, r)
	}
}

func NewGithubSignIn(
	logger fw.Logger,
	tracer fw.Tracer,
	githubOAuth oauth.Github,
	authenticator auth.Authenticator,
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		token := getToken(r, params)
		if authenticator.IsSignedIn(token) {
			http.Redirect(w, r, "/", http.StatusSeeOther)
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
) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		code := params["code"]

		authToken, err := oauthSignIn.SignIn(code)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		setToken(w, authToken)
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func serve404(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/404.html", http.StatusSeeOther)
}

package handle

import (
	"net/http"
	"net/url"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
)

func serve404(w http.ResponseWriter, r *http.Request, webFrontendURL url.URL) {
	webFrontendURL.Path = "/404"
	http.Redirect(w, r, webFrontendURL.String(), http.StatusSeeOther)
}

func getUser(r *http.Request, authenticator authenticator.Authenticator) *entity.User {
	authToken := getBearerToken(r)
	user, err := authenticator.GetUser(authToken)
	if err != nil {
		return nil
	}
	return &user
}

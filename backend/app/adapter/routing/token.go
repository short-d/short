package routing

import (
	"net/http"

	"github.com/byliuyang/app/fw"
)

func getToken(r *http.Request, params fw.Params) string {
	token := params["token"]
	if len(token) > 0 {
		return token
	}

	newToken, err := r.Cookie("token")
	if err != nil {
		return ""
	}
	return newToken.Value
}

func setToken(w http.ResponseWriter, domain string, token string) {
	tokenCookie := http.Cookie{
		Name:   "token",
		Domain: domain,
		Path:   "/",
		Value:  token,
	}
	http.SetCookie(w, &tokenCookie)
}

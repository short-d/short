package routing

import (
	"net/http"
	"short/fw"
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

func setToken(w http.ResponseWriter, token string) {
	tokenCookie := http.Cookie{
		Name:  "token",
		Path:  "/",
		Value: token,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &tokenCookie)
}

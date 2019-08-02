package fw

import "net/http"

type Params map[string]string

type Handle func(w http.ResponseWriter, r *http.Request, params Params)

type Route struct {
	Method      string
	MatchPrefix bool
	Path        string
	Handle      Handle
}

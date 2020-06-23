package handle

import (
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/usecase/search"
)

// Search fetches resources under certain criterias.
func Search(
	search search.Search,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		w.Write([]byte("not implemented"))
	}
}

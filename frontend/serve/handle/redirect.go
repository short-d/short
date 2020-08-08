package handle

import (
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/frontend/serve/shortapi"
	"github.com/short-d/short/frontend/serve/ssr"
)

// Redirect returns a route handler for URL redirection
func Redirect(redirectPage ssr.RedirectPage, gRPC shortapi.GRPC) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		alias := params["alias"]
		openGraphTags, err := gRPC.GetOpenGraphTags(alias)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		twitterTags, err := gRPC.GetTwitterTags(alias)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		page, err := redirectPage.Render(openGraphTags, twitterTags)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(page))
	}
}

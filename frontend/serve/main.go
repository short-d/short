package main

import (
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/service"
	"github.com/short-d/short/frontend/serve/shortapi"
	"github.com/short-d/short/frontend/serve/ssr"
)

func main() {
	gRPC, err := shortapi.NewGRPC("localhost", 8081)
	if err != nil {
		panic(err)
	}

	rootDir := "../build"
	redirectPage := ssr.NewRedirectPage(rootDir)
	routes := []router.Route{
		{
			Method: http.MethodGet,
			Path:   "/r/:alias",
			Handle: func(w http.ResponseWriter, r *http.Request, params router.Params) {
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
			},
		},
		{
			Method:      http.MethodGet,
			Path:        "/",
			MatchPrefix: true,
			Handle: func(w http.ResponseWriter, r *http.Request, params router.Params) {
				fs := http.FileServer(http.Dir(rootDir))
				fs.ServeHTTP(w, r)
			},
		},
	}
	routingService := service.
		NewRoutingBuilder("Short").
		Routes(routes).
		Build()
	routingService.StartAndWait(3000)
}

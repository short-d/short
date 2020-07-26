package main

import (
	"fmt"
	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/service"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "../build"
	routes := []router.Route{
		{
			Method: http.MethodGet,
			Path:   "/r/:alias",
			Handle: func(w http.ResponseWriter, r *http.Request, params router.Params) {
				alias := params["alias"]
				ssrVars := map[string]string{
					"OG_TITLE": alias,
					"OG_DESCRIPTION": "Custom Description",
					"OG_IMAGE": "https://unidoc.io/static/news/unioffice-msword-templates-gophers-main.png",
					"OG_URL": "https://custom.url",
					"TWITTER_SITE": "Twitter Site",
					"TWITTER_TITLE": "Twitter Title",
					"TWITTER_DESCRIPTION": "Twitter Description",
					"TWITTER_IMAGE": "https://unidoc.io/static/news/unioffice-msword-templates-gophers-main.png",
				}
				buf, err := ioutil.ReadFile(filepath.Join(rootDir, "index.html"))
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				page := string(buf)

				for key, val := range ssrVars {
					target := fmt.Sprintf("{{SSR_%s}}", key)
					page = strings.ReplaceAll(page, target, val)
				}

				w.WriteHeader(http.StatusOK)
				w.Write([]byte(page))
			},
		},
		{
			Method: http.MethodGet,
			Path:   "/",
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

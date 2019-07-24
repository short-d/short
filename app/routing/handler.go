package routing

import (
	"net/http"
	"time"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

func NewOriginalUrl(logger fw.Logger, tracer fw.Tracer, urlRetriever usecase.UrlRetriever) fw.Handler {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		alias := params["alias"]

		finish := tracer.Begin()
		url, err := urlRetriever.GetUrlAfter(alias, time.Now())
		finish("routing.handler.NewUrlAlias")

		if err != nil {
			http.Redirect(w, r, "/404", 301)
			logger.Error(err)
			return
		}

		originUrl := url.OriginalUrl
		http.Redirect(w, r, originUrl, 301)
	}
}

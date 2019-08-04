package routing

import (
	"fmt"
	"net/http"
	"short/app/usecase"
	"short/fw"
	"strings"
	"time"
)

func NewOriginalUrl(logger fw.Logger, tracer fw.Tracer, urlRetriever usecase.UrlRetriever) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		trace := tracer.BeginTrace("OriginalUrl")

		alias := params["alias"]

		trace1 := trace.Next("GetUrlAfter")
		url, err := urlRetriever.GetUrlAfter(trace1, alias, time.Now())
		trace1.End()

		if err != nil {
			http.Redirect(w, r, "/404", http.StatusSeeOther)
			logger.Error(err)
			return
		}

		originUrl := url.OriginalUrl

		http.Redirect(w, r, originUrl, http.StatusSeeOther)
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
	fs := http.FileServer(http.Dir(wwwRoot))

	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		fileName := getFilenameFromPath(r.URL.Path, "index.html")
		logger.Info(fmt.Sprintf("serving %s from %s", fileName, wwwRoot))

		fs.ServeHTTP(w, r)
	}
}

package routing

import (
	"short/app/usecase/repo"
	"short/app/usecase/url"
	"short/fw"
)

func NewShort(logger fw.Logger, tracer fw.Tracer, wwwRoot string, urlRepo repo.Url) []fw.Route {
	urlRetriever := url.NewRetrieverPersist(urlRepo)
	fileHandle := NewServeFile(logger, tracer, wwwRoot)

	return []fw.Route{
		{Method: "GET", Path: "/r/:alias", Handle: NewOriginalUrl(logger, tracer, urlRetriever)},
		{Method: "GET", MatchPrefix: true, Path: "/", Handle: fileHandle},
	}
}

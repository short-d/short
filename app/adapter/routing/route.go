package routing

import (
	"short/app/usecase/repo"
	"short/app/usecase/url"
	"short/fw"
)

type WwwRoot string

func NewShort(logger fw.Logger, tracer fw.Tracer, wwwRoot WwwRoot, urlRepo repo.Url) []fw.Route {
	urlRetriever := url.NewRetrieverPersist(urlRepo)
	fileHandle := NewServeFile(logger, tracer, string(wwwRoot))

	return []fw.Route{
		{Method: "GET", Path: "/r/:alias", Handle: NewOriginalUrl(logger, tracer, urlRetriever)},
		{Method: "GET", MatchPrefix: true, Path: "/", Handle: fileHandle},
	}
}

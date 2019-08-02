package routing

import (
	"tinyURL/app/entity"
	"tinyURL/app/repo"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

func NewTinyUrl(logger fw.Logger, tracer fw.Tracer, wwwRoot string) []fw.Route {
	urlRepo := repo.NewUrlFake(map[string]entity.Url{
		"220uFicCJj": {
			Alias:       "220uFicCJj",
			OriginalUrl: "http://www.google.com",
		},
		"yDOBcj5HIPbUAsw": {
			Alias:       "yDOBcj5HIPbUAsw",
			OriginalUrl: "http://www.facebook.com",
		},
	})

	urlRetriever := usecase.NewUrlRetrieverRepo(tracer, urlRepo)
	fileHandle := NewServeFile(logger, tracer, wwwRoot)

	return []fw.Route{
		{Method: "GET", Path: "/api/redirect/:alias", Handle: NewOriginalUrl(logger, tracer, urlRetriever)},
		{Method: "GET", MatchPrefix: true, Path: "/", Handle: fileHandle},
	}
}

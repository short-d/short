package routing

import (
	"tinyURL/app/entity"
	"tinyURL/app/repo"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

func NewTinyUrl(logger fw.Logger, tracer fw.Tracer) fw.Routes {
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
	return fw.Routes{
		"/redirect/{alias}": NewOriginalUrl(logger, tracer, urlRetriever),
	}
}

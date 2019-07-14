package resolver

import (
	"tinyURL/app/repo"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

type Resolver struct {
	Query
	Mutation
}

func NewResolver(logger fw.Logger, tracer fw.Tracer, urlRepo repo.Url) Resolver {
	urlRetriever := usecase.NewUrlRetrieverRepo(tracer, urlRepo)
	return Resolver{
		Query: NewQuery(logger, tracer, urlRetriever),
	}
}

package resolver

import (
	"database/sql"
	"tinyURL/app/repo"
	"tinyURL/app/usecase"
	"tinyURL/fw"
)

type Resolver struct {
	Query
	Mutation
}

func NewResolver(logger fw.Logger, tracer fw.Tracer, db *sql.DB) Resolver {
	urlRepo := repo.NewUrlSql(db)
	urlRetriever := usecase.NewUrlRetrieverRepo(tracer, urlRepo)
	return Resolver{
		Query: NewQuery(logger, tracer, urlRetriever),
	}
}

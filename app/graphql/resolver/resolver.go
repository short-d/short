package resolver

import (
	"database/sql"
	"short/app/repo"
	"short/app/usecase"
	"short/fw"
)

type Resolver struct {
	Query
	Mutation
}

func NewResolver(logger fw.Logger, tracer fw.Tracer, db *sql.DB) Resolver {
	urlRepo := repo.NewUrlSql(db)
	urlRetriever := usecase.NewUrlRetrieverRepo(urlRepo)
	return Resolver{
		Query: NewQuery(logger, tracer, urlRetriever),
	}
}

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
	urlRetriever := usecase.NewUrlRetrieverPersist(urlRepo)
	keyGen := usecase.NewKeyGenInMemory()
	urlCreator := usecase.NewUrlCreatorPersist(urlRepo, keyGen)
	return Resolver{
		Query:    NewQuery(logger, tracer, urlRetriever),
		Mutation: NewMutation(logger, tracer, urlCreator),
	}
}

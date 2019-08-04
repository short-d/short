package routing

import (
	"database/sql"
	"short/app/repo"
	"short/app/usecase"
	"short/fw"
)

func NewShort(logger fw.Logger, tracer fw.Tracer, wwwRoot string, db *sql.DB) []fw.Route {
	urlRepo := repo.NewUrlSql(db)
	urlRetriever := usecase.NewUrlRetrieverPersist(urlRepo)
	fileHandle := NewServeFile(logger, tracer, wwwRoot)

	return []fw.Route{
		{Method: "GET", Path: "/api/redirect/:alias", Handle: NewOriginalUrl(logger, tracer, urlRetriever)},
		{Method: "GET", MatchPrefix: true, Path: "/", Handle: fileHandle},
	}
}

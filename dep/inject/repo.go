package inject

import (
	"database/sql"
	"short/app/adapter/reposql"
	"short/app/usecase/repo"
)

func UrlRepoSql(db *sql.DB) repo.Url {
	url := reposql.NewUrl(db)
	return &url
}

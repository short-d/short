package inject

import (
	"database/sql"
	"short/app/adapter/reposql"
	"short/app/usecase/repo"
)

func URLRepoSQL(db *sql.DB) repo.URL {
	url := reposql.NewURL(db)
	return &url
}

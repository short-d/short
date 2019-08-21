package inject

import (
	"database/sql"
	"short/app/adapter/sqlrepo"
	"short/app/usecase/repo"
)

func URLRepoSQL(db *sql.DB) repo.URL {
	url := sqlrepo.NewURL(db)
	return &url
}

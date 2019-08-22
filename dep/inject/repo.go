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

func UserRepoSQL(db *sql.DB) repo.User {
	user := sqlrepo.NewUser(db)
	return &user
}

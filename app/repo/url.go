package repo

import (
	"database/sql"
	"fmt"
	"tinyURL/app/entity"
	"tinyURL/app/sqlfmt"

	"github.com/pkg/errors"
)

type Url interface {
	GetByAlias(alias string) (entity.Url, error)
}

type UrlSql struct {
	db *sql.DB
}

func NewUrlSql(db *sql.DB) UrlSql {
	return UrlSql{
		db: db,
	}
}

func (u *UrlSql) GetByAlias(alias string) (entity.Url, error) {
	row := u.db.QueryRow("SELECT * FROM Url WHERE alias=$1;", alias)

	fmt.Println(row)

	var originalUrl string
	var expireAt string
	var createdAt string
	var updatedAt string

	err := row.Scan(&alias, &originalUrl, &expireAt, &createdAt, &updatedAt)

	if err == sql.ErrNoRows {
		return entity.Url{}, errors.Errorf("url not found (alias=%s)", alias)
	}

	if err != nil {
		return entity.Url{}, errors.WithStack(err)
	}

	url := entity.Url{
		Alias:       alias,
		OriginalUrl: originalUrl,
		ExpireAt:    sqlfmt.MustParseDatetime(expireAt),
		CreatedAt:   sqlfmt.MustParseDatetime(createdAt),
		UpdatedAt:   sqlfmt.MustParseDatetime(updatedAt),
	}

	return url, nil
}

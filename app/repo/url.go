package repo

import (
	"database/sql"
	"fmt"
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/table"

	"github.com/pkg/errors"
)

type Url interface {
	GetByAlias(alias string) (entity.Url, error)
}

type UrlSql struct {
	db *sql.DB
}

func NewUrlSql(db *sql.DB) Url {
	return &UrlSql{
		db: db,
	}
}

func (u *UrlSql) GetByAlias(alias string) (entity.Url, error) {
	statement := fmt.Sprintf(`SELECT "%s","%s","%s","%s","%s" FROM "%s" WHERE "%s"=$1;`,
		table.Url.ColumnAlias,
		table.Url.ColumnOriginalUrl,
		table.Url.ColumnExpireAt,
		table.Url.ColumnCreatedAt,
		table.Url.ColumnUpdatedAt,
		table.Url.TableName,
		table.Url.ColumnAlias,
	)

	row := u.db.QueryRow(statement, alias)

	var originalUrl string
	var expireAt *time.Time
	var createdAt *time.Time
	var updatedAt *time.Time

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
		ExpireAt:    expireAt,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return url, nil
}

package repo

import (
	"database/sql"
	"fmt"
	"short/app/entity"
	"short/app/table"
	"time"

	"github.com/pkg/errors"
)

type Url interface {
	IsAliasExist(alias string) bool
	GetByAlias(alias string) (entity.Url, error)
	Create(url entity.Url) error
}

type UrlSql struct {
	db *sql.DB
}

func (u UrlSql) IsAliasExist(alias string) bool {
	query := fmt.Sprintf(`
SELECT %s 
FROM "%s" 
WHERE %s=$1;`,
		table.Url.ColumnAlias,
		table.Url.TableName,
		table.Url.ColumnAlias,
	)

	err := u.db.QueryRow(query, alias).Scan()
	return err != sql.ErrNoRows
}

func (u *UrlSql) Create(url entity.Url) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s","%s")
VALUES ($1, $2, $3, $4, $5);`,
		table.Url.TableName,
		table.Url.ColumnAlias,
		table.Url.ColumnOriginalUrl,
		table.Url.ColumnExpireAt,
		table.Url.ColumnCreatedAt,
		table.Url.ColumnUpdatedAt,
	)
	_, err := u.db.Exec(statement, url.Alias, url.OriginalUrl, url.ExpireAt, url.CreatedAt, url.UpdatedAt)
	return err
}

func NewUrlSql(db *sql.DB) Url {
	return &UrlSql{
		db: db,
	}
}

func (u *UrlSql) GetByAlias(alias string) (entity.Url, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s" 
FROM "%s" 
WHERE "%s"=$1;`,
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

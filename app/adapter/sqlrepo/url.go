package sqlrepo

import (
	"database/sql"
	"fmt"
	"short/app/adapter/sqlrepo/table"
	"short/app/entity"
	"short/app/usecase/repo"
)

var _ repo.URL = (*URL)(nil)

type URL struct {
	db *sql.DB
}

func (u URL) IsAliasExist(alias string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;`,
		table.URL.ColumnAlias,
		table.URL.TableName,
		table.URL.ColumnAlias,
	)

	err := u.db.QueryRow(query, alias).Scan(&alias)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, nil
	}
	return true, err
}

func (u *URL) Create(url entity.URL) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s","%s")
VALUES ($1, $2, $3, $4, $5);`,
		table.URL.TableName,
		table.URL.ColumnAlias,
		table.URL.ColumnOriginalURL,
		table.URL.ColumnExpireAt,
		table.URL.ColumnCreatedAt,
		table.URL.ColumnUpdatedAt,
	)
	_, err := u.db.Exec(statement, url.Alias, url.OriginalURL, url.ExpireAt, url.CreatedAt, url.UpdatedAt)
	return err
}

func (u URL) GetByAlias(alias string) (entity.URL, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s" 
FROM "%s" 
WHERE "%s"=$1;`,
		table.URL.ColumnAlias,
		table.URL.ColumnOriginalURL,
		table.URL.ColumnExpireAt,
		table.URL.ColumnCreatedAt,
		table.URL.ColumnUpdatedAt,
		table.URL.TableName,
		table.URL.ColumnAlias,
	)

	row := u.db.QueryRow(statement, alias)

	url := entity.URL{}
	err := row.Scan(&url.Alias, &url.OriginalURL, &url.ExpireAt, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

// GetByUser returns URLs created by user
func (u URL) GetByUser(email string) ([]entity.URL, error) {
	statement := fmt.Sprintf(`
SELECT  "%s","%s","%s","%s","%s"
FROM "%s"
LEFT JOIN "%s" ON  "%s"."%s" = "%s"."%s"
WHERE "%s"."%s" = $1;`,
		table.URL.ColumnAlias,
		table.URL.ColumnOriginalURL,
		table.URL.ColumnExpireAt,
		table.URL.ColumnCreatedAt,
		table.URL.ColumnUpdatedAt,

		table.UserURL.TableName,
		table.URL.TableName,

		table.UserURL.TableName,
		table.UserURL.ColumnUrlAlias,

		table.URL.TableName,
		table.URL.ColumnAlias,

		table.UserURL.TableName,
		table.UserURL.ColumnUserEmail,
	)

	urls := make([]entity.URL, 0)
	rows, err := u.db.Query(statement, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		url := entity.URL{}

		if err := rows.Scan(&url.Alias, &url.OriginalURL, &url.ExpireAt, &url.CreatedAt, &url.UpdatedAt); err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func NewURL(db *sql.DB) *URL {
	return &URL{
		db: db,
	}
}

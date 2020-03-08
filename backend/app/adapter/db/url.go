package db

import (
	"database/sql"
	"fmt"

	"github.com/short-d/short/app/adapter/db/table"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

var _ repository.URL = (*URLSql)(nil)

// URLSql accesses URL information in url table through SQL.
type URLSql struct {
	db *sql.DB
}

// IsAliasExist checks whether a given alias exist in url table.
func (u URLSql) IsAliasExist(alias string) (bool, error) {
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
		return false, err
	}
	return true, nil
}

// Create inserts a new URL into url table.
func (u *URLSql) Create(url entity.URL) error {
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
	_, err := u.db.Exec(
		statement,
		url.Alias,
		url.OriginalURL,
		url.ExpireAt,
		url.CreatedAt,
		url.UpdatedAt,
	)
	return err
}

// GetByAlias finds an URL in url table given alias.
func (u URLSql) GetByAlias(alias string) (entity.URL, error) {
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
	err := row.Scan(
		&url.Alias,
		&url.OriginalURL,
		&url.ExpireAt,
		&url.CreatedAt,
		&url.UpdatedAt,
	)
	if err != nil {
		return entity.URL{}, err
	}

	url.CreatedAt = utc(url.CreatedAt)
	url.UpdatedAt = utc(url.UpdatedAt)
	url.ExpireAt = utc(url.ExpireAt)

	return url, nil
}

// GetByAliases finds URLs for a list of aliases
func (u URLSql) GetByAliases(aliases []string) ([]entity.URL, error) {
	urls := make([]entity.URL, 0)

	for _, alias := range aliases {
		url, err := u.GetByAlias(alias)

		if err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

// NewURLSql creates URLSql
func NewURLSql(db *sql.DB) *URLSql {
	return &URLSql{
		db: db,
	}
}



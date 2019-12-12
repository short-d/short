package db

import (
	"database/sql"
	"fmt"
	"short/app/adapter/db/table"
	"short/app/entity"
	"short/app/usecase/repository"
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
		return false, nil
	}
	return true, err
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
	_, err := u.db.Exec(statement, url.Alias, url.OriginalURL, url.ExpireAt, url.CreatedAt, url.UpdatedAt)
	return err
}

// CreateWithTransaction executes a transaction statement which inserts a new URL into url table.
func (u *URLSql) CreateWithTransaction(tx *sql.Tx, url entity.URL) error {
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
	_, err := tx.Exec(statement, url.Alias, url.OriginalURL, url.ExpireAt, url.CreatedAt, url.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
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
	err := row.Scan(&url.Alias, &url.OriginalURL, &url.ExpireAt, &url.CreatedAt, &url.UpdatedAt)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

// NewURLSql creates URLSql
func NewURLSql(db *sql.DB) *URLSql {
	return &URLSql{
		db: db,
	}
}

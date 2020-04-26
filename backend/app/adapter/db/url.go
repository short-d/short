package db

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

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
// TODO(issue#698): change to CreateURL
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

// UpdateURL updates a URL that exists within the URL table and returns the newly updated URL if there is no error found while
// committing the update.
func (u *URLSql) UpdateURL(key string, newAlias string, newOriginalURL string, expireAt *time.Time) (entity.URL, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3
WHERE "%s"=$4;`,
		table.URL.TableName,
		table.URL.ColumnAlias,
		table.URL.ColumnOriginalURL,
		table.URL.ColumnExpireAt,
		key,
	)

	row := u.db.QueryRow(
		statement,
		newAlias,
		newOriginalURL,
		expireAt,
		key,
	)

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

	return url, nil
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
	if len(aliases) == 0 {
		return []entity.URL{}, nil
	}

	parameterStr := u.composeParamList(len(aliases))

	// create a list of interface{} to hold aliases for db.Query()
	aliasesInterface := []interface{}{}
	for _, alias := range aliases {
		aliasesInterface = append(aliasesInterface, alias)
	}

	var urls []entity.URL

	// TODO: compare performance between Query and QueryRow. Prefer QueryRow for readability
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s" 
FROM "%s"
WHERE "%s" IN (%s);`,
		table.URL.ColumnAlias,
		table.URL.ColumnOriginalURL,
		table.URL.ColumnExpireAt,
		table.URL.ColumnCreatedAt,
		table.URL.ColumnUpdatedAt,
		table.URL.TableName,
		table.URL.ColumnAlias,
		parameterStr,
	)

	stmt, err := u.db.Prepare(statement)
	if err != nil {
		return urls, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(aliasesInterface...)
	if err != nil {
		return urls, nil
	}

	defer rows.Close()
	for rows.Next() {
		url := entity.URL{}
		err := rows.Scan(
			&url.Alias,
			&url.OriginalURL,
			&url.ExpireAt,
			&url.CreatedAt,
			&url.UpdatedAt,
		)
		if err != nil {
			return urls, err
		}

		url.CreatedAt = utc(url.CreatedAt)
		url.UpdatedAt = utc(url.UpdatedAt)
		url.ExpireAt = utc(url.ExpireAt)

		urls = append(urls, url)
	}

	return urls, nil
}

// composeParamList converts an slice to a parameters string with format: $1, $2, $3, ...
func (u URLSql) composeParamList(numParams int) string {
	params := make([]string, 0, numParams)
	for i := 0; i < numParams; i++ {
		params = append(params, fmt.Sprintf("$%d", i+1))
	}

	parameterStr := strings.Join(params, ", ")
	return parameterStr
}

// NewURLSql creates URLSql
func NewURLSql(db *sql.DB) *URLSql {
	return &URLSql{
		db: db,
	}
}

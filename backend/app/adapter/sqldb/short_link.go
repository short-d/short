package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.ShortLink = (*ShortLinkSql)(nil)

// ShortLinkSql accesses ShortLink information in short_link table through SQL.
type ShortLinkSql struct {
	db *sql.DB
}

// IsAliasExist checks whether a given alias exist in short_link table.
func (s ShortLinkSql) IsAliasExist(alias string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
	)

	err := s.db.QueryRow(query, alias).Scan(&alias)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateShortLink inserts a new ShortLink into short_link table.
func (s *ShortLinkSql) CreateShortLink(shortLink entity.ShortLink) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s","%s", "%s", "%s", "%s", "%s", "%s", "%s")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnOGTitle,
		table.ShortLink.ColumnOGDescription,
		table.ShortLink.ColumnOGImageURL,
		table.ShortLink.ColumnTwitterTitle,
		table.ShortLink.ColumnTwitterDescription,
		table.ShortLink.ColumnTwitterImageURL,
	)
	_, err := s.db.Exec(
		statement,
		shortLink.Alias,
		shortLink.LongLink,
		shortLink.ExpireAt,
		shortLink.CreatedAt,
		shortLink.UpdatedAt,
		shortLink.OpenGraphTitle,
		shortLink.OpenGraphDescription,
		shortLink.OpenGraphImageURL,
		shortLink.TwitterTitle,
		shortLink.TwitterDescription,
		shortLink.TwitterImageURL,
	)
	return err
}

// UpdateShortLink updates a ShortLink that exists within the short_link table.
func (s *ShortLinkSql) UpdateShortLink(oldAlias string, newShortLink entity.ShortLink) (entity.ShortLink, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3, "%s"=$4
WHERE "%s"=$5;`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnAlias,
	)

	_, err := s.db.Exec(
		statement,
		newShortLink.Alias,
		newShortLink.LongLink,
		newShortLink.ExpireAt,
		newShortLink.UpdatedAt,
		oldAlias,
	)

	if err != nil {
		return entity.ShortLink{}, err
	}

	return newShortLink, nil
}

// GetShortLinkByAlias finds an ShortLink in short_link table given alias.
func (s ShortLinkSql) GetShortLinkByAlias(alias string) (entity.ShortLink, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s"
FROM "%s" 
WHERE "%s"=$1;`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnOGTitle,
		table.ShortLink.ColumnOGDescription,
		table.ShortLink.ColumnOGImageURL,
		table.ShortLink.ColumnTwitterTitle,
		table.ShortLink.ColumnTwitterDescription,
		table.ShortLink.ColumnTwitterImageURL,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
	)

	row := s.db.QueryRow(statement, alias)

	shortLink := entity.ShortLink{}
	err := row.Scan(
		&shortLink.Alias,
		&shortLink.LongLink,
		&shortLink.ExpireAt,
		&shortLink.CreatedAt,
		&shortLink.UpdatedAt,
		&shortLink.OpenGraphTitle,
		&shortLink.OpenGraphDescription,
		&shortLink.OpenGraphImageURL,
		&shortLink.TwitterTitle,
		&shortLink.TwitterDescription,
		&shortLink.TwitterImageURL,
	)
	if err != nil {
		return entity.ShortLink{}, err
	}

	shortLink.CreatedAt = utc(shortLink.CreatedAt)
	shortLink.UpdatedAt = utc(shortLink.UpdatedAt)
	shortLink.ExpireAt = utc(shortLink.ExpireAt)

	return shortLink, nil
}

// GetShortLinksByAliases finds ShortLinks for a list of aliases
func (s ShortLinkSql) GetShortLinksByAliases(aliases []string) ([]entity.ShortLink, error) {
	if len(aliases) == 0 {
		return []entity.ShortLink{}, nil
	}

	parameterStr := s.composeParamList(len(aliases))

	// create a list of interface{} to hold aliases for db.Query()
	aliasesInterface := []interface{}{}
	for _, alias := range aliases {
		aliasesInterface = append(aliasesInterface, alias)
	}

	var shortLinks []entity.ShortLink

	// TODO: compare performance between Query and QueryRow. Prefer QueryRow for readability
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s","%s","%s","%s" ,"%s","%s","%s" 
FROM "%s"
WHERE "%s" IN (%s);`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnOGTitle,
		table.ShortLink.ColumnOGDescription,
		table.ShortLink.ColumnOGImageURL,
		table.ShortLink.ColumnTwitterTitle,
		table.ShortLink.ColumnTwitterDescription,
		table.ShortLink.ColumnTwitterImageURL,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		parameterStr,
	)

	stmt, err := s.db.Prepare(statement)
	if err != nil {
		return shortLinks, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(aliasesInterface...)
	if err != nil {
		return shortLinks, nil
	}

	defer rows.Close()
	for rows.Next() {
		shortLink := entity.ShortLink{}
		err := rows.Scan(
			&shortLink.Alias,
			&shortLink.LongLink,
			&shortLink.ExpireAt,
			&shortLink.CreatedAt,
			&shortLink.UpdatedAt,
			&shortLink.OpenGraphTitle,
			&shortLink.OpenGraphDescription,
			&shortLink.OpenGraphImageURL,
			&shortLink.TwitterTitle,
			&shortLink.TwitterDescription,
			&shortLink.TwitterImageURL,
		)
		if err != nil {
			return shortLinks, err
		}

		shortLink.CreatedAt = utc(shortLink.CreatedAt)
		shortLink.UpdatedAt = utc(shortLink.UpdatedAt)
		shortLink.ExpireAt = utc(shortLink.ExpireAt)

		shortLinks = append(shortLinks, shortLink)
	}

	return shortLinks, nil
}

// composeParamList converts an slice to a parameters string with format: $1, $2, $3, ...
func (s ShortLinkSql) composeParamList(numParams int) string {
	params := make([]string, 0, numParams)
	for i := 0; i < numParams; i++ {
		params = append(params, fmt.Sprintf("$%d", i+1))
	}

	parameterStr := strings.Join(params, ", ")
	return parameterStr
}

// NewShortLinkSql creates ShortLinkSql
func NewShortLinkSql(db *sql.DB) *ShortLinkSql {
	return &ShortLinkSql{
		db: db,
	}
}

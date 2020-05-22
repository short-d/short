package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.URL = (*ShortLinkSql)(nil)

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

// Create inserts a new ShortLink into short_link table.
// TODO(issue#698): change to CreateURL
func (s *ShortLinkSql) Create(shortLink entity.URL) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s","%s")
VALUES ($1, $2, $3, $4, $5);`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
	)
	_, err := s.db.Exec(
		statement,
		shortLink.Alias,
		shortLink.OriginalURL,
		shortLink.ExpireAt,
		shortLink.CreatedAt,
		shortLink.UpdatedAt,
	)
	return err
}

func (u *ShortLinkSql) UpdateOGMetaTags(alias string, metaOGTags entity.MetaOGTags) (entity.URL, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3
WHERE "%s"=$4;`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnOGTitle,
		table.ShortLink.ColumnOGDescription,
		table.ShortLink.ColumnOGImageURL,
		table.ShortLink.ColumnAlias,
	)

	_, err := u.db.Exec(
		statement,
		metaOGTags.OGTitle,
		metaOGTags.OGDescription,
		metaOGTags.OGImageURL,
		alias,
	)

	if err != nil {
		return entity.URL{}, err
	}

	return u.GetByAlias(alias)
}

func (u *ShortLinkSql) UpdateTwitterMetaTags(alias string, metaTwitterTags entity.MetaTwitterTags) (entity.URL, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3
WHERE "%s"=$4;`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnTwitterTitle,
		table.ShortLink.ColumnTwitterDescription,
		table.ShortLink.ColumnTwitterImageURL,
		table.ShortLink.ColumnAlias,
	)

	_, err := u.db.Exec(
		statement,
		metaTwitterTags.TwitterTitle,
		metaTwitterTags.TwitterDescription,
		metaTwitterTags.TwitterImageURL,
		alias,
	)

	if err != nil {
		return entity.URL{}, err
	}

	return u.GetByAlias(alias)
}

// UpdateURL updates a ShortLink that exists within the short_link table.
func (s *ShortLinkSql) UpdateURL(oldAlias string, newShortLink entity.URL) (entity.URL, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3, "%s"=$4
WHERE "%s"=$5;`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnUpdatedAt,
		oldAlias,
	)

	_, err := s.db.Exec(
		statement,
		newShortLink.Alias,
		newShortLink.OriginalURL,
		newShortLink.ExpireAt,
		newShortLink.UpdatedAt,
		oldAlias,
	)

	if err != nil {
		return entity.URL{}, err
	}

	return newShortLink, nil
}

// GetByAlias finds an ShortLink in short_link table given alias.
func (s ShortLinkSql) GetByAlias(alias string) (entity.URL, error) {
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

	shortLink := entity.URL{}
	err := row.Scan(
		&shortLink.Alias,
		&shortLink.OriginalURL,
		&shortLink.ExpireAt,
		&shortLink.CreatedAt,
		&shortLink.UpdatedAt,
		&shortLink.OGTitle,
		&shortLink.OGDescription,
		&shortLink.OGImageURL,
		&shortLink.TwitterTitle,
		&shortLink.TwitterDescription,
		&shortLink.TwitterImageURL,
	)
	if err != nil {
		return entity.URL{}, err
	}

	shortLink.CreatedAt = utc(shortLink.CreatedAt)
	shortLink.UpdatedAt = utc(shortLink.UpdatedAt)
	shortLink.ExpireAt = utc(shortLink.ExpireAt)

	return shortLink, nil
}

// GetByAliases finds ShortLinks for a list of aliases
func (s ShortLinkSql) GetByAliases(aliases []string) ([]entity.URL, error) {
	if len(aliases) == 0 {
		return []entity.URL{}, nil
	}

	parameterStr := s.composeParamList(len(aliases))

	// create a list of interface{} to hold aliases for db.Query()
	aliasesInterface := []interface{}{}
	for _, alias := range aliases {
		aliasesInterface = append(aliasesInterface, alias)
	}

	var shortLinks []entity.URL

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
		shortLink := entity.URL{}
		err := rows.Scan(
			&shortLink.Alias,
			&shortLink.OriginalURL,
			&shortLink.ExpireAt,
			&shortLink.CreatedAt,
			&shortLink.UpdatedAt,
			&shortLink.OGTitle,
			&shortLink.OGDescription,
			&shortLink.OGImageURL,
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

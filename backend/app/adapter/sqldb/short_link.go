package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.ShortLink = (*ShortLinkSQL)(nil)

// ShortLinkSQL accesses ShortLink information in short_link table through SQL.
type ShortLinkSQL struct {
	db *sql.DB
}

// UpdateOpenGraphTags updates OpenGraph meta tags for a given short link.
func (s ShortLinkSQL) UpdateOpenGraphTags(alias string, openGraphTags metatag.OpenGraph) (entity.ShortLink, error) {
	statement := fmt.Sprintf(`
UPDATE "%s"
SET "%s"=$1, "%s"=$2, "%s"=$3
WHERE "%s"=$4;`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnOpenGraphTitle,
		table.ShortLink.ColumnOpenGraphDescription,
		table.ShortLink.ColumnOpenGraphImageURL,
		table.ShortLink.ColumnAlias,
	)

	_, err := s.db.Exec(
		statement,
		openGraphTags.Title,
		openGraphTags.Description,
		openGraphTags.ImageURL,
		alias,
	)
	if err != nil {
		return entity.ShortLink{}, err
	}

	return s.GetShortLinkByAlias(alias)
}

// UpdateTwitterTags updates Twitter meta tags for a given short link.
func (s ShortLinkSQL) UpdateTwitterTags(alias string, twitterTags metatag.Twitter) (entity.ShortLink, error) {
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

	_, err := s.db.Exec(
		statement,
		twitterTags.Title,
		twitterTags.Description,
		twitterTags.ImageURL,
		alias,
	)
	if err != nil {
		return entity.ShortLink{}, err
	}

	return s.GetShortLinkByAlias(alias)
}

// IsAliasExist checks whether a given alias exist in short_link table.
func (s ShortLinkSQL) IsAliasExist(alias string) (bool, error) {
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
func (s ShortLinkSQL) CreateShortLink(shortLinkInput entity.ShortLinkInput) error {
	statement := fmt.Sprintf(`
INSERT INTO "%s" ("%s","%s","%s","%s")
VALUES ($1, $2, $3, $4);`,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
	)
	_, err := s.db.Exec(
		statement,
		shortLinkInput.CustomAlias,
		shortLinkInput.LongLink,
		shortLinkInput.ExpireAt,
		shortLinkInput.CreatedAt,
	)
	return err
}

// UpdateShortLink updates a ShortLink that exists within the short_link table.
func (s ShortLinkSQL) UpdateShortLink(oldAlias string, shortLinkInput entity.ShortLinkInput) (entity.ShortLink, error) {
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
		shortLinkInput.CustomAlias,
		shortLinkInput.LongLink,
		shortLinkInput.ExpireAt,
		shortLinkInput.UpdatedAt,
		oldAlias,
	)

	if err != nil {
		return entity.ShortLink{}, err
	}

	return entity.ShortLink{
		Alias: *shortLinkInput.CustomAlias,
		LongLink: *shortLinkInput.LongLink,
		ExpireAt: shortLinkInput.ExpireAt,
		UpdatedAt: shortLinkInput.UpdatedAt,
	}, nil
}

// GetShortLinkByAlias finds an ShortLink in short_link table given alias.
func (s ShortLinkSQL) GetShortLinkByAlias(alias string) (entity.ShortLink, error) {
	statement := fmt.Sprintf(`
SELECT "%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s"
FROM "%s" 
WHERE "%s"=$1;`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnOpenGraphTitle,
		table.ShortLink.ColumnOpenGraphDescription,
		table.ShortLink.ColumnOpenGraphImageURL,
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
		&shortLink.OpenGraphTags.Title,
		&shortLink.OpenGraphTags.Description,
		&shortLink.OpenGraphTags.ImageURL,
		&shortLink.TwitterTags.Title,
		&shortLink.TwitterTags.Description,
		&shortLink.TwitterTags.ImageURL,
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
func (s ShortLinkSQL) GetShortLinksByAliases(aliases []string) ([]entity.ShortLink, error) {
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
SELECT "%s","%s","%s","%s","%s","%s","%s","%s","%s","%s","%s" 
FROM "%s"
WHERE "%s" IN (%s);`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.ColumnLongLink,
		table.ShortLink.ColumnExpireAt,
		table.ShortLink.ColumnCreatedAt,
		table.ShortLink.ColumnUpdatedAt,
		table.ShortLink.ColumnOpenGraphTitle,
		table.ShortLink.ColumnOpenGraphDescription,
		table.ShortLink.ColumnOpenGraphImageURL,
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
			&shortLink.OpenGraphTags.Title,
			&shortLink.OpenGraphTags.Description,
			&shortLink.OpenGraphTags.ImageURL,
			&shortLink.TwitterTags.Title,
			&shortLink.TwitterTags.Description,
			&shortLink.TwitterTags.ImageURL,
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
func (s ShortLinkSQL) composeParamList(numParams int) string {
	params := make([]string, 0, numParams)
	for i := 0; i < numParams; i++ {
		params = append(params, fmt.Sprintf("$%d", i+1))
	}

	parameterStr := strings.Join(params, ", ")
	return parameterStr
}

// NewShortLinkSQL creates ShortLinkSQL
func NewShortLinkSQL(db *sql.DB) ShortLinkSQL {
	return ShortLinkSQL{
		db: db,
	}
}

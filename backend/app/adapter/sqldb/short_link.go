package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.URL = (*URLSql)(nil)

// URLSql accesses ShortLink information in short_link table through SQL.
type URLSql struct {
	db *sql.DB
}

// IsAliasExist checks whether a given alias exist in short_link table.
func (u URLSql) IsAliasExist(alias string) (bool, error) {
	query := fmt.Sprintf(`
SELECT "%s" 
FROM "%s" 
WHERE "%s"=$1;`,
		table.ShortLink.ColumnAlias,
		table.ShortLink.TableName,
		table.ShortLink.ColumnAlias,
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

// Create inserts a new ShortLink into short_link table.
// TODO(issue#698): change to CreateURL
func (u *URLSql) Create(url entity.URL) error {
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

func (u *URLSql) UpdateOGMetaTags(alias string, metaOGTags entity.MetaOGTags) (entity.URL, error) {
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

func (u *URLSql) UpdateTwitterMetaTags(alias string, metaTwitterTags entity.MetaTwitterTags) (entity.URL, error) {
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
func (u *URLSql) UpdateURL(oldAlias string, newURL entity.URL) (entity.URL, error) {
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

	_, err := u.db.Exec(
		statement,
		newURL.Alias,
		newURL.OriginalURL,
		newURL.ExpireAt,
		newURL.UpdatedAt,
		oldAlias,
	)

	if err != nil {
		return entity.URL{}, err
	}

	return newURL, nil
}

// GetByAlias finds an ShortLink in short_link table given alias.
func (u URLSql) GetByAlias(alias string) (entity.URL, error) {
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

	row := u.db.QueryRow(statement, alias)

	url := entity.URL{}
	err := row.Scan(
		&url.Alias,
		&url.OriginalURL,
		&url.ExpireAt,
		&url.CreatedAt,
		&url.UpdatedAt,
		&url.OGTitle,
		&url.OGDescription,
		&url.OGImageURL,
		&url.TwitterTitle,
		&url.TwitterDescription,
		&url.TwitterImageURL,
	)
	if err != nil {
		return entity.URL{}, err
	}

	url.CreatedAt = utc(url.CreatedAt)
	url.UpdatedAt = utc(url.UpdatedAt)
	url.ExpireAt = utc(url.ExpireAt)

	return url, nil
}

// GetByAliases finds ShortLinks for a list of aliases
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
			&url.OGTitle,
			&url.OGDescription,
			&url.OGImageURL,
			&url.TwitterTitle,
			&url.TwitterDescription,
			&url.TwitterImageURL,
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

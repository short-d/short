package repository

import (
	"database/sql"
	"errors"
	"short/app/entity"
)

var _ URL = (*URLFake)(nil)

// URLFake accesses URL information in url table through SQL.
type URLFake struct {
	urls map[string]entity.URL
}

func (u URLFake) CreateWithTransaction(tx *sql.Tx, url entity.URL) error {
	panic("implement me")
}

// IsAliasExist checks whether a given alias exist in url table.
func (u URLFake) IsAliasExist(alias string) (bool, error) {
	_, ok := u.urls[alias]
	return ok, nil
}

// Create inserts a new URL into url table.
func (u *URLFake) Create(url entity.URL) error {
	isExist, err := u.IsAliasExist(url.Alias)
	if err != nil {
		return err
	}
	if isExist {
		return errors.New("alias exists")
	}
	u.urls[url.Alias] = url
	return nil
}

// GetByAlias finds an URL in url table given alias.
func (u URLFake) GetByAlias(alias string) (entity.URL, error) {
	isExist, err := u.IsAliasExist(alias)
	if err != nil {
		return entity.URL{}, err
	}
	if !isExist {
		return entity.URL{}, errors.New("alias not found")
	}
	url := u.urls[alias]
	return url, nil
}

// NewURLFake creates in memory URL repository
func NewURLFake(urls map[string]entity.URL) URLFake {
	return URLFake{
		urls: urls,
	}
}

package repository

import (
	"errors"
	"time"

	"github.com/short-d/short/backend/app/entity"
)

var _ URL = (*URLFake)(nil)

// URLFake accesses ShortLink information in url table through SQL.
type URLFake struct {
	urls map[string]entity.ShortLink
}

// IsAliasExist checks whether a given alias exist in url table.
func (u URLFake) IsAliasExist(alias string) (bool, error) {
	_, ok := u.urls[alias]
	return ok, nil
}

// Create inserts a new ShortLink into url table.
func (u *URLFake) Create(url entity.ShortLink) error {
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

// GetByAlias finds an ShortLink in url table given alias.
func (u URLFake) GetByAlias(alias string) (entity.ShortLink, error) {
	isExist, err := u.IsAliasExist(alias)
	if err != nil {
		return entity.ShortLink{}, err
	}
	if !isExist {
		return entity.ShortLink{}, errors.New("alias not found")
	}
	url := u.urls[alias]
	return url, nil
}

// GetByAliases finds all ShortLink for a list of aliases
func (u URLFake) GetByAliases(aliases []string) ([]entity.ShortLink, error) {
	if len(aliases) == 0 {
		return []entity.ShortLink{}, nil
	}

	var urls []entity.ShortLink
	for _, alias := range aliases {
		url, err := u.GetByAlias(alias)

		if err != nil {
			return urls, err
		}
		urls = append(urls, url)
	}
	return urls, nil
}

// UpdateURL updates an existing ShortLink with new properties.
func (u URLFake) UpdateURL(oldAlias string, newURL entity.ShortLink) (entity.ShortLink, error) {
	prevURL, ok := u.urls[oldAlias]
	if !ok {
		return entity.ShortLink{}, errors.New("alias not found")
	}

	now := time.Now().UTC()
	createdBy := prevURL.CreatedBy
	createdAt := prevURL.CreatedAt
	return entity.ShortLink{
		Alias:     newURL.Alias,
		LongLink:  newURL.LongLink,
		ExpireAt:  newURL.ExpireAt,
		CreatedBy: createdBy,
		CreatedAt: createdAt,
		UpdatedAt: &now,
	}, nil
}

// NewURLFake creates in memory ShortLink repository
func NewURLFake(urls map[string]entity.ShortLink) URLFake {
	return URLFake{
		urls: urls,
	}
}

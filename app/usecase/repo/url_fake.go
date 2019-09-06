package repo

import (
	"errors"
	"fmt"
	"short/app/entity"
)

var _ URL = (*URLFake)(nil)

type URLFake struct {
	urls map[string]entity.URL
}

func (u *URLFake) IsAliasExist(alias string) (bool, error) {
	_, ok := u.urls[alias]
	return ok, nil
}

func (u *URLFake) Create(url entity.URL) error {
	alias := url.Alias
	url, ok := u.urls[alias]

	if ok {
		return errors.New(fmt.Sprintf("url exists (alias=%s)", alias))
	}

	u.urls[alias] = url
	return nil
}

func (u URLFake) GetByAlias(alias string) (entity.URL, error) {
	url, ok := u.urls[alias]

	if !ok {
		return entity.URL{}, errors.New(fmt.Sprintf("url not found (alias=%s)", alias))
	}

	return url, nil
}

func NewURLFake(urls map[string]entity.URL) URLFake {
	return URLFake{
		urls: urls,
	}
}

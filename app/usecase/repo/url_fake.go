package repo

import (
	"short/app/entity"

	"github.com/pkg/errors"
)

type UrlFake struct {
	urls map[string]entity.Url
}

func (u *UrlFake) IsAliasExist(alias string) (bool, error) {
	_, ok := u.urls[alias]
	return ok, nil
}

func (u *UrlFake) Create(url entity.Url) error {
	alias := url.Alias
	url, ok := u.urls[alias]

	if ok {
		return errors.Errorf("url exists (alias=%s)", alias)
	}

	u.urls[alias] = url
	return nil
}

func (u UrlFake) GetByAlias(alias string) (entity.Url, error) {
	url, ok := u.urls[alias]

	if !ok {
		return entity.Url{}, errors.Errorf("url not found (alias=%s)", alias)
	}

	return url, nil
}

func NewUrlFake(urls map[string]entity.Url) Url {
	return &UrlFake{
		urls: urls,
	}
}

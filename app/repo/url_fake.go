package repo

import (
	"github.com/pkg/errors"
	"tinyURL/app/entity"
)

type UrlFake struct {
	urls map[string]entity.Url
}

func (u UrlFake) GetByAlias(alias string) (entity.Url, error) {
	url, ok := u.urls[alias]

	if  !ok {
		return entity.Url{}, errors.Errorf("url not found (alias=%s)", alias)
	}

	return url, nil
}

func NewUrlFake(urls map[string]entity.Url) Url {
	return UrlFake{
		urls:urls,
	}
}


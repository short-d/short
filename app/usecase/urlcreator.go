package usecase

import (
	"short/app/entity"
	"short/app/repo"
)

type UrlCreator interface {
	CreateUrl(url entity.Url) (entity.Url, error)
	CreateUrlWithCustomAlias(url entity.Url, alias string) (entity.Url, error)
}

type UrlCreatorPersist struct {
	urlRepo repo.Url
	keyGen  KeyGenerator
}

func (a UrlCreatorPersist) CreateUrl(url entity.Url) (entity.Url, error) {
	randomAlias := a.keyGen.NewKey()
	return a.CreateUrlWithCustomAlias(url, randomAlias)
}

func (a UrlCreatorPersist) CreateUrlWithCustomAlias(url entity.Url, alias string) (entity.Url, error) {
	url.Alias = alias

	err := a.urlRepo.Create(url)
	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

func NewUrlCreatorPersist(urlRepo repo.Url, keyGen KeyGenerator) UrlCreator {
	return UrlCreatorPersist{
		urlRepo: urlRepo,
		keyGen:  keyGen,
	}
}

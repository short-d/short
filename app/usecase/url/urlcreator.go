package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
)

type ErrAliasExist string

func (e ErrAliasExist) Error() string {
	return string(e)
}

type Creator interface {
	Create(url entity.Url) (entity.Url, error)
	CreateWithCustomAlias(url entity.Url, alias string) (entity.Url, error)
}

type CreatorPersist struct {
	urlRepo repo.Url
	keyGen  keygen.KeyGenerator
}

func (a CreatorPersist) Create(url entity.Url) (entity.Url, error) {
	randomAlias := a.keyGen.NewKey()
	return a.CreateWithCustomAlias(url, randomAlias)
}

func (a CreatorPersist) CreateWithCustomAlias(url entity.Url, alias string) (entity.Url, error) {
	url.Alias = alias

	if a.urlRepo.IsAliasExist(alias) {
		return entity.Url{}, ErrAliasExist("usecase: url alias already exist")
	}

	err := a.urlRepo.Create(url)
	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

func NewCreatorPersist(urlRepo repo.Url, keyGen keygen.KeyGenerator) Creator {
	return CreatorPersist{
		urlRepo: urlRepo,
		keyGen:  keyGen,
	}
}

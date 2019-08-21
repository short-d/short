package url

import (
	"short/app/entity"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
)

var _ Creator = (*CreatorPersist)(nil)

type ErrAliasExist string

func (e ErrAliasExist) Error() string {
	return string(e)
}

type Creator interface {
	Create(url entity.URL) (entity.URL, error)
	CreateWithCustomAlias(url entity.URL, alias string) (entity.URL, error)
}

type CreatorPersist struct {
	urlRepo repo.URL
	keyGen  keygen.KeyGenerator
}

func (a CreatorPersist) Create(url entity.URL) (entity.URL, error) {
	randomAlias := a.keyGen.NewKey()
	return a.CreateWithCustomAlias(url, randomAlias)
}

func (a CreatorPersist) CreateWithCustomAlias(url entity.URL, alias string) (entity.URL, error) {
	url.Alias = alias

	isExist, err := a.urlRepo.IsAliasExist(alias)
	if err != nil {
		return entity.URL{}, err
	}

	if isExist {
		return entity.URL{}, ErrAliasExist("usecase: url alias already exist")
	}

	err = a.urlRepo.Create(url)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func NewCreatorPersist(urlRepo repo.URL, keyGen keygen.KeyGenerator) CreatorPersist {
	return CreatorPersist{
		urlRepo: urlRepo,
		keyGen:  keyGen,
	}
}

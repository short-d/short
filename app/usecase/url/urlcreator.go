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
	Create(url entity.URL, userEmail string) (entity.URL, error)
	CreateWithCustomAlias(url entity.URL, alias string, userEmail string) (entity.URL, error)
}

type CreatorPersist struct {
	urlRepo     repo.URL
	userURLRepo repo.UserURL
	keyGen      keygen.KeyGenerator
}

func (a CreatorPersist) Create(url entity.URL, userEmail string) (entity.URL, error) {
	randomAlias := a.keyGen.NewKey()
	return a.CreateWithCustomAlias(url, randomAlias, userEmail)
}

func (a CreatorPersist) CreateWithCustomAlias(url entity.URL, alias string, userEmail string) (entity.URL, error) {
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

	err = a.userURLRepo.CreateRelation(userEmail, url.Alias)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func NewCreatorPersist(
	urlRepo repo.URL,
	userURLRepo repo.UserURL,
	keyGen keygen.KeyGenerator,
) CreatorPersist {
	return CreatorPersist{
		urlRepo:     urlRepo,
		userURLRepo: userURLRepo,
		keyGen:      keyGen,
	}
}

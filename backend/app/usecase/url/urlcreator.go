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

// Creator represents a URL alias creator
type Creator interface {
	CreateURL(url entity.URL, userEmail string) (entity.URL, error)
	CreateURLWithCustomAlias(url entity.URL, alias string, userEmail string) (entity.URL, error)
}

// Creator represents a URL alias creator which persist the generated alias in
// the repository
type CreatorPersist struct {
	urlRepo             repo.URL
	userURLRelationRepo repo.UserURLRelation
	keyGen              keygen.KeyGenerator
}

func (a CreatorPersist) CreateURL(url entity.URL, userEmail string) (entity.URL, error) {
	key, err := a.keyGen.NewKey()
	if err != nil {
		return entity.URL{}, err
	}
	randomAlias := string(key)
	return a.CreateURLWithCustomAlias(url, randomAlias, userEmail)
}

func (a CreatorPersist) CreateURLWithCustomAlias(url entity.URL, alias string, userEmail string) (entity.URL, error) {
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

	err = a.userURLRelationRepo.CreateRelation(userEmail, url.Alias)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func NewCreatorPersist(
	urlRepo repo.URL,
	userURLRelationRepo repo.UserURLRelation,
	keyGen keygen.KeyGenerator,
) CreatorPersist {
	return CreatorPersist{
		urlRepo:             urlRepo,
		userURLRelationRepo: userURLRelationRepo,
		keyGen:              keyGen,
	}
}

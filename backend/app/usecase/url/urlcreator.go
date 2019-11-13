package url

import (
	"short/app/entity"
	"short/app/usecase/input"
	"short/app/usecase/keygen"
	"short/app/usecase/repo"
)

var _ Creator = (*CreatorPersist)(nil)

type ErrAliasExist string

func (e ErrAliasExist) Error() string {
	return string(e)
}

type ErrInvalidLongLink string

func (e ErrInvalidLongLink) Error() string {
	return string(e)
}

type ErrInvalidCustomAlias string

func (e ErrInvalidCustomAlias) Error() string {
	return string(e)
}

// Creator represents a URL alias creator
type Creator interface {
	CreateURL(url entity.URL, alias *string, user entity.User) (entity.URL, error)
}

// CreatorPersist represents a URL alias creator which persist the generated
// alias in the repository
type CreatorPersist struct {
	urlRepo             repo.URL
	userURLRelationRepo repo.UserURLRelation
	keyGen              keygen.KeyGenerator
	longLinkValidator   input.Validator
	aliasValidator      input.Validator
}

func (c CreatorPersist) CreateURL(url entity.URL, customAlias *string, user entity.User) (entity.URL, error) {
	longLink := url.OriginalURL
	if !c.longLinkValidator.IsValid(&longLink) {
		return entity.URL{}, ErrInvalidLongLink(longLink)
	}

	if customAlias == nil {
		return c.createURLWithAutoAlias(url, user)
	}

	if !c.aliasValidator.IsValid(customAlias) {
		return entity.URL{}, ErrInvalidCustomAlias(*customAlias)
	}
	return c.createURLWithCustomAlias(url, *customAlias, user)
}

// CreateURL persists a new url with a generated alias in the repository.
func (c CreatorPersist) createURLWithAutoAlias(url entity.URL, user entity.User) (entity.URL, error) {
	key, err := c.keyGen.NewKey()
	if err != nil {
		return entity.URL{}, err
	}
	randomAlias := string(key)
	return c.createURLWithCustomAlias(url, randomAlias, user)
}

// CreateURLWithCustomAlias persists a new url with a custom alias in
// the repository.
func (c CreatorPersist) createURLWithCustomAlias(url entity.URL, alias string, user entity.User) (entity.URL, error) {
	url.Alias = alias

	isExist, err := c.urlRepo.IsAliasExist(alias)
	if err != nil {
		return entity.URL{}, err
	}

	if isExist {
		return entity.URL{}, ErrAliasExist("url alias already exist")
	}

	err = c.urlRepo.Create(url)
	if err != nil {
		return entity.URL{}, err
	}

	err = c.userURLRelationRepo.CreateRelation(user, url.Alias)
	return url, err
}

// NewCreatorPersist creates CreatorPersist
func NewCreatorPersist(
	urlRepo repo.URL,
	userURLRelationRepo repo.UserURLRelation,
	keyGen keygen.KeyGenerator,
) CreatorPersist {
	return CreatorPersist{
		urlRepo:             urlRepo,
		userURLRelationRepo: userURLRelationRepo,
		keyGen:              keyGen,
		longLinkValidator:   input.NewLongLink(),
		aliasValidator:      input.NewCustomAlias(),
	}
}

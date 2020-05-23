package url

import (
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

var _ Creator = (*CreatorPersist)(nil)

// ErrAliasExist represents alias unavailable error
type ErrAliasExist string

func (e ErrAliasExist) Error() string {
	return string(e)
}

// ErrInvalidLongLink represents incorrect long link format error
type ErrInvalidLongLink string

func (e ErrInvalidLongLink) Error() string {
	return string(e)
}

// ErrInvalidCustomAlias represents incorrect custom alias format error
type ErrInvalidCustomAlias string

func (e ErrInvalidCustomAlias) Error() string {
	return string(e)
}

// ErrMaliciousLongLink represents malicious long link error
type ErrMaliciousLongLink string

func (e ErrMaliciousLongLink) Error() string {
	return string(e)
}

// Creator represents a ShortLink alias creator
type Creator interface {
	CreateURL(url entity.ShortLink, alias string, user entity.User, isPublic bool) (entity.ShortLink, error)
}

// CreatorPersist represents a ShortLink alias creator which persist the generated
// alias in the repository
type CreatorPersist struct {
	urlRepo             repository.URL
	userURLRelationRepo repository.UserURLRelation
	keyGen              keygen.KeyGenerator
	longLinkValidator   validator.LongLink
	aliasValidator      validator.CustomAlias
	timer               timer.Timer
	riskDetector        risk.Detector
}

// CreateURL persists a new url with a given or auto generated alias in the repository.
// TODO(issue#235): add functionality for public URLs
func (c CreatorPersist) CreateURL(url entity.ShortLink, customAlias string, user entity.User, isPublic bool) (entity.ShortLink, error) {
	longLink := url.LongLink
	if !c.longLinkValidator.IsValid(longLink) {
		return entity.ShortLink{}, ErrInvalidLongLink(longLink)
	}

	if c.riskDetector.IsURLMalicious(longLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(longLink)
	}

	if !c.isAliasProvided(customAlias) {
		return c.createURLWithAutoAlias(url, user)
	}

	if !c.aliasValidator.IsValid(customAlias) {
		return entity.ShortLink{}, ErrInvalidCustomAlias(customAlias)
	}
	return c.createURLWithCustomAlias(url, customAlias, user)
}

func (c CreatorPersist) isAliasProvided(customAlias string) bool {
	return customAlias != ""
}

func (c CreatorPersist) createURLWithAutoAlias(url entity.ShortLink, user entity.User) (entity.ShortLink, error) {
	key, err := c.keyGen.NewKey()
	if err != nil {
		return entity.ShortLink{}, err
	}
	randomAlias := string(key)
	return c.createURLWithCustomAlias(url, randomAlias, user)
}

func (c CreatorPersist) createURLWithCustomAlias(url entity.ShortLink, alias string, user entity.User) (entity.ShortLink, error) {
	url.Alias = alias

	isExist, err := c.urlRepo.IsAliasExist(alias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	if isExist {
		return entity.ShortLink{}, ErrAliasExist("url alias already exist")
	}

	now := c.timer.Now().UTC()
	url.CreatedAt = &now

	err = c.urlRepo.Create(url)
	if err != nil {
		return entity.ShortLink{}, err
	}

	err = c.userURLRelationRepo.CreateRelation(user, url)
	return url, err
}

// NewCreatorPersist creates CreatorPersist
func NewCreatorPersist(
	urlRepo repository.URL,
	userURLRelationRepo repository.UserURLRelation,
	keyGen keygen.KeyGenerator,
	longLinkValidator validator.LongLink,
	aliasValidator validator.CustomAlias,
	timer timer.Timer,
	riskDetector risk.Detector,
) CreatorPersist {
	return CreatorPersist{
		urlRepo:             urlRepo,
		userURLRelationRepo: userURLRelationRepo,
		keyGen:              keyGen,
		longLinkValidator:   longLinkValidator,
		aliasValidator:      aliasValidator,
		timer:               timer,
		riskDetector:        riskDetector,
	}
}

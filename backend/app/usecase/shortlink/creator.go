package shortlink

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
	CreateShortLink(shortLink entity.ShortLink, alias *string, user entity.User, isPublic bool) (entity.ShortLink, error)
}

// CreatorPersist represents a ShortLink alias creator which persist the generated
// alias in the repository
type CreatorPersist struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	keyGen            keygen.KeyGenerator
	longLinkValidator validator.LongLink
	aliasValidator    validator.CustomAlias
	timer             timer.Timer
	riskDetector      risk.Detector
}

// CreateShortLink persists a new short link with a given or auto generated alias in the repository.
// TODO(issue#235): add functionality for public ShortLinks
func (c CreatorPersist) CreateShortLink(shortLink entity.ShortLink, customAlias *string, user entity.User, isPublic bool) (entity.ShortLink, error) {
	longLink := shortLink.LongLink
	if !c.longLinkValidator.IsValid(&longLink) {
		return entity.ShortLink{}, ErrInvalidLongLink(longLink)
	}

	if c.riskDetector.IsURLMalicious(longLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(longLink)
	}

	if !c.isAliasProvided(customAlias) {
		return c.createShortLinkWithAutoAlias(shortLink, user)
	}

	if !c.aliasValidator.IsValid(customAlias) {
		return entity.ShortLink{}, ErrInvalidCustomAlias(*customAlias)
	}
	return c.createShortLinkWithCustomAlias(shortLink, *customAlias, user)
}

func (c CreatorPersist) isAliasProvided(customAlias *string) bool {
	return customAlias != nil && *customAlias != ""
}

func (c CreatorPersist) createShortLinkWithAutoAlias(shortLink entity.ShortLink, user entity.User) (entity.ShortLink, error) {
	key, err := c.keyGen.NewKey()
	if err != nil {
		return entity.ShortLink{}, err
	}
	randomAlias := string(key)
	return c.createShortLinkWithCustomAlias(shortLink, randomAlias, user)
}

func (c CreatorPersist) createShortLinkWithCustomAlias(shortLink entity.ShortLink, alias string, user entity.User) (entity.ShortLink, error) {
	shortLink.Alias = alias

	isExist, err := c.shortLinkRepo.IsAliasExist(alias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	if isExist {
		return entity.ShortLink{}, ErrAliasExist("short link alias already exist")
	}

	now := c.timer.Now().UTC()
	shortLink.CreatedAt = &now

	err = c.shortLinkRepo.CreateShortLink(shortLink)
	if err != nil {
		return entity.ShortLink{}, err
	}

	err = c.userShortLinkRepo.CreateRelation(user, shortLink)
	return shortLink, err
}

// NewCreatorPersist creates CreatorPersist
func NewCreatorPersist(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	keyGen keygen.KeyGenerator,
	longLinkValidator validator.LongLink,
	aliasValidator validator.CustomAlias,
	timer timer.Timer,
	riskDetector risk.Detector,
) CreatorPersist {
	return CreatorPersist{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		keyGen:            keyGen,
		longLinkValidator: longLinkValidator,
		aliasValidator:    aliasValidator,
		timer:             timer,
		riskDetector:      riskDetector,
	}
}

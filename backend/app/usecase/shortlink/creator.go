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
type ErrInvalidLongLink struct {
	LongLink  string
	Violation validator.Violation
}

func (e ErrInvalidLongLink) Error() string {
	return string(e.LongLink)
}

// ErrInvalidCustomAlias represents incorrect custom alias format error
type ErrInvalidCustomAlias struct {
	customAlias string
	Violation   validator.Violation
}

func (e ErrInvalidCustomAlias) Error() string {
	return string(e.customAlias)
}

// ErrMaliciousLongLink represents malicious long link error
type ErrMaliciousLongLink string

func (e ErrMaliciousLongLink) Error() string {
	return string(e)
}

// Creator represents a ShortLink alias creator
type Creator interface {
	CreateShortLink(shortLinkInput entity.ShortLinkInput, user entity.User, isPublic bool) (entity.ShortLink, error)
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
// TODO(issue#235): add functionality for public URLs
func (c CreatorPersist) CreateShortLink(shortLinkInput entity.ShortLinkInput, user entity.User, isPublic bool) (entity.ShortLink, error) {
	longLink := shortLinkInput.GetLongLink("")
	isValid, violation := c.longLinkValidator.IsValid(longLink)
	if !isValid {
		return entity.ShortLink{}, ErrInvalidLongLink{longLink, violation}
	}

	if c.riskDetector.IsURLMalicious(longLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(longLink)
	}

	customAlias := shortLinkInput.GetCustomAlias("")
	isValid, violation = c.aliasValidator.IsValid(customAlias)
	if !isValid {
		return entity.ShortLink{}, ErrInvalidCustomAlias{customAlias, violation}
	}

	if customAlias == "" {
		autoAlias, err := c.generateAlias()
		if err != nil {
			// TODO(issue#950) create error type for fail create auto alias
			return entity.ShortLink{}, err
		}
		customAlias = autoAlias
	}

	shortLinkInput.LongLink = &longLink
	shortLinkInput.CustomAlias = &customAlias

	return c.createShortLink(shortLinkInput, user)
}

func (c CreatorPersist) generateAlias() (string, error) {
	key, err := c.keyGen.NewKey()
	if err != nil {
		return "", err
	}
	return string(key), nil
}

func (c CreatorPersist) createShortLink(shortLinkInput entity.ShortLinkInput, user entity.User) (entity.ShortLink, error) {
	isExist, err := c.shortLinkRepo.IsAliasExist(*shortLinkInput.CustomAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	if isExist {
		return entity.ShortLink{}, ErrAliasExist("short link alias already exist")
	}

	now := c.timer.Now().UTC()
	shortLinkInput.CreatedAt = &now

	err = c.shortLinkRepo.CreateShortLink(shortLinkInput)
	if err != nil {
		return entity.ShortLink{}, err
	}

	err = c.userShortLinkRepo.CreateRelation(user, shortLinkInput)
	return entity.ShortLink{
		LongLink:  *shortLinkInput.LongLink,
		Alias:     *shortLinkInput.CustomAlias,
		ExpireAt:  shortLinkInput.ExpireAt,
		CreatedAt: shortLinkInput.CreatedAt,
	}, err
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

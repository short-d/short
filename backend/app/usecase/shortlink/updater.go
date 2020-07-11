package shortlink

import (
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

var _ Updater = (*UpdaterPersist)(nil)

// ErrShortLinkNotFound represents the failure of finding certain short link in the data store
type ErrShortLinkNotFound string

func (e ErrShortLinkNotFound) Error() string {
	return string(e)
}

// ErrNewAliasNotSpecified represents no alias specified for update.
type ErrNewAliasNotSpecified string

func (e ErrNewAliasNotSpecified) Error() string {
	return string(e)
}

// Updater mutates existing short links.
type Updater interface {
	UpdateShortLink(oldAlias string, shortLinkInput entity.ShortLinkInput, user entity.User) (entity.ShortLink, error)
}

// UpdaterPersist persists the mutated short link in the data store.
type UpdaterPersist struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	longLinkValidator validator.LongLink
	aliasValidator    validator.CustomAlias
	timer             timer.Timer
	riskDetector      risk.Detector
}

// UpdateShortLink mutates a short link in the repository.
func (u UpdaterPersist) UpdateShortLink(
	oldAlias string,
	shortLinkInput entity.ShortLinkInput,
	user entity.User,
) (entity.ShortLink, error) {
	hasMapping, err := u.userShortLinkRepo.HasMapping(user, oldAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}
	if !hasMapping {
		return entity.ShortLink{}, ErrShortLinkNotFound("short link not found")
	}

	newAlias := shortLinkInput.GetCustomAlias(oldAlias)
	if newAlias == "" {
		return entity.ShortLink{}, ErrNewAliasNotSpecified("new alias not specified")
	}

	// Only check if it exists if user is changing the alias to something else
	if newAlias != oldAlias {
		aliasExist, err := u.shortLinkRepo.IsAliasExist(newAlias)
		if err != nil {
			return entity.ShortLink{}, err
		}
		if aliasExist {
			return entity.ShortLink{}, ErrAliasExist("short link alias already exists")
		}
	}

	shortLink, err := u.shortLinkRepo.GetShortLinkByAlias(oldAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	shortLink.Alias = newAlias
	shortLink.LongLink = shortLinkInput.GetLongLink("")

	isValid, violation := u.aliasValidator.IsValid(shortLink.Alias)
	if !isValid {
		return entity.ShortLink{}, ErrInvalidCustomAlias{shortLink.Alias, violation}
	}

	isValid, violation = u.longLinkValidator.IsValid(shortLink.LongLink)
	if !isValid {
		return entity.ShortLink{}, ErrInvalidLongLink{shortLink.LongLink, violation}
	}

	if u.riskDetector.IsURLMalicious(shortLink.LongLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(shortLink.LongLink)
	}

	updateTime := u.timer.Now()
	shortLink.UpdatedAt = &updateTime

	return u.shortLinkRepo.UpdateShortLink(oldAlias, shortLink)
}

// NewUpdaterPersist creates a new UpdaterPersist instance.
func NewUpdaterPersist(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	longLinkValidator validator.LongLink,
	aliasValidator validator.CustomAlias,
	timer timer.Timer,
	riskDetector risk.Detector,
) UpdaterPersist {
	return UpdaterPersist{
		shortLinkRepo,
		userShortLinkRepo,
		longLinkValidator,
		aliasValidator,
		timer,
		riskDetector,
	}
}

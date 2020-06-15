package shortlink

import (
	"errors"

	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

var _ Updater = (*UpdaterPersist)(nil)

// ErrShortLinkNotFound represents the failure of finding certain short link in the data store
var ErrShortLinkNotFound = errors.New("short link not found")

// Updater mutates existing short links.
type Updater interface {
	UpdateShortLink(oldAlias string, update entity.ShortLink, user entity.User) (entity.ShortLink, error)
}

// UpdaterPersist persists the mutated short link in the data store..
type UpdaterPersist struct {
	shortLinkRepo             repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	keyGen                    keygen.KeyGenerator
	longLinkValidator         validator.LongLink
	aliasValidator            validator.CustomAlias
	timer                     timer.Timer
	riskDetector              risk.Detector
}

// UpdateShortLink persists mutations for a given short link in the repository.
func (u UpdaterPersist) UpdateShortLink(
	oldAlias string,
	update entity.ShortLink,
	user entity.User,
) (entity.ShortLink, error) {
	ok, err := u.userShortLinkRelationRepo.HasMapping(user, oldAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}
	if !ok {
		return entity.ShortLink{}, ErrShortLinkNotFound
	}

	shortLink, err := u.shortLinkRepo.GetShortLinkByAlias(oldAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	shortLink = u.updateAlias(shortLink, update)
	shortLink = u.updateLongLink(shortLink, update)

	if isValid, violation := u.aliasValidator.IsValid(&shortLink.Alias); !isValid {
		return entity.ShortLink{}, ErrInvalidCustomAlias{shortLink.Alias, violation}
	}

	if isValid, violation := u.longLinkValidator.IsValid(&shortLink.LongLink); !isValid {
		return entity.ShortLink{}, ErrInvalidLongLink{shortLink.LongLink, violation}
	}

	if u.riskDetector.IsURLMalicious(shortLink.LongLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(shortLink.LongLink)
	}

	updateTime := u.timer.Now()
	shortLink.UpdatedAt = &updateTime

	return u.shortLinkRepo.UpdateShortLink(oldAlias, shortLink)
}

func (u UpdaterPersist) updateAlias(shortLink, update entity.ShortLink) entity.ShortLink {
	newAlias := update.Alias
	if newAlias != "" {
		shortLink.Alias = newAlias
	}

	return shortLink
}

func (u *UpdaterPersist) updateLongLink(shortLink, update entity.ShortLink) entity.ShortLink {
	newLongLink := update.LongLink
	if newLongLink != "" {
		shortLink.LongLink = newLongLink
	}

	return shortLink
}

// NewUpdaterPersist creates a new UpdaterPersist instance.
func NewUpdaterPersist(
	shortLinkRepo repository.ShortLink,
	userShortLinkRelationRepo repository.UserShortLink,
	keyGen keygen.KeyGenerator,
	longLinkValidator validator.LongLink,
	aliasValidator validator.CustomAlias,
	timer timer.Timer,
	riskDetector risk.Detector,
) UpdaterPersist {
	return UpdaterPersist{
		shortLinkRepo,
		userShortLinkRelationRepo,
		keyGen,
		longLinkValidator,
		aliasValidator,
		timer,
		riskDetector,
	}
}

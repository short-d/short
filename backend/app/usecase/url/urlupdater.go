package url

import (
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

var _ Updater = (*UpdaterPersist)(nil)

// Updater represents a short link attribute updater.
type Updater interface {
	UpdateURL(oldAlias string, update entity.URL, user entity.User) (entity.URL, error)
}

// CreatorPersist represents a short link updater that persists the given changes
type UpdaterPersist struct {
	urlRepo             repository.URL
	userURLRelationRepo repository.UserURLRelation
	keyGen              keygen.KeyGenerator
	longLinkValidator   validator.LongLink
	aliasValidator      validator.CustomAlias
	timer               timer.Timer
	riskDetector        risk.Detector
}

// UpdateURL persists mutations for a given short link in the repository.
func (u UpdaterPersist) UpdateURL(
	oldAlias string,
	update entity.URL,
	user entity.User,
) (entity.URL, error) {
	url, err := u.urlRepo.GetByAlias(oldAlias)

	if err != nil {
		return entity.URL{}, err
	}

	url = u.updateAlias(url, update)
	url = u.updateLongLink(url, update)

	if u.riskDetector.IsURLMalicious(url.OriginalURL) {
		return entity.URL{}, ErrMaliciousLongLink(url.OriginalURL)
	}

	if !u.aliasValidator.IsValid(&url.Alias) {
		return entity.URL{}, ErrInvalidCustomAlias(url.Alias)
	}

	if !u.longLinkValidator.IsValid(&url.OriginalURL) {
		return entity.URL{}, ErrInvalidLongLink(url.OriginalURL)
	}

	return u.urlRepo.UpdateURL(oldAlias, url)
}

func (u UpdaterPersist) updateAlias(url, update entity.URL) entity.URL {
	newAlias := update.Alias

	if newAlias != "" {
		url.Alias = newAlias
	}

	return url
}

func (u *UpdaterPersist) updateLongLink(url, update entity.URL) entity.URL {
	newLongLink := update.OriginalURL

	if newLongLink != "" {
		url.OriginalURL = newLongLink
	}

	return url
}

// NewUpdaterPersist creates a new UpdaterPersist instance.
func NewUpdaterPersist(
	urlRepo repository.URL,
	userURLRelationRepo repository.UserURLRelation,
	keyGen keygen.KeyGenerator,
	longLinkValidator validator.LongLink,
	aliasValidator validator.CustomAlias,
	timer timer.Timer,
	riskDetector risk.Detector,
) UpdaterPersist {
	return UpdaterPersist{
		urlRepo,
		userURLRelationRepo,
		keyGen,
		longLinkValidator,
		aliasValidator,
		timer,
		riskDetector,
	}
}

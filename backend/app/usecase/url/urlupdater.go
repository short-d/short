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

type Updater interface {
	UpdateURL(oldAlias string, update entity.URL, user entity.User) (entity.URL, error)
}

type UpdaterPersist struct {
	urlRepo             repository.URL
	userURLRelationRepo repository.UserURLRelation
	keyGen              keygen.KeyGenerator
	longLinkValidator   validator.LongLink
	aliasValidator      validator.CustomAlias
	timer               timer.Timer
	riskDetector        risk.Detector
}

func (u UpdaterPersist) UpdateURL(
	oldAlias string,
	update entity.URL,
	user entity.User,
) (entity.URL, error) {
	oldURL, err := u.urlRepo.GetByAlias(oldAlias)

	if err != nil {
		return entity.URL{}, err
	}

	update = u.updateAlias(oldURL, update)
	update = u.updateLongLink(oldURL, update)

	if u.riskDetector.IsURLMalicious(update.OriginalURL) {
		return entity.URL{}, ErrMaliciousLongLink(update.OriginalURL)
	}

	if !u.aliasValidator.IsValid(&update.Alias) {
		return entity.URL{}, ErrInvalidCustomAlias(oldAlias)
	}

	if !u.longLinkValidator.IsValid(&update.OriginalURL) {
		return entity.URL{}, ErrInvalidLongLink(update.OriginalURL)
	}

	return u.urlRepo.UpdateURL(oldAlias, update)
}

func (u UpdaterPersist) updateAlias(url, update entity.URL) entity.URL {
	newAlias := update.Alias

	if newAlias == "" {
		update.Alias = url.Alias
	}

	return update
}

func (u *UpdaterPersist) updateLongLink(url, update entity.URL) entity.URL {
	newLongLink := update.OriginalURL

	if newLongLink == "" {
		url.OriginalURL = update.OriginalURL
	}

	return update
}

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

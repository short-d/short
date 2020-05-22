package url

import (
	"errors"
	"sort"

	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/keygen"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/risk"
	"github.com/short-d/short/backend/app/usecase/validator"
)

var _ Updater = (*UpdaterPersist)(nil)

// Updater modifies the properties of existing short links.
type Updater interface {
	UpdateURL(oldAlias string, update entity.URL, user entity.User) (entity.URL, error)
}

// UpdaterPersist persists the given changes to a short link in the database store.
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
	if !u.isURLRelated(oldAlias, user) {
		return entity.URL{}, errors.New("short link not found")
	}

	url, err := u.urlRepo.GetByAlias(oldAlias)
	if err != nil {
		return entity.URL{}, err
	}

	url = u.updateAlias(url, update)
	url = u.updateLongLink(url, update)

	if !u.aliasValidator.IsValid(&url.Alias) {
		return entity.URL{}, ErrInvalidCustomAlias(url.Alias)
	}

	if !u.longLinkValidator.IsValid(&url.OriginalURL) {
		return entity.URL{}, ErrInvalidLongLink(url.OriginalURL)
	}

	if u.riskDetector.IsURLMalicious(url.OriginalURL) {
		return entity.URL{}, ErrMaliciousLongLink(url.OriginalURL)
	}

	updateTime := u.timer.Now()
	url.UpdatedAt = &updateTime

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

func (u *UpdaterPersist) isURLRelated(shortlink string, user entity.User) bool {
	userShortLinks, err := u.userURLRelationRepo.FindAliasesByUser(user)
	if err != nil {
		return false
	}

	if len(userShortLinks) == 0 {
		return false
	}

	idx := sort.Search(len(userShortLinks), func(i int) bool {
		return userShortLinks[i] == shortlink
	})

	// sort.Search uses a binary search to find the index of the first
	// match and returns length of the slice if no match is found.
	return idx != len(userShortLinks)
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

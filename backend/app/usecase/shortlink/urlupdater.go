package shortlink

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
	UpdateURL(oldAlias string, update entity.ShortLink, user entity.User) (entity.ShortLink, error)
}

// UpdaterPersist persists the given changes to a short link in the database store.
type UpdaterPersist struct {
	urlRepo             repository.ShortLink
	userURLRelationRepo repository.UserShortLink
	keyGen              keygen.KeyGenerator
	longLinkValidator   validator.LongLink
	aliasValidator      validator.CustomAlias
	timer               timer.Timer
	riskDetector        risk.Detector
}

// UpdateURL persists mutations for a given short link in the repository.
func (u UpdaterPersist) UpdateURL(
	oldAlias string,
	update entity.ShortLink,
	user entity.User,
) (entity.ShortLink, error) {
	if !u.isURLRelated(oldAlias, user) {
		return entity.ShortLink{}, errors.New("short link not found")
	}

	url, err := u.urlRepo.GetShortLinkByAlias(oldAlias)
	if err != nil {
		return entity.ShortLink{}, err
	}

	url = u.updateAlias(url, update)
	url = u.updateLongLink(url, update)

	if !u.aliasValidator.IsValid(&url.Alias) {
		return entity.ShortLink{}, ErrInvalidCustomAlias(url.Alias)
	}

	if !u.longLinkValidator.IsValid(&url.LongLink) {
		return entity.ShortLink{}, ErrInvalidLongLink(url.LongLink)
	}

	if u.riskDetector.IsURLMalicious(url.LongLink) {
		return entity.ShortLink{}, ErrMaliciousLongLink(url.LongLink)
	}

	updateTime := u.timer.Now()
	url.UpdatedAt = &updateTime

	return u.urlRepo.UpdateShortLink(oldAlias, url)
}

func (u UpdaterPersist) updateAlias(url, update entity.ShortLink) entity.ShortLink {
	newAlias := update.Alias
	if newAlias != "" {
		url.Alias = newAlias
	}

	return url
}

func (u *UpdaterPersist) updateLongLink(url, update entity.ShortLink) entity.ShortLink {
	newLongLink := update.LongLink
	if newLongLink != "" {
		url.LongLink = newLongLink
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
	urlRepo repository.ShortLink,
	userURLRelationRepo repository.UserShortLink,
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

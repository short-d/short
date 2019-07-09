package usecase

import (
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/repo"

	"github.com/pkg/errors"
)

type UrlRetriever struct {
	urlRepo repo.Url
}

func NewUrlRetriever(urlRepo repo.Url) UrlRetriever {
	return UrlRetriever{
		urlRepo: urlRepo,
	}
}

func (u UrlRetriever) GetUrlAfter(alias string, expiringAt time.Time) (entity.Url, error) {
	url, err := u.urlRepo.GetByAlias(alias)

	if err != nil {
		return entity.Url{}, err
	}

	if url.ExpireAt != nil && expiringAt.After(*url.ExpireAt) {
		return entity.Url{}, errors.Errorf("url expired (alias=%s,expiringAt=%v)", alias, expiringAt)
	}

	return url, nil
}

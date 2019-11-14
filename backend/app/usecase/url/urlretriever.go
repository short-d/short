package url

import (
	"errors"
	"fmt"
	"short/app/entity"
	"short/app/usecase/repo"
	"time"
)

var _ Retriever = (*RetrieverPersist)(nil)

type Retriever interface {
	GetURL(alias string, expiringAt *time.Time) (entity.URL, error)
}

type RetrieverPersist struct {
	urlRepo repo.URL
}

func (r RetrieverPersist) GetURL(alias string, expiringAt *time.Time) (entity.URL, error) {
	if expiringAt == nil {
		return r.getURL(alias)
	}
	return r.getURLExpireAfter(alias, *expiringAt)
}

func (r RetrieverPersist) getURLExpireAfter(alias string, expiringAt time.Time) (entity.URL, error) {
	url, err := r.getURL(alias)
	if err != nil {
		return entity.URL{}, err
	}

	if url.ExpireAt == nil {
		return url, nil
	}

	if expiringAt.After(*url.ExpireAt) {
		return entity.URL{}, errors.New(fmt.Sprintf("url expired (alias=%s,expiringAt=%v)", alias, expiringAt))
	}

	return url, nil
}

func (r RetrieverPersist) getURL(alias string) (entity.URL, error) {
	url, err := r.urlRepo.GetByAlias(alias)
	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func NewRetrieverPersist(urlRepo repo.URL) RetrieverPersist {
	return RetrieverPersist{
		urlRepo: urlRepo,
	}
}

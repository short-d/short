package url

import (
	"fmt"
	"short/app/entity"
	"short/app/usecase/repository"
	"time"
)

var _ Retriever = (*RetrieverPersist)(nil)

// Retriever represents URL retriever
type Retriever interface {
	GetURL(alias string, expiringAt *time.Time) (entity.URL, error)
}

// RetrieverPersist represents URL retriever that fetches URL from persistent
// storage, such as database
type RetrieverPersist struct {
	urlRepo repository.URL
}

// GetURL retrieves URL from persistent storage given alias
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
		return entity.URL{}, fmt.Errorf("url expired (alias=%s,expiringAt=%v)", alias, expiringAt)
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

// NewRetrieverPersist creates persistent URL retriever
func NewRetrieverPersist(urlRepo repository.URL) RetrieverPersist {
	return RetrieverPersist{
		urlRepo: urlRepo,
	}
}

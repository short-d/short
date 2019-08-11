package url

import (
	"short/app/adapter/repo"
	"short/app/entity"
	"short/fw"
	"time"

	"github.com/pkg/errors"
)

type Retriever interface {
	GetAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error)
	Get(trace fw.Trace, alias string) (entity.Url, error)
}

type RetrieverPersist struct {
	urlRepo repo.Url
}

func (u RetrieverPersist) GetAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error) {
	trace1 := trace.Next("Get")
	url, err := u.Get(trace1, alias)
	trace1.End()

	if err != nil {
		return entity.Url{}, err
	}

	if url.ExpireAt == nil {
		return url, nil
	}

	if expiringAt.After(*url.ExpireAt) {
		return entity.Url{}, errors.Errorf("url expired (alias=%s,expiringAt=%v)", alias, expiringAt)
	}

	return url, nil
}

func (u RetrieverPersist) Get(trace fw.Trace, alias string) (entity.Url, error) {
	trace1 := trace.Next("GetByAlias")
	url, err := u.urlRepo.GetByAlias(alias)
	trace1.End()

	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

func NewRetrieverPersist(urlRepo repo.Url) Retriever {
	return RetrieverPersist{
		urlRepo: urlRepo,
	}
}

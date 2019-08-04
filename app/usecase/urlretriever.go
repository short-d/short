package usecase

import (
	"short/app/entity"
	"short/app/repo"
	"short/fw"
	"time"

	"github.com/pkg/errors"
)

type UrlRetriever interface {
	GetUrlAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error)
	GetUrl(trace fw.Trace, alias string) (entity.Url, error)
}

type UrlRetrieverPersist struct {
	urlRepo repo.Url
}

func NewUrlRetrieverPersist(urlRepo repo.Url) UrlRetriever {
	return UrlRetrieverPersist{
		urlRepo: urlRepo,
	}
}

func (u UrlRetrieverPersist) GetUrlAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error) {
	trace1 := trace.Next("GetUrl")
	url, err := u.GetUrl(trace1, alias)
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

func (u UrlRetrieverPersist) GetUrl(trace fw.Trace, alias string) (entity.Url, error) {
	trace1 := trace.Next("GetByAlias")
	url, err := u.urlRepo.GetByAlias(alias)
	trace1.End()

	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

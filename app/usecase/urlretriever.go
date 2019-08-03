package usecase

import (
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/repo"
	"tinyURL/fw"

	"github.com/pkg/errors"
)

type UrlRetriever interface {
	GetUrlAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error)
	GetUrl(trace fw.Trace, alias string) (entity.Url, error)
}

type UrlRetrieverRepo struct {
	urlRepo repo.Url
}

func NewUrlRetrieverRepo(urlRepo repo.Url) UrlRetriever {
	return UrlRetrieverRepo{
		urlRepo: urlRepo,
	}
}

func (u UrlRetrieverRepo) GetUrlAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error) {
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

func (u UrlRetrieverRepo) GetUrl(trace fw.Trace, alias string) (entity.Url, error) {
	trace1 := trace.Next("GetByAlias")
	url, err := u.urlRepo.GetByAlias(alias)
	trace1.End()

	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

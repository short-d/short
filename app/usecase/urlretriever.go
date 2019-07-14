package usecase

import (
	"time"
	"tinyURL/app/entity"
	"tinyURL/app/repo"
	"tinyURL/fw"

	"github.com/pkg/errors"
)

type UrlRetriever interface {
	GetUrlAfter(alias string, expiringAt time.Time) (entity.Url, error)
	GetUrl(alias string) (entity.Url, error)
}

type UrlRetrieverRepo struct {
	tracer  fw.Tracer
	urlRepo repo.Url
}

func NewUrlRetrieverRepo(tracer fw.Tracer, urlRepo repo.Url) UrlRetriever {
	return UrlRetrieverRepo{
		tracer:  tracer,
		urlRepo: urlRepo,
	}
}

func (u UrlRetrieverRepo) GetUrlAfter(alias string, expiringAt time.Time) (entity.Url, error) {
	finish := u.tracer.Begin()
	url, err := u.GetUrl(alias)
	finish("usecase.UrlRetriever.GetUrl")

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

func (u UrlRetrieverRepo) GetUrl(alias string) (entity.Url, error) {
	finish := u.tracer.Begin()
	url, err := u.urlRepo.GetByAlias(alias)
	finish("repo.Url.GetByAlias")

	if err != nil {
		return entity.Url{}, err
	}

	return url, nil
}

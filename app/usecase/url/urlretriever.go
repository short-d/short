package url

import (
	"errors"
	"fmt"
	"short/app/entity"
	"short/app/usecase/repo"
	"time"

	"github.com/byliuyang/app/fw"
)

var _ Retriever = (*RetrieverPersist)(nil)

type Retriever interface {
	GetAfter(trace fw.Segment, alias string, expiringAt time.Time) (entity.URL, error)
	Get(trace fw.Segment, alias string) (entity.URL, error)
	GetList(trace fw.Segment, email string) ([]entity.URL, error)
}

type RetrieverPersist struct {
	urlRepo repo.URL
}

func (u RetrieverPersist) GetAfter(trace fw.Segment, alias string, expiringAt time.Time) (entity.URL, error) {
	trace1 := trace.Next("Get")
	url, err := u.Get(trace1, alias)
	trace1.End()

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

func (u RetrieverPersist) Get(trace fw.Segment, alias string) (entity.URL, error) {
	trace1 := trace.Next("GetByAlias")
	url, err := u.urlRepo.GetByAlias(alias)
	trace1.End()

	if err != nil {
		return entity.URL{}, err
	}

	return url, nil
}

func (u RetrieverPersist) GetList(trace fw.Segment, email string) ([]entity.URL, error) {
	trace1 := trace.Next("GetByUser")
	urls, err := u.urlRepo.GetByUser(email)
	trace1.End()

	if err != nil {
		return nil, err
	}

	return urls, nil
}

func NewRetrieverPersist(urlRepo repo.URL) RetrieverPersist {
	return RetrieverPersist{
		urlRepo: urlRepo,
	}
}

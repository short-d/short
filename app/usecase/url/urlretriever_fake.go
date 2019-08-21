package url

import (
	"errors"
	"short/app/entity"
	"short/fw"
	"time"
)

var _ Retriever = (*FakeRetriever)(nil)

type FakeRetriever struct {
	urls map[string]entity.Url
}

func (f FakeRetriever) GetAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.Url, error) {
	url, ok := f.urls[alias]
	if !ok {
		return entity.Url{}, errors.New("url not found")
	}

	if url.ExpireAt == nil {
		return url, nil
	}
	if expiringAt.After(*url.ExpireAt) {
		return entity.Url{}, errors.New("url expired")
	}

	return url, nil
}

func (f FakeRetriever) Get(trace fw.Trace, alias string) (entity.Url, error) {
	url, ok := f.urls[alias]
	if !ok {
		return entity.Url{}, errors.New("url not found")
	}

	return url, nil
}

func NewRetrieverFake(urls map[string]entity.Url) FakeRetriever {
	return FakeRetriever{
		urls: urls,
	}
}

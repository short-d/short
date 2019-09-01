package url

import (
	"errors"
	"short/app/entity"
	"time"

	"short/fw"
)

var _ Retriever = (*FakeRetriever)(nil)

type FakeRetriever struct {
	urls map[string]entity.URL
}

func (f FakeRetriever) GetAfter(trace fw.Trace, alias string, expiringAt time.Time) (entity.URL, error) {
	url, ok := f.urls[alias]
	if !ok {
		return entity.URL{}, errors.New("url not found")
	}

	if url.ExpireAt == nil {
		return url, nil
	}
	if expiringAt.After(*url.ExpireAt) {
		return entity.URL{}, errors.New("url expired")
	}

	return url, nil
}

func (f FakeRetriever) Get(trace fw.Trace, alias string) (entity.URL, error) {
	url, ok := f.urls[alias]
	if !ok {
		return entity.URL{}, errors.New("url not found")
	}

	return url, nil
}

func NewRetrieverFake(urls map[string]entity.URL) FakeRetriever {
	return FakeRetriever{
		urls: urls,
	}
}

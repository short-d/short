package usecase

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/usecase/url"
)

type UseCase interface {
	ViewLongLink(
		alias string,
		aliasNotFoundCallback func(),
		viewLongLinkCallback func(longLink string),
	)
}

var _ UseCase = (*Short)(nil)

type Short struct {
	logger       fw.Logger
	urlRetriever url.Retriever
	timer        fw.Timer
}

func (s Short) ViewLongLink(
	alias string,
	aliasNotFoundCallback func(),
	viewLongLinkCallback func(longLink string),
) {
	now := s.timer.Now()
	u, err := s.urlRetriever.GetURL(alias, &now)
	if err != nil {
		s.logger.Error(err)
		aliasNotFoundCallback()
		return
	}
	viewLongLinkCallback(u.OriginalURL)
}

func NewShort(
	logger fw.Logger,
	urlRetriever url.Retriever,
	timer fw.Timer,
) Short {
	return Short{
		logger:       logger,
		urlRetriever: urlRetriever,
		timer:        timer,
	}
}

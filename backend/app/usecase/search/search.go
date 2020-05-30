package search

import (
	"errors"
	"strings"
	"time"

	"github.com/short-d/short/backend/app/usecase/search/order"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

func (s Search) Search(query Query, filter Filter) (Result, error) {
	resultCh := make(chan Result)
	defer close(resultCh)

	var orders []order.Order
	for _, orderBy := range filter.orders {
		orders = append(orders, order.NewOrder(orderBy))
	}

	for i := range filter.resources {
		go func() {
			result, err := s.searchResource(filter.resources[i], orders[i], query, filter)
			if err != nil {
				resultCh <- nil
				return
			}
			resultCh <- result
		}()
	}

	timeout := time.After(s.timeout)
	var results []Result
	for i := 0; i < len(filter.resources); i++ {
		select {
		case result := <-resultCh:
			results = append(results, result)
		case <-timeout:
			return mergeResults(results), nil
		}
	}

	return mergeResults(results), nil
}

func mergeResults(results []Result) Result {
	var mergedResult Result

	for _, result := range results {
		mergedResult.shortLinks = append(mergedResult.shortLinks, result.shortLinks...)
		mergedResult.users = append(mergedResult.users, result.users...)
	}

	return mergedResult
}

func (s Search) searchShortLink(query Query, orderBy order.Order, filter Filter) (Result, error) {
	if query.user == nil {
		return Result{}, errors.New("user not provided")
	}

	aliases, err := s.userShortLinkRepo.FindAliasesByUser(*query.user)
	if err != nil {
		return Result{}, err
	}
	shortLinks, err := s.shortLinkRepo.GetShortLinksByAliases(aliases)
	if err != nil {
		return Result{}, err
	}

	var matchedByAlias []entity.ShortLink
	for _, shortLink := range shortLinks {
		if strings.Contains(shortLink.Alias, query.query) {
			matchedByAlias = append(matchedByAlias, shortLink)
		}
	}
	matchedByAlias = orderBy.ArrangeShortLinks(matchedByAlias)

	var matchedByLongLink []entity.ShortLink
	for _, shortLink := range shortLinks {
		if strings.Contains(shortLink.LongLink, query.query) {
			matchedByLongLink = append(matchedByLongLink, shortLink)
		}
	}
	matchedByLongLink = orderBy.ArrangeShortLinks(matchedByLongLink)

	mergedShortLinks := append(matchedByAlias, matchedByLongLink...)

	return Result{
		shortLinks: mergedShortLinks,
		users:      nil,
	}, nil
}

func (s Search) searchUser(query Query, orderBy order.Order, filter Filter) (Result, error) {
	return Result{}, nil
}

func (s Search) searchResource(resource Resource, orderBy order.Order, query Query, filter Filter) (Result, error) {
	switch resource {
	case ShortLink:
		return s.searchShortLink(query, orderBy, filter)
	case User:
		return s.searchUser(query, orderBy, filter)
	}
	return Result{}, errors.New("error")
}

func NewSearch(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	timeout time.Duration,
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		timeout:           timeout,
	}
}

package search

import (
	"errors"
	"strings"
	"time"

	"github.com/short-d/short/backend/app/usecase/search/order"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

// Search represents the search handler of short links and users
// from a persistent storage
type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

// Result represents the result of a search
type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

// Search searches the short links and users for given query and filter
func (s Search) Search(query Query, filter Filter) (Result, error) {
	resultCh := make(chan Result)
	defer close(resultCh)

	var orders []order.Order
	// register orders
	for _, orderBy := range filter.Orders {
		orders = append(orders, order.NewOrder(orderBy))
	}

	// search resources in concurrently
	for i := range filter.Resources {
		go func() {
			result, err := s.searchResource(filter.Resources[i], orders[i], query, filter)
			if err != nil {
				resultCh <- nil
				return
			}
			resultCh <- result
		}()
	}

	timeout := time.After(s.timeout)
	var results []Result
	for i := 0; i < len(filter.Resources); i++ {
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

func (s Search) searchResource(resource Resource, orderBy order.Order, query Query, filter Filter) (Result, error) {
	switch resource {
	case ShortLink:
		return s.searchShortLink(query, orderBy, filter)
	case User:
		return s.searchUser(query, orderBy, filter)
	default:
		return Result{}, errors.New("unknown resource")
	}
}

func (s Search) searchShortLink(query Query, orderBy order.Order, filter Filter) (Result, error) {
	if query.User == nil {
		return Result{}, errors.New("user not provided")
	}

	aliases, err := s.userShortLinkRepo.FindAliasesByUser(*query.User)
	if err != nil {
		return Result{}, err
	}
	shortLinks, err := s.shortLinkRepo.GetShortLinksByAliases(aliases)
	if err != nil {
		return Result{}, err
	}

	var matchedAliasByAnd, matchedAliasByOr []entity.ShortLink
	for _, shortLink := range shortLinks {
		// check if string is contained by "and" operator and by "or" operator
		and, or := stringContains(shortLink.Alias, strings.Split(query.Query, " "))
		if and {
			matchedAliasByAnd = append(matchedAliasByAnd, shortLink)
		} else if or {
			matchedAliasByOr = append(matchedAliasByOr, shortLink)
		}
	}

	// sort short links
	matchedAliasByAnd = orderBy.ArrangeShortLinks(matchedAliasByAnd)
	matchedAliasByOr = orderBy.ArrangeShortLinks(matchedAliasByOr)

	var matchedLongLinkByAnd, matchedLongLinkByOr []entity.ShortLink
	for _, shortLink := range shortLinks {
		and, or := stringContains(shortLink.LongLink, strings.Split(query.Query, " "))
		if and {
			matchedLongLinkByAnd = append(matchedLongLinkByAnd, shortLink)
		} else if or {
			matchedLongLinkByOr = append(matchedLongLinkByOr, shortLink)
		}
	}

	// sort short links
	matchedLongLinkByAnd = orderBy.ArrangeShortLinks(matchedLongLinkByAnd)
	matchedLongLinkByOr = orderBy.ArrangeShortLinks(matchedLongLinkByOr)

	// merge all the short links
	mergedShortLinks := append(matchedAliasByAnd, matchedAliasByOr...)
	mergedShortLinks = append(mergedShortLinks, matchedLongLinkByAnd...)
	mergedShortLinks = append(mergedShortLinks, matchedLongLinkByOr...)

	return Result{
		shortLinks: mergedShortLinks,
		users:      nil,
	}, nil
}

func stringContains(s string, words []string) (bool, bool) {
	and := true
	or := false
	for _, word := range words {
		if !and && or {
			return and, or
		}
		if strings.Contains(s, word) {
			or = true
		} else {
			and = false
		}
	}
	return and, or
}

func (s Search) searchUser(query Query, orderBy order.Order, filter Filter) (Result, error) {
	return Result{}, nil
}

// NewSearch creates Search
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

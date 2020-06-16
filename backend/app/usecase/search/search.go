package search

import (
	"errors"
	"strings"
	"time"

	"github.com/short-d/app/fw/logger"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/matcher"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

// Search finds different types of resources matching certain criteria and sort them based on predefined orders.
type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
	logger            logger.Logger
}

// Result represents the result of a search query.
type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

// Search finds resources based on specified criteria.
func (s Search) Search(query Query, filter Filter) (Result, error) {
	resultCh := make(chan Result)
	defer close(resultCh)

	orders := toOrders(filter.Orders)

	for i := range filter.Resources {
		go func() {
			result, err := s.searchResource(filter.Resources[i], orders[i], query, filter)
			if err != nil {
				// TODO(issue#865): Handle errors of Search API
				s.logger.Error(err)
				resultCh <- Result{}
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

// TODO(issue#866): Simplify searchShortLink function
func (s Search) searchShortLink(query Query, orderBy order.Order, filter Filter) (Result, error) {
	if query.User == nil {
		s.logger.Error(errors.New("user not provided"))
		return Result{}, nil
	}

	shortLinks, err := s.getShortLinkByUser(*query.User)
	if err != nil {
		return Result{}, err
	}

	var matchedAliasByAll, matchedAliasByAny, matchedLongLinkByAll, matchedLongLinkByAny []entity.ShortLink
	keywords := getKeywords(query.Query)
	for _, shortLink := range shortLinks {
		if matcher.ContainsAll(keywords, shortLink.Alias) {
			matchedAliasByAll = append(matchedAliasByAll, shortLink)
		} else if matcher.ContainsAny(keywords, shortLink.Alias) {
			matchedAliasByAny = append(matchedAliasByAny, shortLink)
		} else if matcher.ContainsAll(keywords, shortLink.LongLink) {
			matchedLongLinkByAll = append(matchedLongLinkByAll, shortLink)
		} else if matcher.ContainsAny(keywords, shortLink.LongLink) {
			matchedLongLinkByAny = append(matchedLongLinkByAny, shortLink)
		}
	}

	sortShortLinks(orderBy, matchedAliasByAll, matchedAliasByAny, matchedLongLinkByAll, matchedLongLinkByAny)

	mergedShortLinks := mergeShortLinks(
		matchedAliasByAll,
		matchedAliasByAny,
		matchedLongLinkByAll,
		matchedLongLinkByAny,
	)

	filteredShortLinks := filterShortLinks(mergedShortLinks, filter)

	return Result{
		shortLinks: filteredShortLinks,
		users:      nil,
	}, nil
}

func (s Search) searchUser(query Query, orderBy order.Order, filter Filter) (Result, error) {
	return Result{}, nil
}

func (s Search) getShortLinkByUser(user entity.User) ([]entity.ShortLink, error) {
	aliases, err := s.userShortLinkRepo.FindAliasesByUser(user)
	if err != nil {
		return []entity.ShortLink{}, err
	}

	return s.shortLinkRepo.GetShortLinksByAliases(aliases)
}

func getKeywords(query string) []string {
	return strings.Split(query, " ")
}

func filterShortLinks(shortLinks []entity.ShortLink, filter Filter) []entity.ShortLink {
	if len(shortLinks) > filter.MaxResults {
		shortLinks = shortLinks[:filter.MaxResults]
	}
	return shortLinks
}

func sortShortLinks(orderBy order.Order, shortLinkLists ...[]entity.ShortLink) {
	for _, shortLinks := range shortLinkLists {
		orderBy.ArrangeShortLinks(shortLinks)
	}
}

func mergeShortLinks(shortLinkLists ...[]entity.ShortLink) []entity.ShortLink {
	var merged []entity.ShortLink
	for _, shortLinks := range shortLinkLists {
		merged = append(merged, shortLinks...)
	}
	return merged
}

func toOrders(ordersBy []order.By) []order.Order {
	var orders []order.Order
	for _, orderBy := range ordersBy {
		orders = append(orders, order.NewOrder(orderBy))
	}
	return orders
}

func mergeResults(results []Result) Result {
	var mergedResult Result

	for _, result := range results {
		mergedResult.shortLinks = append(mergedResult.shortLinks, result.shortLinks...)
		mergedResult.users = append(mergedResult.users, result.users...)
	}

	return mergedResult
}

// NewSearch creates Search
func NewSearch(
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	timeout time.Duration,
	logger logger.Logger,
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		timeout:           timeout,
		logger:            logger,
	}
}

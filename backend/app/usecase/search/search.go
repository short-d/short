package search

import (
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
	logger            logger.Logger
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

// Result represents the result of a search query.
type Result struct {
	ShortLinks []entity.ShortLink
	Users      []entity.User
}

// ErrUnknownResource represents unknown search resource error.
type ErrUnknownResource struct{}

func (e ErrUnknownResource) Error() string {
	return "unknown resource"
}

// ErrUserNotProvided represents user not provided for search query.
type ErrUserNotProvided struct{}

func (e ErrUserNotProvided) Error() string {
	return "user not provided"
}

// Search finds resources based on specified criteria.
func (s Search) Search(query Query, filter Filter) (Result, error) {
	resultCh := make(chan Result)
	errCh := make(chan error)
	defer close(resultCh)
	defer close(errCh)

	orders := toOrders(filter.orders)

	for i := range filter.resources {
		i := i
		go func() {
			result, err := s.searchResource(filter.resources[i], orders[i], query, filter)
			if err != nil {
				s.logger.Error(err)
				resultCh <- Result{}
				errCh <- err
				return
			}
			resultCh <- result
			errCh <- nil
		}()
	}

	timeout := time.After(s.timeout)
	var results []Result
	var resultErr error
	for i := 0; i < len(filter.resources); i++ {
		select {
		case result := <-resultCh:
			results = append(results, result)
		case <-timeout:
			return mergeResults(results), nil
		}
		select {
		case err := <-errCh:
			// Only return the first error encountered
			if resultErr == nil {
				resultErr = err
			}
		}
	}

	return mergeResults(results), resultErr
}

func (s Search) searchResource(resource Resource, orderBy order.Order, query Query, filter Filter) (Result, error) {
	switch resource {
	case ShortLink:
		return s.searchShortLink(query, orderBy, filter)
	case User:
		return s.searchUser(query, orderBy, filter)
	default:
		return Result{}, ErrUnknownResource{}
	}
}

// TODO(issue#866): Simplify searchShortLink function
func (s Search) searchShortLink(query Query, orderBy order.Order, filter Filter) (Result, error) {
	if query.User == nil {
		return Result{}, ErrUserNotProvided{}
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
			continue
		}
		if matcher.ContainsAny(keywords, shortLink.Alias) {
			matchedAliasByAny = append(matchedAliasByAny, shortLink)
			continue
		}
		if matcher.ContainsAll(keywords, shortLink.LongLink) {
			matchedLongLinkByAll = append(matchedLongLinkByAll, shortLink)
			continue
		}
		if matcher.ContainsAny(keywords, shortLink.LongLink) {
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
		ShortLinks: filteredShortLinks,
		Users:      nil,
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
	if filter.maxResults == 0 {
		return shortLinks
	}

	if len(shortLinks) > filter.maxResults {
		shortLinks = shortLinks[:filter.maxResults]
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
		mergedResult.ShortLinks = append(mergedResult.ShortLinks, result.ShortLinks...)
		mergedResult.Users = append(mergedResult.Users, result.Users...)
	}

	return mergedResult
}

// NewSearch creates Search
func NewSearch(
	logger logger.Logger,
	shortLinkRepo repository.ShortLink,
	userShortLinkRepo repository.UserShortLink,
	timeout time.Duration,
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		timeout:           timeout,
		logger:            logger,
	}
}

package search

import (
	"errors"
	"strings"
	"time"

	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

// Search finds different types of resources matching certain criteria and sort them based on predefined orders.
type Search struct {
	shortLinkRepo     repository.ShortLink
	userShortLinkRepo repository.UserShortLink
	timeout           time.Duration
}

// Result represents the result of a search query.
type Result struct {
	shortLinks []entity.ShortLink
	users      []entity.User
}

// Search finds resources based on specified criteria.
func (s Search) Search(query Query, filter Filter) (Result, error) {
	var finalErr error
	resultCh := make(chan Result)
	defer close(resultCh)

	orders := toOrders(filter.Orders)

	for i := range filter.Resources {
		go func() {
			result, err := s.searchResource(filter.Resources[i], orders[i], query, filter)
			if err != nil {
				finalErr = err
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
			return mergeResults(results), finalErr
		}
	}

	return mergeResults(results), finalErr
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
	if len(query.Query) == 0 {
		return Result{}, errors.New("query not provided")
	}

	shortLinks, err := s.getShortLinkByUser(*query.User)
	if err != nil {
		return Result{}, err
	}

	matchedShortLinks := matchShortLinks(shortLinks, query, orderBy)

	filteredShortLinks := filterShortLinks(matchedShortLinks, filter)

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

func matchShortLinks(shortLinks []entity.ShortLink, query Query, orderBy order.Order) []entity.ShortLink {
	var matchedAliasByAnd, matchedAliasByOr, matchedLongLinkByAnd, matchedLongLinkByOr []entity.ShortLink
	keywords := getKeywordsFromQuery(query.Query)
	for _, shortLink := range shortLinks {
		// check if query is contained in alias by "and" operator and by "or" operator
		aliasByAnd := containsAllWords(shortLink.Alias, keywords)
		aliasByOr := containsSomeWords(shortLink.Alias, keywords)
		if aliasByAnd {
			matchedAliasByAnd = append(matchedAliasByAnd, shortLink)
			continue
		} else if aliasByOr {
			matchedAliasByOr = append(matchedAliasByOr, shortLink)
			continue
		}

		// check if query is contained in long link by "and" operator and by "or" operator
		longLinkByAnd := containsAllWords(shortLink.LongLink, keywords)
		longLinkByOr := containsSomeWords(shortLink.LongLink, keywords)
		if longLinkByAnd {
			matchedLongLinkByAnd = append(matchedLongLinkByAnd, shortLink)
		} else if longLinkByOr {
			matchedLongLinkByOr = append(matchedLongLinkByOr, shortLink)
		}
	}

	// sort short links
	sortShortLinks(orderBy, matchedAliasByAnd, matchedAliasByOr, matchedLongLinkByAnd, matchedLongLinkByOr)

	// merge all the short links
	return mergeShortLinks(matchedAliasByAnd, matchedAliasByOr, matchedLongLinkByAnd, matchedLongLinkByOr)
}

func getKeywordsFromQuery(query string) []string {
	return strings.Split(query, " ")
}

func containsAllWords(s string, words []string) bool {
	for _, word := range words {
		if !strings.Contains(s, word) {
			return false
		}
	}
	return true
}

func containsSomeWords(s string, words []string) bool {
	for _, word := range words {
		if strings.Contains(s, word) {
			return true
		}
	}
	return false
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
) Search {
	return Search{
		shortLinkRepo:     shortLinkRepo,
		userShortLinkRepo: userShortLinkRepo,
		timeout:           timeout,
	}
}

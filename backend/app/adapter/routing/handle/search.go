package handle

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/search"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

// SearchRequest represents the request received from Search API.
type SearchRequest struct {
	query  search.Query
	filter search.Filter
}

// Filter represents the filter field received from Search API.
type Filter struct {
	maxResults int
	resources  []search.Resource
	orders     []order.By
}

// Search fetches resources under certain criteria.
func Search(
	searcher search.Search,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		user := getUser(r, authenticator)
		if user == nil {
			w.Write([]byte("user not authenticated"))
			return
		}

		searchBody, err := parseSearchBody(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		searchBody.setUser(user)

		w.Write([]byte("request is read"))
	}
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (s *SearchRequest) UnmarshalJSON(data []byte) error {
	temp := struct {
		Query  string `json:"query"`
		Filter Filter `json:"filter"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.setQuery(temp.Query)
	err := s.setFilter(temp.Filter.maxResults, temp.Filter.resources, temp.Filter.orders)

	return err
}

func (s *SearchRequest) setQuery(query string) {
	s.query.Query = query
}

func (s *SearchRequest) setFilter(maxResults int, resources []search.Resource, orders []order.By) error {
	filter, err := search.NewFilter(maxResults, resources, orders)
	if err != nil {
		return err
	}

	s.filter = filter
	return nil
}

func (s *SearchRequest) setUser(user *entity.User) {
	s.query.User = user
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (f *Filter) UnmarshalJSON(data []byte) error {
	temp := struct {
		MaxResults int      `json:"max_results"`
		Resources  []string `json:"resources"`
		Orders     []string `json:"orders"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	resources, err := parseResources(temp.Resources)
	if err != nil {
		return err
	}
	orders, err := parseOrders(temp.Orders)
	if err != nil {
		return err
	}

	f.maxResults = temp.MaxResults
	f.resources = resources
	f.orders = orders

	return nil
}

func parseSearchBody(r *http.Request) (*SearchRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var searchRequest SearchRequest
	err = json.Unmarshal(b, &searchRequest)
	return &searchRequest, err
}

func parseResources(ss []string) ([]search.Resource, error) {
	var resources []search.Resource
	for _, s := range ss {
		if resource, err := stringToResource(s); err == nil {
			resources = append(resources, resource)
		} else {
			return resources, err
		}
	}
	return resources, nil
}

func stringToResource(s string) (search.Resource, error) {
	switch s {
	case "short_link":
		return search.ShortLink, nil
	case "user":
		return search.User, nil
	default:
		return search.Resource(0), errors.New("unknown resource")
	}
}

func parseOrders(ss []string) ([]order.By, error) {
	var ordersBy []order.By
	for _, s := range ss {
		if orderBy, err := stringToOrder(s); err == nil {
			ordersBy = append(ordersBy, orderBy)
		} else {
			return ordersBy, err
		}
	}
	return ordersBy, nil
}

func stringToOrder(s string) (order.By, error) {
	switch s {
	case "by_created_time_asc":
		return order.ByCreatedTimeASC, nil
	default:
		return order.By(0), errors.New("unknown order")
	}
}

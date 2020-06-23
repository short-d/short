package handle

import (
	"encoding/json"
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
		Filter struct {
			MaxResults int      `json:"max_results"`
			Resources  []string `json:"resources"`
			Orders     []string `json:"orders"`
		} `json:"filter"`
	}{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	s.setQuery(temp.Query)

	resources := parseResources(temp.Filter.Resources)
	orders := parseOrders(temp.Filter.Orders)
	err := s.setFilter(temp.Filter.MaxResults, resources, orders)

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

func parseResources(data []string) []search.Resource {
	var resources []search.Resource
	for _, resource := range data {
		if resource == "short_link" {
			resources = append(resources, search.ShortLink)
		} else if resource == "user" {
			resources = append(resources, search.User)
		}
	}
	return resources
}

func parseOrders(data []string) []order.By {
	var orders []order.By
	for _, o := range data {
		if o == "by_created_time_asc" {
			orders = append(orders, order.ByCreatedTimeASC)
		}
	}
	return orders
}

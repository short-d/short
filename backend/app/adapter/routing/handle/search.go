package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/short-d/short/backend/app/usecase/search/order"

	"github.com/short-d/short/backend/app/usecase/authenticator"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/usecase/search"
)

// Search fetches resources under certain criterias.
func Search(
	searcher search.Search,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		//user := getUser(r, authenticator)
		//if user == nil {
		//	w.Write([]byte("user not logged in"))
		//}

		searchRequest, err := readSearchRequest(r)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		_ = search.Query{
			Query: searchRequest.Query,
			User:  nil,
		}
		_, err = getFilter(searchRequest)
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		w.Write([]byte("request is read"))

	}
}

// SearchRequest represents the JSON request received from Search API.
type SearchRequest struct {
	Query  string `json:"query"`
	Filter Filter `json:"filter"`
}

// Filter represents the filter JSON field received from Search API.
type Filter struct {
	MaxResults int      `json:"max_results"`
	Resources  []string `json:"resources"`
	Orders     []string `json:"orders"`
}

// SearchResponse represents the JSON response received from Search API.
type SearchResponse struct {
}

func readSearchRequest(r *http.Request) (*SearchRequest, error) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}

	var searchRequest SearchRequest
	err = json.Unmarshal(b, &searchRequest)
	return &searchRequest, err
}

func getFilter(request *SearchRequest) (search.Filter, error) {
	return search.NewFilter(request.Filter.MaxResults, getResources(request), getOrders(request))
}

func getResources(request *SearchRequest) []search.Resource {
	var resources []search.Resource
	for _, resource := range request.Filter.Resources {
		if resource == "short_link" {
			resources = append(resources, search.ShortLink)
		} else if resource == "user" {
			resources = append(resources, search.User)
		}
	}
	return resources
}

func getOrders(request *SearchRequest) []order.By {
	var orders []order.By
	for _, o := range request.Filter.Orders {
		if o == "time asc" {
			orders = append(orders, order.ByCreatedTimeASC)
		}
	}
	return orders
}

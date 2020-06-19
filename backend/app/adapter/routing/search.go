package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/short-d/short/backend/app/usecase/search/order"

	"github.com/short-d/short/backend/app/usecase/search"
)

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

func getFilter(request *SearchRequest) search.Filter {
	var filter search.Filter

	filter.MaxResults = request.Filter.MaxResults
	filter.Resources = getResources(request)
	filter.Orders = getOrders(request)

	return filter
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

func writeSearchResponse(w http.ResponseWriter, response *SearchResponse) error {
	output, err := json.Marshal(response)
	if err != nil {
		return err
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	return nil
}

package handle

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/search"
	"github.com/short-d/short/backend/app/usecase/search/order"
)

var searchResource = map[string]search.Resource{
	"short_link": search.ShortLink,
	"user":       search.User,
}

var searchOrder = map[string]order.By{
	"created_time_asc": order.ByCreatedTimeASC,
}

// SearchRequest represents the request received from Search API.
type SearchRequest struct {
	Query  string `json:"query"`
	Filter Filter `json:"filter"`
}

// Filter represents the filter field received from Search API.
type Filter struct {
	MaxResults int
	Resources  []search.Resource
	Orders     []order.By
}

// Search fetches resources under certain criteria.
func Search(
	search search.Search,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		buf, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var body SearchRequest
		err = json.Unmarshal(buf, &body)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(fmt.Sprintf("%v", body)))
	}
}

// UnmarshalJSON parses json into Filter
func (f *Filter) UnmarshalJSON(data []byte) error {
	buf := struct {
		MaxResults int      `json:"max_results"`
		Resources  []string `json:"resources"`
		Orders     []string `json:"orders"`
	}{}

	if err := json.Unmarshal(data, &buf); err != nil {
		return err
	}

	f.MaxResults = buf.MaxResults

	for _, resource := range buf.Resources {
		val, ok := searchResource[resource]
		if !ok {
			f.Resources = append(f.Resources, search.Unknown)
			continue
		}
		f.Resources = append(f.Resources, val)
	}

	for _, o := range buf.Orders {
		val, ok := searchOrder[o]
		if !ok {
			f.Orders = append(f.Orders, order.ByUnsorted)
			continue
		}
		f.Orders = append(f.Orders, val)
	}
	return nil
}

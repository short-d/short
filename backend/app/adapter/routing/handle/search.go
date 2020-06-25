package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/entity"
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

// SearchResponse represents the response to the Search API request.
type SearchResponse struct {
	ShortLinks []ShortLink `json:"short_links,omitempty"`
	Users      []User      `json:"users,omitempty"`
}

// ShortLink represents the short_link field of Search API respond.
type ShortLink struct {
	Alias     string     `json:"alias,omitempty"`
	LongLink  string     `json:"long_link,omitempty"`
	ExpireAt  *time.Time `json:"expire_at,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// User represents the user field of Search API respond.
type User struct {
	ID             string     `json:"id,omitempty"`
	Name           string     `json:"name,omitempty"`
	Email          string     `json:"email,omitempty"`
	LastSignedInAt *time.Time `json:"last_signed_in_at,omitempty"`
	CreatedAt      *time.Time `json:"created_at,omitempty"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty"`
}

// Search fetches resources under certain criteria.
func Search(
	instrumentationFactory request.InstrumentationFactory,
	searcher search.Search,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		i := instrumentationFactory.NewHTTP(r)

		buf, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			i.SearchFailed(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var body SearchRequest
		err = json.Unmarshal(buf, &body)
		if err != nil {
			i.SearchFailed(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := getUser(r, authenticator)
		query := search.Query{
			Query: body.Query,
			User:  user,
		}
		filter, err := search.NewFilter(body.Filter.MaxResults, body.Filter.Resources, body.Filter.Orders)
		if err != nil {
			i.SearchFailed(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		results, err := searcher.Search(query, filter)
		if err != nil {
			i.SearchFailed(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := newSearchResponse(results)
		respBody, err := json.Marshal(&response)
		if err != nil {
			i.SearchFailed(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(respBody)
		i.SearchSucceed()
	}
}

// UnmarshalJSON parses json into Filter.
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

func newSearchResponse(result search.Result) SearchResponse {
	shortLinks := make([]ShortLink, len(result.ShortLinks))
	for i := 0; i < len(result.ShortLinks); i++ {
		shortLinks[i] = newShortLink(result.ShortLinks[i])
	}

	users := make([]User, len(result.Users))
	for i := 0; i < len(result.Users); i++ {
		users[i] = newUser(result.Users[i])
	}

	return SearchResponse{
		ShortLinks: shortLinks,
		Users:      users,
	}
}

func newShortLink(shortLink entity.ShortLink) ShortLink {
	return ShortLink{
		Alias:     shortLink.Alias,
		LongLink:  shortLink.LongLink,
		ExpireAt:  shortLink.ExpireAt,
		CreatedAt: shortLink.CreatedAt,
		UpdatedAt: shortLink.UpdatedAt,
	}
}

func newUser(user entity.User) User {
	return User{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		LastSignedInAt: user.LastSignedInAt,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
	}
}

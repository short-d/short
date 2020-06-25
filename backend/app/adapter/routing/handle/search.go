package handle

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/short-d/app/fw/router"
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

// SearchRespond represents the respond to the Search API request.
type SearchRespond search.Result

// ShortLink represents the short_link field of Search API respond.
type ShortLink entity.ShortLink

// User represents the user field of Search API respond.
type User entity.User

// Search fetches resources under certain criteria.
func Search(
	searcher search.Search,
	authenticator authenticator.Authenticator,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		buf, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var body SearchRequest
		err = json.Unmarshal(buf, &body)
		if err != nil {
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		results, err := searcher.Search(query, filter)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		respond := SearchRespond(results)
		marshaled, err := json.Marshal(&respond)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(marshaled)
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

// MarshalJSON formats the search result into json format.
func (s *SearchRespond) MarshalJSON() ([]byte, error) {
	shortLinks := make([]ShortLink, len(s.ShortLinks))
	for i := 0; i < len(s.ShortLinks); i++ {
		shortLinks[i] = ShortLink(s.ShortLinks[i])
	}
	users := make([]User, len(s.Users))
	for i := 0; i < len(s.Users); i++ {
		users[i] = User(s.Users[i])
	}

	buf := struct {
		ShortLink []ShortLink `json:"short_link,omitempty"`
		User      []User      `json:"user,omitempty"`
	}{
		ShortLink: shortLinks,
		User:      users,
	}

	return json.Marshal(&buf)
}

// MarshalJSON formats the short link into json format.
func (s *ShortLink) MarshalJSON() ([]byte, error) {
	buf := struct {
		Alias     string     `json:"alias,omitempty"`
		LongLink  string     `json:"long_link,omitempty"`
		ExpireAt  *time.Time `json:"expire_at,omitempty"`
		CreatedAt *time.Time `json:"created_at,omitempty"`
		UpdatedAt *time.Time `json:"updated_at,omitempty"`
	}{
		s.Alias,
		s.LongLink,
		s.ExpireAt,
		s.CreatedAt,
		s.UpdatedAt,
	}

	return json.Marshal(&buf)
}

// MarshalJSON formats the user into json format.
func (u *User) MarshalJSON() ([]byte, error) {
	buf := struct {
		ID             string     `json:"id,omitempty"`
		Name           string     `json:"name,omitempty"`
		Email          string     `json:"email,omitempty"`
		LastSignedInAt *time.Time `json:"last_signed_in_at,omitempty"`
		CreatedAt      *time.Time `json:"created_at,omitempty"`
		UpdatedAt      *time.Time `json:"updated_at,omitempty"`
	}{
		u.ID,
		u.Name,
		u.Email,
		u.LastSignedInAt,
		u.CreatedAt,
		u.UpdatedAt,
	}

	return json.Marshal(&buf)
}

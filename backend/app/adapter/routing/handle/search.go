package handle

import (
	"encoding/json"
	"errors"
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

var searchResourceString = map[search.Resource]string{
	search.ShortLink: "short_link",
	search.User:      "user",
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

// SearchError represents an error with the Search API request.
type SearchError struct {
	Message string `json:"message"`
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
			emitSearchError(w, err)
			return
		}

		var body SearchRequest
		err = json.Unmarshal(buf, &body)
		if err != nil {
			i.SearchFailed(err)
			emitSearchError(w, err)
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
			emitSearchError(w, err)
			return
		}

		results, err := searcher.Search(query, filter)
		if err != nil {
			i.SearchFailed(err)
			emitSearchError(w, err)
			return
		}

		response := newSearchResponse(results)
		respBody, err := json.Marshal(&response)
		if err != nil {
			i.SearchFailed(err)
			emitSearchError(w, err)
			return
		}

		w.Write(respBody)
		i.SearchSucceed(user, query.Query, body.Filter.resourcesString())
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

func (f *Filter) resourcesString() []string {
	var resources []string
	for _, resource := range f.Resources {
		val, ok := searchResourceString[resource]
		if ok {
			resources = append(resources, val)
		}
	}
	return resources
}

func newSearchResponse(result search.ResourceResult) SearchResponse {
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

func emitSearchError(w http.ResponseWriter, err error) {
	var (
	       code http.StatusInternalServerError
		u search.ErrUserNotProvided
		r search.ErrUnknownResource
	)
	if errors.As(err, &u) {
		code = http.StatusUnauthorized
	}
	if errors.As(err, &r) {
		code = http.StatusNotFound
	}
	errResp, err := json.Marshal(SearchError{
		Message: err.Error(),
	})
	if err != nil {
		return
	}
	http.Error(w, string(errResp), code)
}

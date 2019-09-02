package github

import (
	"fmt"
	"short/app/entity"

	"github.com/byliuyang/app/fw"
)

const githubAPI = "https://api.github.com/graphql"

// API access user's account & repository information through Github API v4.
type API struct {
	graphql fw.GraphQlRequest
}

// GetUserProfile retrieves user's email and name from Github.
func (g API) GetUserProfile(accessToken string) (entity.UserProfile, error) {
	type response struct {
		Viewer struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"viewer"`
	}

	var profileResponse response
	query := fw.GraphQlQuery{
		Query: `
query {
	viewer {
		email
		name
	}
}
`,
		Variables: nil,
	}

	err := g.sendGraphQlRequest(accessToken, query, &profileResponse)
	if err != nil {
		return entity.UserProfile{}, err
	}

	return entity.UserProfile{
		Email: profileResponse.Viewer.Email,
		Name:  profileResponse.Viewer.Name,
	}, nil
}

func (g API) sendGraphQlRequest(accessToken string, query fw.GraphQlQuery, response interface{}) error {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
	}
	return g.graphql.Query(query, headers, &response)
}

// NewAPI initializes Github API access service.
func NewAPI(graphql fw.GraphQlRequest) API {
	return API{
		graphql: graphql.RootUrl(githubAPI),
	}
}

package github

import (
	"fmt"
	"short/app/entity"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const githubAPI = "https://api.github.com/graphql"

var _ service.SSOAccount = (*Account)(nil)

// Account access user's account data through Github API v4.
type Account struct {
	graphql fw.GraphQlRequest
}

// GetSingleSignOnUser retrieves user's email and name from Github.
func (a Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
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

	err := a.sendGraphQlRequest(accessToken, query, &profileResponse)
	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: profileResponse.Viewer.Email,
		Name:  profileResponse.Viewer.Name,
	}, nil
}

func (a Account) sendGraphQlRequest(accessToken string, query fw.GraphQlQuery, response interface{}) error {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
	}
	return a.graphql.Query(query, headers, &response)
}

// NewAccount initializes Github account API client.
func NewAccount(graphql fw.GraphQlRequest) Account {
	return Account{
		graphql: graphql.RootUrl(githubAPI),
	}
}

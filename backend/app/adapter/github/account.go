package github

import (
	"fmt"

	"github.com/short-d/app/fw/graphql"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/sso"
)

const githubAPI = "https://api.github.com/graphql"

var _ sso.Account = (*Account)(nil)

// Account accesses user's account data through Github API v4.
type Account struct {
	gqlClient graphql.Client
}

// GetSingleSignOnUser retrieves user's email and name from Github.
func (a Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Viewer struct {
			ID    string `json:"id"`
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"viewer"`
	}

	var profileResponse response
	query := graphql.Query{
		Query: `
query {
	viewer {
		id
		email
		name
	}
}
`,
		Variables: nil,
	}

	err := a.sendGraphQLRequest(accessToken, query, &profileResponse)
	if err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		ID:    profileResponse.Viewer.ID,
		Email: profileResponse.Viewer.Email,
		Name:  profileResponse.Viewer.Name,
	}, nil
}

func (a Account) sendGraphQLRequest(accessToken string, query graphql.Query, response interface{}) error {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
	}
	return a.gqlClient.Query(query, headers, response)
}

// NewAccount initializes Github account API client.
func NewAccount(gqlClientFactory graphql.ClientFactory) Account {
	return Account{
		gqlClient: gqlClientFactory.NewClient(githubAPI),
	}
}

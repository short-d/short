package account

import (
	"fmt"
	"short/fw"
)

const githubApi = "https://api.github.com/graphql"

type Github struct {
	graphql fw.GraphQlRequest
}

func (g Github) GetEmail(accessToken string) (string, error) {
	headers := map[string]string{
		"Authorization": fmt.Sprintf("bearer %s", accessToken),
	}

	type response struct {
		Viewer struct {
			Email string `json:"email"`
		} `json:"viewer"`
	}

	var emailResponse response
	query := fw.GraphQlQuery{
		Query: `
query {
	viewer {
		email
	}
}
`,
		Variables: nil,
	}
	err := g.graphql.Query(query, headers, &emailResponse)
	if err != nil {
		return "", err
	}
	return emailResponse.Viewer.Email, nil
}

func NewGithub(graphql fw.GraphQlRequest) Github {
	return Github{
		graphql: graphql.RootUrl(githubApi),
	}
}

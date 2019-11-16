package graphql

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"short/app/adapter/db"
	"short/app/adapter/recaptcha"
	"short/app/usecase/auth"
	"short/app/usecase/keygen"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"testing"
	"time"

	"github.com/byliuyang/app/fw"
	"github.com/byliuyang/app/mdtest"
	"github.com/byliuyang/app/modern/mdcrypto"
)

func TestGraphQLAPI(t *testing.T) {
	authenticator := auth.NewAuthenticatorFake(time.Now(), time.Hour)
	graphqlAPI, err := newShortAPI(authenticator)
	mdtest.Equal(t, nil, err)
	mdtest.Equal(t, true, mdtest.IsGraphQlAPIValid(graphqlAPI))
}

func TestNewShortQueryViewer(t *testing.T) {
	nowStr := "2019-11-16T04:19:59Z"
	now := mustParseTime(t, nowStr)
	timer := mdtest.NewTimerFake(now)

	jwt := mdcrypto.NewJwtGo("test_secret")
	authToken, err := jwt.Encode(map[string]interface{}{
		"email":     "alpha@short-d.com",
		"issued_at": nowStr,
	})
	mdtest.Equal(t, nil, err)
	authenticator := auth.NewAuthenticator(jwt, timer, time.Hour)

	testCases := []struct {
		name        string
		query       graphQLQuery
		expHasErr   bool
		expResponse interface{}
	}{
		{
			name: "viewer",
			query: graphQLQuery{
				Query: `
query viewer($authToken: String) {
	viewer(authToken: $authToken) {
		email
	}
}
`,
				Variables: map[string]interface{}{
					"authToken": authToken,
				},
			},
			expResponse: map[string]interface{}{
				"viewer": map[string]interface{}{
					"email": "alpha@short-d.com",
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			serverURL, err := newGraphQLServerURL(authenticator)
			mdtest.Equal(t, nil, err)

			buf, err := json.Marshal(testCase.query)
			mdtest.Equal(t, nil, err)
			reqBody := bytes.NewReader(buf)
			res, err := http.Post(serverURL, "application/json", reqBody)
			if testCase.expHasErr {
				mdtest.NotEqual(t, nil, err)
				return
			}
			mdtest.Equal(t, nil, err)
			buf, err = ioutil.ReadAll(res.Body)
			mdtest.Equal(t, nil, err)

			var response graphQLResponse
			err = json.Unmarshal(buf, &response)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expResponse, response.Data)
		})
	}
}

type graphQLQuery struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

type graphQLResponse struct {
	Data map[string]interface{} `json:"data"`
}

func mustParseTime(t *testing.T, timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	mdtest.Equal(t, nil, err)
	return parsedTime
}

func newShortAPI(authenticator auth.Authenticator) (fw.GraphQlAPI, error) {
	sqlDB, _, err := mdtest.NewSQLStub()
	if err != nil {
		return nil, err
	}
	defer sqlDB.Close()

	urlRepo := db.NewURLSql(sqlDB)
	retriever := url.NewRetrieverPersist(urlRepo)
	urlRelationRepo := db.NewUserURLRelationSQL(sqlDB)
	keyGen := keygen.NewFake([]string{})
	creator := url.NewCreatorPersist(urlRepo, urlRelationRepo, &keyGen)

	s := recaptcha.NewFake()
	verifier := requester.NewVerifier(s)

	logger := mdtest.NewLoggerFake()
	tracer := mdtest.NewTracerFake()
	graphqlAPI := NewShort(&logger, &tracer, retriever, creator, verifier, authenticator)
	return graphqlAPI, nil
}

func newGraphQLServerURL(authenticator auth.Authenticator) (string, error) {
	graphqlAPI, err := newShortAPI(authenticator)
	if err != nil {
		return "", nil
	}
	server := mdtest.NewGraphQLServerFake(graphqlAPI)
	return server.URL(), nil
}

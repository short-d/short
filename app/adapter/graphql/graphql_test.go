package graphql

import (
	"short/app/adapter/recaptcha"
	"short/app/entity"
	"short/app/usecase/requester"
	"short/app/usecase/url"
	"short/mdtest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGraphQlAPI(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	urls := map[string]entity.URL{}
	retriever := url.NewRetrieverFake(urls)
	var availableUrls []string
	creator := url.NewCreatorFake(urls, availableUrls)

	s := recaptcha.NewFake()
	verifier := requester.NewVerifier(s)

	graphqlAPI := NewShort(mdtest.FakeLogger, mdtest.FakeTracer, retriever, creator, verifier)
	assert.True(t, mdtest.IsGraphQlAPIValid(graphqlAPI))
}

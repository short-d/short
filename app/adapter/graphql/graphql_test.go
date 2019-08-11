package graphql

import (
	"short/app/entity"
	"short/app/usecase/captcha"
	"short/app/usecase/url"
	"short/modern/mdtest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGraphQlApi(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	urls := map[string]entity.Url{}
	retriever := url.NewRetrieverFake(urls)
	creator := url.NewCreatorFake(urls)

	graphqlApi := NewShort(mdtest.FakeLogger, mdtest.FakeTracer, retriever, creator)
	assert.True(t, mdtest.IsGraphQlApiValid(graphqlApi))
}

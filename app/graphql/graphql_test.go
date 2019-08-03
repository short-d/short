package graphql

import (
	"testing"
	"tinyURL/modern/mdtest"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestGraphQlApi(t *testing.T) {
	db, _, err := sqlmock.New()

	assert.Nil(t, err)
	defer db.Close()

	graphqlApi := NewTinyUrl(mdtest.FakeLogger, mdtest.FakeTracer, db)
	assert.True(t, mdtest.IsGraphQlApiValid(graphqlApi))
}

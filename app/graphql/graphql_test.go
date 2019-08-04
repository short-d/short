package graphql

import (
	"short/modern/mdtest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestGraphQlApi(t *testing.T) {
	db, _, err := sqlmock.New()

	assert.Nil(t, err)
	defer db.Close()

	graphqlApi := NewShort(mdtest.FakeLogger, mdtest.FakeTracer, db)
	assert.True(t, mdtest.IsGraphQlApiValid(graphqlApi))
}

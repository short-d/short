package graphql

import (
	"testing"
	"tinyURL/modern/mdtest"

	"github.com/stretchr/testify/assert"
)

func TestGraphQlApi(t *testing.T) {
	graphqlApi := NewTinyUrl()
	assert.True(t, mdtest.IsGraphQlApiValid(graphqlApi))
}

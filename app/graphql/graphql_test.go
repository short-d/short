package graphql

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tinyURL/modern/mdtest"
)

func TestGraphQlApi(t *testing.T)  {
	graphqlApi := NewTinyUrl()
	assert.True(t, mdtest.IsGraphQlApiValid(graphqlApi))
}

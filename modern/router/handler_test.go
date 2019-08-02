package router

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpHandler_ServeHTTP(t *testing.T) {
	handler := NewHttpHandler()
	var gotParams Params

	err := handler.AddRoute("GET", false, "/users/:userId/articles/:articleId", func(w http.ResponseWriter, r *http.Request, params Params) {
		w.WriteHeader(200)

		body, err := ioutil.ReadAll(r.Body)
		assert.Nil(t, err)

		_, err = w.Write(body)
		assert.Nil(t, err)

		gotParams = params
	})
	assert.Nil(t, err)

	req, err := http.NewRequest("GET", "/users/fr4esw1rdf/articles/1dsd2DwxS/", strings.NewReader("Test"))
	assert.Nil(t, err)
	res := httptest.NewRecorder()

	handler.ServeHTTP(res, req)

	assert.Equal(t, 200, res.Code)

	body, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Test", string(body))

	expParams := Params{
		"articleId": "1dsd2DwxS",
		"userId":    "fr4esw1rdf",
	}
	assert.Equal(t, expParams, gotParams)
}

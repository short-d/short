package mdtest

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type TransportHandleFunc func(req *http.Request) (*http.Response, error)

type TransportMock struct {
	handle TransportHandleFunc
}

func (r TransportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.handle(req)
}

func NewTransportMock(handleFunc TransportHandleFunc) http.RoundTripper {
	return TransportMock{
		handle: handleFunc,
	}
}

func JSONResponse(jsonObj map[string]interface{}) (*http.Response, error) {
	jsonStr, err := json.Marshal(jsonObj)
	if err != nil {
		return nil, err
	}

	body := ioutil.NopCloser(strings.NewReader(string(jsonStr)))
	return &http.Response{
		StatusCode:    http.StatusOK,
		Body:          body,
		ContentLength: int64(len(jsonStr)),
	}, nil
}

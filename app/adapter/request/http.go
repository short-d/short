package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type Http struct {
	client http.Client
}

func (h Http) Json(
	method string,
	url string,
	headers map[string]string,
	body string,
	v interface{},
) error {
	req, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		return err
	}

	return nil
}

func NewHttp(client http.Client) Http {
	return Http{
		client: client,
	}
}

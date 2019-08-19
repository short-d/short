package mdrequest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"short/fw"
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

	req.Header.Add("Accept", "application/json")
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return errors.New(res.Status)
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

func NewHttp(client http.Client) fw.HttpRequest {
	return Http{
		client: client,
	}
}

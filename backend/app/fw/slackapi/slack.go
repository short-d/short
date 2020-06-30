package slackapi

import (
	"encoding/json"
	"github.com/short-d/app/fw/webreq"
	"net/http"
)

type Slack struct {
	webRequest webreq.HTTP
}

func (s Slack) SendMessage(webHookURL string, message string) error {
	type jsonRequest struct {
		Text string `json:"text"`
	}
	body, err := json.Marshal(jsonRequest{
		Text: message,
	})
	if err != nil {
		return err
	}
	return s.webRequest.JSON(http.MethodPost, webHookURL, map[string]string{}, string(body), nil)
}

func NewSlack(webRequest webreq.HTTP) Slack {
	return Slack{webRequest: webRequest}
}
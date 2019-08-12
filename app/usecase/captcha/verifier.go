package captcha

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const recaptchaVerifyApi = "https://www.google.com/recaptcha/api/siteverify"

type RecaptchaV3Secret string

type Verifier interface {
	IsHuman(captchaResponse string) (bool, error)
}

type RecaptchaV3Verifier struct {
	httpClient http.Client
	secret     string
}

type captchaVerifyResponse struct {
	Success       bool      `json:"success"`
	ChallengeTime time.Time `json:"challenge_ts"`
	Hostname      string    `json:"hostname"`
	Score         float32   `json:"score"`
	Action        string    `json:"action"`
}

func (r RecaptchaV3Verifier) IsHuman(captchaResponse string) (bool, error) {
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	req, err := http.NewRequest("POST", recaptchaVerifyApi, strings.NewReader(body))
	if err != nil {
		return false, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := r.httpClient.Do(req)
	if err != nil {
		return false, err
	}

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return false, err
	}

	apiRes := captchaVerifyResponse{}
	err = json.Unmarshal(buf, &apiRes)
	if err != nil {
		return false, err
	}

	if !apiRes.Success {
		return false, nil
	}

	return apiRes.Score > 0.7, nil
}

func NewRecaptchaV3Verifier(httpClient http.Client, secret RecaptchaV3Secret) Verifier {
	return RecaptchaV3Verifier{
		httpClient: httpClient,
		secret:     string(secret),
	}
}

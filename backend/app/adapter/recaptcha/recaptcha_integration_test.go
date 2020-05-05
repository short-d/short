// +build !integration all

package recaptcha

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/app/usecase/external"
)

func TestReCaptcha_Verify(t *testing.T) {
	expSecret := "ZPDIGNFj1EQJeNfs"
	expCaptchaResponse := "qHwha3zZh9G9mquEUOKZ"

	testCases := []struct {
		name         string
		httpResponse *http.Response
		httpErr      error
		expRes       external.VerifyResponse
	}{
		{
			name: "successful request with score = 0.8",
			httpResponse: &http.Response{
				StatusCode: http.StatusOK,
				Body: ioutil.NopCloser(bytes.NewReader([]byte(`
{
	"success": true,
	"action":  "homepage",
	"score":   0.8,
	"challenge_ts": "2006-01-02T15:04:05+07:00",
	"hostname": "s.time4hacks.com"
}
`,
				)))},
			expRes: external.VerifyResponse{
				Success:       true,
				Action:        "homepage",
				Score:         0.8,
				ChallengeTime: "2006-01-02T15:04:05+07:00",
				Hostname:      "s.time4hacks.com",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			httpRequest := webreq.NewHTTPFake(func(req *http.Request) (response *http.Response, e error) {
				assert.Equal(t, "https://www.google.com/recaptcha/api/siteverify", req.URL.String())
				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
				assert.Equal(t, "application/json", req.Header.Get("Accept"))

				buf, err := ioutil.ReadAll(req.Body)
				assert.Equal(t, nil, err)
				params, err := url.ParseQuery(string(buf))
				assert.Equal(t, nil, err)

				assert.Equal(t, expSecret, params.Get("secret"))
				assert.Equal(t, expCaptchaResponse, params.Get("response"))
				return testCase.httpResponse, testCase.httpErr
			})

			rc := NewService(httpRequest, expSecret)
			gotRes, err := rc.Verify(expCaptchaResponse)

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expRes, gotRes)
		})
	}
}

package service

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"short/app/adapter/request"
	"short/app/usecase/service"
	"short/modern/mdtest"
	"testing"

	"github.com/stretchr/testify/assert"
	//"time"
)

func TestReCaptcha_Verify(t *testing.T) {
	expSecret := ReCaptchaSecret("ZPDIGNFj1EQJeNfs")
	expCaptchaResponse := "qHwha3zZh9G9mquEUOKZ"
	//now := time.Now()

	testCases := []struct {
		name        string
		apiResponse map[string]interface{}
		expRes      service.VerifyResponse
	}{
		{
			name: "successful request with score = 0.8",
			apiResponse: map[string]interface{}{
				"success": true,
				"action":  "homepage",
				"score":   0.8,
				//"challenge_ts": now.Format(time.RFC3339),
				"hostname": "s.time4hacks.com",
			},
			expRes: service.VerifyResponse{
				Success: true,
				Action:  "homepage",
				Score:   0.8,
				//ChallengeTime: service.JSONTime(now),
				Hostname: "s.time4hacks.com",
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			transport := mdtest.NewTransportMock(func(req *http.Request) (response *http.Response, e error) {
				assert.Equal(t, "https://www.google.com/recaptcha/api/siteverify", req.URL.String())
				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
				assert.Equal(t, "application/json", req.Header.Get("Accept"))

				buf, err := ioutil.ReadAll(req.Body)
				assert.Nil(t, err)
				params, err := url.ParseQuery(string(buf))
				assert.Nil(t, err)

				assert.Equal(t, string(expSecret), params.Get("secret"))
				assert.Equal(t, expCaptchaResponse, params.Get("response"))
				return mdtest.JsonResponse(testCase.apiResponse)
			})

			client := http.Client{
				Transport: transport,
			}
			req := request.NewHttp(client)

			recaptcha := NewReCaptcha(req, expSecret)
			gotRes, err := recaptcha.Verify(expCaptchaResponse)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expRes, gotRes)
		})
	}
}

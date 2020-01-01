// +build !integration all

package recaptcha

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"short/app/usecase/service"
	"testing"

	"github.com/byliuyang/app/mdtest"
	"github.com/byliuyang/app/modern/mdrequest"
	//"time"
)

func TestReCaptcha_Verify(t *testing.T) {
	expSecret := "ZPDIGNFj1EQJeNfs"
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
				mdtest.Equal(t, "https://www.google.com/recaptcha/api/siteverify", req.URL.String())
				mdtest.Equal(t, "POST", req.Method)
				mdtest.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))
				mdtest.Equal(t, "application/json", req.Header.Get("Accept"))

				buf, err := ioutil.ReadAll(req.Body)
				mdtest.Equal(t, nil, err)
				params, err := url.ParseQuery(string(buf))
				mdtest.Equal(t, nil, err)

				mdtest.Equal(t, expSecret, params.Get("secret"))
				mdtest.Equal(t, expCaptchaResponse, params.Get("response"))
				return mdtest.JSONResponse(testCase.apiResponse)
			})

			client := http.Client{
				Transport: transport,
			}
			req := mdrequest.NewHTTP(client)

			rc := NewService(req, expSecret)
			gotRes, err := rc.Verify(expCaptchaResponse)
			mdtest.Equal(t, nil, err)
			mdtest.Equal(t, testCase.expRes, gotRes)
		})
	}
}

package captcha

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"short/modern/mdtest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRecaptchaV3Verifier_IsHuman(t *testing.T) {
	expSecret := "ZPDIGNFj1EQJeNfs"
	expCaptchaResponse := "qHwha3zZh9G9mquEUOKZ"

	testCases := []struct {
		name        string
		apiResponse map[string]interface{}
		expIsHuman  bool
	}{
		{
			name: "successful request with score = 0.8",
			apiResponse: map[string]interface{}{
				"success":      true,
				"score":        0.8,
				"challenge_ts": time.Now().Format(time.RFC3339),
				"hostname":     "s.time4hacks.com",
			},
			expIsHuman: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			transport := mdtest.NewTransportMock(func(req *http.Request) (response *http.Response, e error) {
				// https://developers.google.com/recaptcha/docs/verify
				assert.Equal(t, "https://www.google.com/recaptcha/api/siteverify", req.URL.String())
				assert.Equal(t, "POST", req.Method)
				assert.Equal(t, "application/x-www-form-urlencoded", req.Header.Get("Content-Type"))

				buf, err := ioutil.ReadAll(req.Body)
				assert.Nil(t, err)
				params, err := url.ParseQuery(string(buf))
				assert.Nil(t, err)

				assert.Equal(t, expSecret, params.Get("secret"))
				assert.Equal(t, expCaptchaResponse, params.Get("response"))
				return mdtest.JsonResponse(testCase.apiResponse)
			})

			client := http.Client{
				Transport: transport,
			}

			verifier := NewRecaptchaV3Verifier(client, RecaptchaV3Secret(expSecret))
			isHuman, err := verifier.IsHuman(expCaptchaResponse)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expIsHuman, isHuman)
		})
	}
}

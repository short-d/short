package recaptcha

import (
	"fmt"
	"net/http"

	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/app/usecase/external"
)

const verifyAPI = "https://www.google.com/recaptcha/api/siteverify"

var _ external.ReCaptcha = (*Service)(nil)

// Service consumes with Google ReCaptcha V3 APIs through network.
// https://developers.google.com/recaptcha/docs/verify
type Service struct {
	http   webreq.HTTP
	secret string
}

// Verify checks whether a captcha response is valid.
func (r Service) Verify(captchaResponse string) (external.VerifyResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	apiRes := external.VerifyResponse{}
	err := r.http.JSON(http.MethodPost, verifyAPI, headers, body, &apiRes)
	if err != nil {
		return external.VerifyResponse{}, err
	}
	return apiRes, nil
}

// NewService initializes ReCaptcha API consumer.
func NewService(http webreq.HTTP, secret string) Service {
	return Service{
		http:   http,
		secret: secret,
	}
}

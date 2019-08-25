package recaptcha

import (
	"fmt"
	"net/http"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

const verifyAPI = "https://www.google.com/recaptcha/api/siteverify"

var _ service.ReCaptcha = (*Service)(nil)

// https://developers.google.com/recaptcha/docs/verify
type Service struct {
	http   fw.HTTPRequest
	secret string
}

func (r Service) Verify(captchaResponse string) (service.VerifyResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	apiRes := service.VerifyResponse{}
	err := r.http.JSON(http.MethodPost, verifyAPI, headers, body, &apiRes)
	if err != nil {
		return service.VerifyResponse{}, err
	}
	return apiRes, nil
}

func NewService(http fw.HTTPRequest, secret string) Service {
	return Service{
		http:   http,
		secret: secret,
	}
}

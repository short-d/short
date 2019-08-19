package recaptcha

import (
	"fmt"
	"net/http"
	"short/app/usecase/service"
	"short/fw"
)

const verifyApi = "https://www.google.com/recaptcha/api/siteverify"

// https://developers.google.com/recaptcha/docs/verify
type Service struct {
	http   fw.HttpRequest
	secret string
}

func (r Service) Verify(captchaResponse string) (service.VerifyResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	apiRes := service.VerifyResponse{}
	err := r.http.Json(http.MethodPost, verifyApi, headers, body, &apiRes)
	if err != nil {
		return service.VerifyResponse{}, err
	}
	return apiRes, nil
}

func NewService(http fw.HttpRequest, secret string) service.ReCaptcha {
	return Service{
		http:   http,
		secret: secret,
	}
}

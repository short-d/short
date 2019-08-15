package recaptcha

import (
	"fmt"
	"net/http"
	"short/app/adapter/request"
	"short/app/usecase/service"
)

const verifyApi = "https://www.google.com/recaptcha/api/siteverify"

// https://developers.google.com/recaptcha/docs/verify
type Service struct {
	req    request.Http
	secret string
}

func (r Service) Verify(captchaResponse string) (service.VerifyResponse, error) {
	headers := map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	apiRes := service.VerifyResponse{}
	err := r.req.Json(http.MethodPost, verifyApi, headers, body, &apiRes)
	if err != nil {
		return service.VerifyResponse{}, err
	}
	return apiRes, nil
}

func NewService(req request.Http, secret string) service.ReCaptcha {
	return Service{
		req:    req,
		secret: secret,
	}
}

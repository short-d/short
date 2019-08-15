package service

import (
	"fmt"
	"net/http"
	"short/app/adapter/request"
	"short/app/usecase/service"
)

const verifyApi = "https://www.google.com/recaptcha/api/siteverify"

// https://developers.google.com/recaptcha/docs/verify
type ReCaptcha struct {
	req    request.Http
	secret string
}

type ReCaptchaSecret string

func (r ReCaptcha) Verify(captchaResponse string) (service.VerifyResponse, error) {
	body := fmt.Sprintf("secret=%s&response=%s", r.secret, captchaResponse)
	apiRes := service.VerifyResponse{}
	err := r.req.Json(http.MethodPost, verifyApi, map[string]string{}, body, &apiRes)
	if err != nil {
		return service.VerifyResponse{}, err
	}
	return apiRes, nil
}

func NewReCaptcha(req request.Http, secret ReCaptchaSecret) service.Recaptcha {
	return ReCaptcha{
		req:    req,
		secret: string(secret),
	}
}

package provider

import (
	"github.com/short-d/app/fw/webreq"
	"github.com/short-d/short/backend/app/adapter/recaptcha"
	"github.com/short-d/short/backend/app/usecase/requester"
)

// ReCaptchaSecret represents the secret used to verify reCAPTCHA.
type ReCaptchaSecret string

// NewReCaptchaService creates reCAPTCHA service with ReCaptchaSecret to uniquely identify secret during dependency injection.
func NewReCaptchaService(req webreq.HTTP, secret ReCaptchaSecret) requester.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

package provider

import (
	"short/app/adapter/recaptcha"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

// ReCaptchaSecret represents the secret used to verify reCAPTCHA.
type ReCaptchaSecret string

// NewReCaptchaService creates reCAPTCHA service with ReCaptchaSecret to uniquely identify secret during dependency injection.
func NewReCaptchaService(req fw.HTTPRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/recaptcha"
	"github.com/short-d/short/app/usecase/service"
)

// ReCaptchaSecret represents the secret used to verify reCAPTCHA.
type ReCaptchaSecret string

// NewReCaptchaService creates reCAPTCHA service with ReCaptchaSecret to uniquely identify secret during dependency injection.
func NewReCaptchaService(req fw.HTTPRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

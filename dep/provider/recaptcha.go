package provider

import (
	"short/app/adapter/recaptcha"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

// ReCaptchaSecret reCAPTCHA secret.
type ReCaptchaSecret string

// ReCaptchaService initializes reCAPTCHA service.
func ReCaptchaService(req fw.HTTPRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

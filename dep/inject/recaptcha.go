package inject

import (
	"short/app/adapter/recaptcha"
	"short/app/usecase/service"
	"short/fw"
)

type ReCaptchaSecret string

func ReCaptchaService(req fw.HttpRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

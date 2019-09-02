package inject

import (
	"short/app/adapter/recaptcha"
	"short/app/usecase/service"

	"github.com/byliuyang/app/fw"
)

type ReCaptchaSecret string

func ReCaptchaService(req fw.HTTPRequest, secret ReCaptchaSecret) service.ReCaptcha {
	return recaptcha.NewService(req, string(secret))
}

package service

import "short/app/usecase/service"

type ReCaptchaFake struct {
}

func (v ReCaptchaFake) Verify(recaptchaResponse string) (service.VerifyResponse, error) {
	panic("implement me")
}

func NewServiceFake() service.Recaptcha {
	return ReCaptchaFake{}
}

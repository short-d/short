package service

var _ ReCaptcha = (*ReCaptchaFake)(nil)

// ReCaptchaFake represents in memory implementation of ReCaptcha service.
type ReCaptchaFake struct {
	verifyResponse VerifyResponse
}

// ReCaptchaFake verifies captcha response.
func (r ReCaptchaFake) Verify(recaptchaResponse string) (VerifyResponse, error) {
	return r.verifyResponse, nil
}

// NewReCaptchaFake creates in memory fake reCaptcha service with predefined
// response.
func NewReCaptchaFake(verifyResponse VerifyResponse) ReCaptchaFake {
	return ReCaptchaFake{verifyResponse: verifyResponse}
}

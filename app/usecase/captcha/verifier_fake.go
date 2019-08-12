package captcha

type VerifierFake struct {
}

func (v VerifierFake) IsHuman(captchaResponse string) (bool, error) {
	return true, nil
}

func NewVerifierFake() Verifier {
	return VerifierFake{}
}

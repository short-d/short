package service

// VerifyResponse represents the JSON response received from ReCaptcha Verify
// API.
type VerifyResponse struct {
	Success       bool    `json:"success"`
	ChallengeTime string  `json:"challenge_ts"`
	Hostname      string  `json:"hostname"`
	Score         float32 `json:"score"`
	Action        string  `json:"action"`
}

// ReCaptcha verifies captcha response.
type ReCaptcha interface {
	Verify(recaptchaResponse string) (VerifyResponse, error)
}

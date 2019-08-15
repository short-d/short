package service

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime time.Time

func (j *JSONTime) UnmarshalJSON(buf []byte) error {
	var timeStr string
	err := json.Unmarshal(buf, &timeStr)
	if err != nil {
		return err
	}
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return err
	}
	fmt.Println(t)
	*j = JSONTime(t)
	return nil
}

type VerifyResponse struct {
	Success       bool     `json:"success"`
	ChallengeTime JSONTime `json:"challenge_ts"`
	Hostname      string   `json:"hostname"`
	Score         float32  `json:"score"`
	Action        string   `json:"action"`
}

type ReCaptcha interface {
	Verify(recaptchaResponse string) (VerifyResponse, error)
}

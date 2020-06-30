package resolver

import (
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

type Feedback struct {
	feedback entity.Feedback
}

func (f Feedback) AppID() string {
	return f.feedback.AppID
}

func (f Feedback) FeedbackID() string {
	return f.feedback.FeedbackID
}

func (f Feedback) CustomerRating() int32 {
	return int32(f.feedback.CustomerRating)
}

func (f Feedback) Comment() *string {
	return f.feedback.Comment
}

func (f Feedback) CustomerEmail() *string {
	return f.feedback.CustomerEmail
}

func (f Feedback) ReceivedAt() scalar.Time {
	return scalar.Time{Time: f.feedback.ReceivedAt}
}

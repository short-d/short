package entity

import "time"

type Feedback struct {
	AppID string
	FeedbackID string
	CustomerRating int
	Comment *string
	CustomerEmail *string
	ReceivedAt time.Time
}

func (f Feedback) GetComment(defaultVal string) string {
	if f.Comment == nil {
		return defaultVal
	}
	return *f.Comment
}

func (f Feedback) GetCustomerEmail(defaultVal string) string {
	if f.CustomerEmail == nil {
		return defaultVal
	}
	return *f.CustomerEmail
}

type FeedbackInput struct {
	AppID *string
	FeedbackID *string
	CustomerRating *int
	Comment *string
	CustomerEmail *string
	ReceivedAt *time.Time
}

func (f FeedbackInput) GetAppID() string {
	if f.AppID == nil {
		return ""
	}
	return *f.AppID
}

func (f FeedbackInput) GetFeedbackID() string {
	if f.FeedbackID == nil {
		return ""
	}
	return *f.FeedbackID
}

func (f FeedbackInput) GetCustomerRating() int {
	if f.CustomerRating == nil {
		return 0
	}
	return *f.CustomerRating
}

func (f FeedbackInput) GetReceivedAt() time.Time {
	if f.ReceivedAt == nil {
		return time.Time{}
	}
	return *f.ReceivedAt
}
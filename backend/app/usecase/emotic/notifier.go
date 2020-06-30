package emotic

import "github.com/short-d/short/backend/app/entity"

type Notifier interface {
	NotifyFeedbackReceived(feedback entity.Feedback) error
}

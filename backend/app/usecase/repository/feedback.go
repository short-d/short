package repository

import (
	"github.com/short-d/short/backend/app/entity"
)

type Feedback interface {
	CreateFeedback(input entity.FeedbackInput) (entity.Feedback, error)
}

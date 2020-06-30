package sqldb

import (
	"database/sql"
	"fmt"
	"github.com/short-d/short/backend/app/adapter/sqldb/table"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/repository"
)

var _ repository.Feedback = (*FeedbackSQL)(nil)

type FeedbackSQL struct {
	db *sql.DB
}

func (f FeedbackSQL) CreateFeedback(input entity.FeedbackInput) (entity.Feedback, error) {
	stmt := fmt.Sprintf(`
INSERT INTO "%s"("%s", "%s", "%s", "%s", "%s", "%s")
VALUES ($1, $2, $3, $4, $5, $6);
`,
		table.Feedback.TableName,
		table.Feedback.ColumnAppID,
		table.Feedback.ColumnFeedbackID,
		table.Feedback.ColumnCustomerRating,
		table.Feedback.ColumnComment,
		table.Feedback.ColumnCustomerEmail,
		table.Feedback.ColumnReceivedAt,
	)

	_, err := f.db.Exec(
		stmt,
		input.GetAppID(),
		input.GetFeedbackID(),
		input.GetCustomerRating(),
		input.Comment,
		input.CustomerEmail,
		input.GetReceivedAt(),
		)
	return entity.Feedback{
		AppID: input.GetAppID(),
		FeedbackID: input.GetFeedbackID(),
		CustomerRating: input.GetCustomerRating(),
		Comment: input.Comment,
		CustomerEmail: input.CustomerEmail,
		ReceivedAt: input.GetReceivedAt(),
	}, err
}

func NewFeedbackSQL(db *sql.DB) FeedbackSQL {
	return FeedbackSQL{
		db: db,
	}
}

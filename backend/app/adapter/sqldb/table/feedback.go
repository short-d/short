package table

var Feedback = struct {
	TableName            string
	ColumnAppID          string
	ColumnFeedbackID     string
	ColumnCustomerRating string
	ColumnComment        string
	ColumnCustomerEmail  string
	ColumnReceivedAt     string
}{
	TableName:            "feedback",
	ColumnAppID:          "app_id",
	ColumnFeedbackID:     "feedback_id",
	ColumnCustomerRating: "customer_rating",
	ColumnComment:        "comment",
	ColumnCustomerEmail:  "customer_email",
	ColumnReceivedAt:     "received_at",
}

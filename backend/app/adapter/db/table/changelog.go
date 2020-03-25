package table

// ChangeLog represents database table columns for 'change_log' table
var ChangeLog = struct {
	TableName             string
	ColumnID              string
	ColumnTitle           string
	ColumnSummaryMarkdown string
	ColumnReleasedAt      string
}{
	TableName:             "change_log",
	ColumnID:              "id",
	ColumnTitle:           "title",
	ColumnSummaryMarkdown: "summary_markdown",
	ColumnReleasedAt:      "released_at",
}

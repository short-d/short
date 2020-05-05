package table

// GithubSSO represents database table columns for 'github_sso' table.
var GithubSSO = struct {
	TableName          string
	ColumnGithubUserID string
	ColumnShortUserID  string
}{
	TableName:          "github_sso",
	ColumnGithubUserID: "github_user_id",
	ColumnShortUserID:  "short_user_id",
}

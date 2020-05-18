package table

// GoogleSSO represents database table columns for 'google_sso' table.
var GoogleSSO = struct {
	TableName          string
	ColumnGoogleUserID string
	ColumnShortUserID  string
}{
	TableName:          "google_sso",
	ColumnGoogleUserID: "google_user_id",
	ColumnShortUserID:  "short_user_id",
}

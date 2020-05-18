package table

// FacebookSSO represents database table columns for 'facebook_sso' table.
var FacebookSSO = struct {
	TableName            string
	ColumnFacebookUserID string
	ColumnShortUserID    string
}{
	TableName:            "facebook_sso",
	ColumnFacebookUserID: "facebook_user_id",
	ColumnShortUserID:    "short_user_id",
}

package table

// UserChangeLog represents database table columns for 'user_changelog' table
var UserChangeLog = struct {
	TableName          string
	ColumnUserID       string
	ColumnEmail        string
	ColumnLastViewedAt string
}{
	TableName:          "user_changelog",
	ColumnUserID:       "user_id",
	ColumnEmail:        "email",
	ColumnLastViewedAt: "last_viewed_at",
}

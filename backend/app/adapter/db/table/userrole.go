package table

// UserRole represents database table columns for 'user_role' table
var UserRole = struct {
	TableName    string
	ColumnUserID string
	ColumnRoles  string
}{
	TableName:    "user_role",
	ColumnUserID: "user_id",
	ColumnRoles:  "roles",
}

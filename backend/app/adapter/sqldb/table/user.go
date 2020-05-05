package table

// User represents database table columns for 'user' table
var User = struct {
	TableName            string
	ColumnID             string
	ColumnEmail          string
	ColumnName           string
	ColumnLastSignedInAt string
	ColumnCreatedAt      string
	ColumnUpdatedAt      string
}{
	TableName:            "user",
	ColumnID:             "id",
	ColumnEmail:          "email",
	ColumnName:           "name",
	ColumnLastSignedInAt: "last_signed_in_at",
	ColumnCreatedAt:      "created_at",
	ColumnUpdatedAt:      "updated_at",
}

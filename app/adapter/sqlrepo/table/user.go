package table

var User = struct {
	TableName            string
	ColumnEmail          string
	ColumnName           string
	ColumnLastSignedInAt string
	ColumnCreatedAt      string
	ColumnUpdatedAt      string
}{
	TableName:            "user",
	ColumnEmail:          "email",
	ColumnName:           "name",
	ColumnLastSignedInAt: "last_signed_in_at",
	ColumnCreatedAt:      "created_at",
	ColumnUpdatedAt:      "updated_at",
}

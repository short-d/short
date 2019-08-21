package table

var User = struct {
	TableName            string
	ColumnEmail          string
	ColumnName           string
	ColumnLastSignedInAt string
	ColumnCreatedAt      string
	ColumnUpdatedAt      string
}{
	TableName:            "User",
	ColumnEmail:          "email",
	ColumnName:           "name",
	ColumnLastSignedInAt: "lastSignedInAt",
	ColumnCreatedAt:      "createdAt",
	ColumnUpdatedAt:      "updatedAt",
}

package table

// URL represents database table columns for 'url' table
var URL = struct {
	TableName         string
	ColumnAlias       string
	ColumnOriginalURL string
	ColumnCreatedAt   string
	ColumnExpireAt    string
	ColumnUpdatedAt   string
}{
	TableName:         "short_link",
	ColumnAlias:       "alias",
	ColumnOriginalURL: "long_link",
	ColumnCreatedAt:   "created_at",
	ColumnExpireAt:    "expire_at",
	ColumnUpdatedAt:   "updated_at",
}

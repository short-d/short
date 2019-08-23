package table

var URL = struct {
	TableName         string
	ColumnAlias       string
	ColumnOriginalURL string
	ColumnCreatedAt   string
	ColumnExpireAt    string
	ColumnUpdatedAt   string
}{
	TableName:         "url",
	ColumnAlias:       "alias",
	ColumnOriginalURL: "original_url",
	ColumnCreatedAt:   "created_at",
	ColumnExpireAt:    "expire_at",
	ColumnUpdatedAt:   "updated_at",
}

package table

var URL = struct {
	TableName         string
	ColumnAlias       string
	ColumnOriginalURL string
	ColumnCreatedAt   string
	ColumnExpireAt    string
	ColumnUpdatedAt   string
}{
	TableName:         "Url",
	ColumnAlias:       "alias",
	ColumnOriginalURL: "originalUrl",
	ColumnCreatedAt:   "createdAt",
	ColumnExpireAt:    "expireAt",
	ColumnUpdatedAt:   "updatedAt",
}

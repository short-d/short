package table

var Url = struct {
	TableName         string
	ColumnAlias       string
	ColumnOriginalUrl string
	ColumnCreatedAt   string
	ColumnExpireAt    string
	ColumnUpdatedAt   string
}{
	TableName:         "Url",
	ColumnAlias:       "alias",
	ColumnOriginalUrl: "originalUrl",
	ColumnCreatedAt:   "createdAt",
	ColumnExpireAt:    "expireAt",
	ColumnUpdatedAt:   "updatedAt",
}

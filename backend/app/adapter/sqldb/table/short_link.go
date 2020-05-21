package table

// ShortLink represents database table columns for 'short_link' table
var ShortLink = struct {
	TableName       string
	ColumnAlias     string
	ColumnLongLink  string
	ColumnCreatedAt string
	ColumnExpireAt  string
	ColumnUpdatedAt string
}{
	TableName:       "short_link",
	ColumnAlias:     "alias",
	ColumnLongLink:  "long_link",
	ColumnCreatedAt: "created_at",
	ColumnExpireAt:  "expire_at",
	ColumnUpdatedAt: "updated_at",
}

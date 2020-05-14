package table

// UserURLRelation represents database table columns for 'user_url_relation' table
var UserURLRelation = struct {
	TableName      string
	ColumnUserID   string
	ColumnURLAlias string
}{
	TableName:      "user_url_relation",
	ColumnUserID:   "user_id",
	ColumnURLAlias: "url_alias",
}

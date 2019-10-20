package table

// UserURLRelation represents database table columns for 'user_url_relation' table
var UserURLRelation = struct {
	TableName       string
	ColumnUserEmail string
	ColumnURLAlias  string
}{
	TableName:       "user_url_relation",
	ColumnUserEmail: "user_email",
	ColumnURLAlias:  "url_alias",
}

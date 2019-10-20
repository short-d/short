package table

// UserURLRelation represents database table columns for 'user_url_relation' table
var UserURLRelation = struct {
	TableName       string
	ColumnUserEmail string
	ColumnUrlAlias  string
}{
	TableName:       "user_url_relation",
	ColumnUserEmail: "user_email",
	ColumnUrlAlias:  "url_alias",
}

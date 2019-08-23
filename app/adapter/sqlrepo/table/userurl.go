package table

var UserURL = struct {
	TableName       string
	ColumnUserEmail string
	ColumnUrlAlias  string
}{
	TableName:       "user_url_relation",
	ColumnUserEmail: "user_email",
	ColumnUrlAlias:  "url_alias",
}

package table

// UserShortLink represents database table columns for 'user_short_link' table
var UserShortLink = struct {
	TableName            string
	ColumnUserID         string
	ColumnShortLinkAlias string
}{
	TableName:            "user_short_link",
	ColumnUserID:         "user_id",
	ColumnShortLinkAlias: "short_link_alias",
}

package table

// ShortLink represents database table columns for 'short_link' table
var ShortLink = struct {
	TableName                string
	ColumnAlias              string
	ColumnLongLink           string
	ColumnCreatedAt          string
	ColumnExpireAt           string
	ColumnUpdatedAt          string
	ColumnOGTitle            string
	ColumnOGDescription      string
	ColumnOGImageURL         string
	ColumnTwitterTitle       string
	ColumnTwitterDescription string
	ColumnTwitterImageURL    string
}{
	TableName:                "short_link",
	ColumnAlias:              "alias",
	ColumnLongLink:           "long_link",
	ColumnCreatedAt:          "created_at",
	ColumnExpireAt:           "expire_at",
	ColumnUpdatedAt:          "updated_at",
	ColumnOGTitle:            "og_title",
	ColumnOGDescription:      "og_description",
	ColumnOGImageURL:         "og_image_url",
	ColumnTwitterTitle:       "twitter_title",
	ColumnTwitterDescription: "twitter_description",
	ColumnTwitterImageURL:    "twitter_image_url",
}
package entity

type MetaTags struct {
	MetaOGTags
	MetaTwitterTags
}

type MetaOGTags struct {
	OGTitle       string
	OGDescription string
	OGImageURL    string
}

type MetaTwitterTags struct {
	TwitterTitle       string
	TwitterDescription string
	TwitterImageURL    string
}

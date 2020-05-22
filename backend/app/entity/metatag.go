package entity

type MetaTags struct {
	OpenGraphTags
	TwitterTags
}

type OpenGraphTags struct {
	OpenGraphTitle       *string
	OpenGraphDescription *string
	OpenGraphImageURL    *string
}

type TwitterTags struct {
	TwitterTitle       *string
	TwitterDescription *string
	TwitterImageURL    *string
}

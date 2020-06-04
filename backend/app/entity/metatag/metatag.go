package metatag

// OpenGraph represents OpenGraph Meta Tags for a ShortLink.
type OpenGraph struct {
	Title       *string
	Description *string
	ImageURL    *string
}

// Twitter represents Twitter Meta Tags for a ShortLink.
type Twitter struct {
	Title       *string
	Description *string
	ImageURL    *string
}

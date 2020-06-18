package provider

import "time"

// SearchTimeout represents timeout duration of a search request.
type SearchTimeout time.Duration

// NewSearchTimeout creates Duration.
func NewSearchTimeout(timeout SearchTimeout) time.Duration {
	return time.Duration(timeout)
}

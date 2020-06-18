package provider

import "time"

// SearchAPITimeout represents timeout duration of a search request.
type SearchAPITimeout time.Duration

// NewSearchAPITimeout creates Duration.
func NewSearchAPITimeout(timeout SearchAPITimeout) time.Duration {
	return time.Duration(timeout)
}

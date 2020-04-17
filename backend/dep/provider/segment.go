package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdanalytics"
)

// SegmentAPIKey represents credential for Segment APIs.
type SegmentAPIKey string

// NewSegment creates Segment with SegmentAPIKey to uniquely identify apiKey
// during dependency injection.
func NewSegment(apiKey SegmentAPIKey, timer fw.Timer, logger fw.Logger) mdanalytics.Segment {
	return mdanalytics.NewSegment(string(apiKey), timer, logger)
}

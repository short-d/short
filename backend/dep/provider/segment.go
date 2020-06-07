package provider

import (
	"github.com/short-d/app/fw/analytics"
	"github.com/short-d/app/fw/logger"
	"github.com/short-d/app/fw/timer"
)

// SegmentAPIKey represents credential for Segment APIs.
type SegmentAPIKey string

// NewSegment creates Segment with SegmentAPIKey to uniquely identify apiKey
// during dependency injection.
func NewSegment(apiKey SegmentAPIKey, timer timer.Timer, logger logger.Logger) analytics.Segment {
	return analytics.NewSegment(string(apiKey), timer, logger)
}

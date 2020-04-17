package provider

import (
	"github.com/short-d/app/fw"
	"github.com/short-d/app/modern/mdanalytics"
)

type SegmentAPIKey string

func NewSegment(apiKey SegmentAPIKey, timer fw.Timer, logger fw.Logger) mdanalytics.Segment {
	return mdanalytics.NewSegment(string(apiKey), timer, logger)
}

package must

import (
	"testing"
	"time"
)

// Time parses time from it's string representation. This simplifies test case
// inputs involving configurable time and normalizes it to UTC.
// TODO(issue#977): replace mustParseTime with must.Time across the codebase.
func Time(t *testing.T, value string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatal(err)
	}
	return parsedTime.UTC()
}

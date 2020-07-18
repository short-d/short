package must

import (
	"testing"
	"time"
)

// Time parses time from it's string representation. This simplifies test case
// inputs involving configurable time and normalizes it to UTC.
func Time(t *testing.T, value string) time.Time {
	tm, err := time.Parse(time.RFC3339, value)
	if err != nil {
		t.Fatal(err)
	}
	return tm.UTC()
}

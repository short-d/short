package ptr

import "time"

// Time returns the address of a time variable.
func Time(t time.Time) *time.Time {
	return &t
}

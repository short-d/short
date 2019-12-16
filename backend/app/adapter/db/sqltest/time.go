package sqltest

import "time"

var dateTimeFmt = "2006-01-02 15:04:05"

// MustParseSQLTime parses time.Time from SQL time string.
func MustParseSQLTime(dateTime string) *time.Time {
	if dateTime == "NULL" {
		return nil
	}

	dt, err := time.Parse(dateTimeFmt, dateTime)
	if err != nil {
		panic(err)
	}
	return &dt
}

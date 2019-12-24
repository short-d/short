package db

import "time"

func utc(dbTime *time.Time) *time.Time {
	if dbTime == nil {
		return nil
	}
	dbTimeUTC := dbTime.UTC()
	return &dbTimeUTC
}

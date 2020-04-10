package unit

import (
	"errors"
	"strconv"
	"time"
)

const (
	oneDay  = time.Hour * 24
	oneWeak = oneDay * 7
)

// ParseDuration reads a human-readable duration and returns duration in seconds
func ParseDuration(readableDuration string) (time.Duration, error) {
	var (
		duration time.Duration
		err      error
		value    int
	)
	unit := readableDuration[len(readableDuration)-1]
	value, err = strconv.Atoi(readableDuration[:len(readableDuration)-1])
	switch unit {
	case 's':
		duration = time.Duration(value) * time.Second
	case 'm':
		duration = time.Duration(value) * time.Minute
	case 'h':
		duration = time.Duration(value) * time.Hour
	case 'd':
		duration = time.Duration(value) * oneDay
	case 'w':
		duration = time.Duration(value) * oneWeak
	default:
		err = errors.New("unknown time type")
	}
	return duration, err
}

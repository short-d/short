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
func ParseDuration(s string) (time.Duration, error) {
	var (
		duration time.Duration
		err      error
		length   int
	)
	timeVariable := s[len(s)-1]
	switch timeVariable {
	case 's':
		length, err = strconv.Atoi(s[:len(s)-1])
		duration = time.Duration(length) * time.Second
	case 'm':
		length, err = strconv.Atoi(s[:len(s)-1])
		duration = time.Duration(length) * time.Minute
	case 'h':
		length, err = strconv.Atoi(s[:len(s)-1])
		duration = time.Duration(length) * time.Hour
	case 'd':
		length, err = strconv.Atoi(s[:len(s)-1])
		duration = time.Duration(length) * oneDay
	case 'w':
		length, err = strconv.Atoi(s[:len(s)-1])
		duration = time.Duration(length) * oneWeak
	default:
		err = errors.New("unknown time type")
	}
	return duration, err
}

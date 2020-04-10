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
	timeVariable := s[len(s)-1]
	switch timeVariable {
	case 's':
		length, err := strconv.Atoi(s[:len(s)-1])
		return time.Duration(length) * time.Second, err
	case 'm':
		length, err := strconv.Atoi(s[:len(s)-1])
		return time.Duration(length) * time.Minute, err
	case 'h':
		length, err := strconv.Atoi(s[:len(s)-1])
		return time.Duration(length) * time.Hour, err
	case 'd':
		length, err := strconv.Atoi(s[:len(s)-1])
		return time.Duration(length) * oneDay, err
	case 'w':
		length, err := strconv.Atoi(s[:len(s)-1])
		return time.Duration(length) * oneWeak, err
	default:
		return 0, errors.New("unknown time type")
	}
}

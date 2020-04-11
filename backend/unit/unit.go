package unit

import (
	"errors"
	"strconv"
	"time"
)

// TODO(issue#659): extend and add other units

const (
	oneDay  = time.Hour * 24
	oneWeak = oneDay * 7
)

// ParseDuration converts a human-readable duration and returns duration into seconds
func ParseDuration(readableDuration string) (time.Duration, error) {
	if len(readableDuration) == 0 {
		return 0, errors.New("readable duration can't be empty")
	}

	end := len(readableDuration) - 1
	valueStr := readableDuration[:end]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0, err
	}

	unit := readableDuration[end]
	switch unit {
	case 's':
		return time.Duration(value) * time.Second, nil
	case 'm':
		return time.Duration(value) * time.Minute, nil
	case 'h':
		return time.Duration(value) * time.Hour, nil
	case 'd':
		return time.Duration(value) * oneDay, nil
	case 'w':
		return time.Duration(value) * oneWeak, nil
	default:
		return 0, errors.New("unknown duration unit")
	}
}

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
// TODO: extend to ms or add other units, such as storage (B, KB, MB, GB, TB, PB) and temperature ( C, F).
func ParseDuration(readableDuration string) (time.Duration, error) {
	if len(readableDuration) == 0 {
		return 0, errors.New("duration string is empty")
	}

	value, err := strconv.Atoi(readableDuration[:len(readableDuration)-1])
	if err != nil {
		return 0, err
	}

	unit := readableDuration[len(readableDuration)-1]
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
		return 0, errors.New("unknown time type")
	}
}

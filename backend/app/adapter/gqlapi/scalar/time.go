package scalar

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/short-d/app/fw/graphql"
)

var _ graphql.Scalar = &Time{}

// Time maps GraphQL Time scalar to time.Time.
type Time struct {
	time.Time
}

// ImplementsGraphQLType checks whether a given GraphQL scalar is Time.
func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

// UnmarshalGraphQL parses GraphQL Time scalar from various data format.
func (t *Time) UnmarshalGraphQL(input interface{}) error {
	var timeTmp time.Time
	var err error

	switch input := input.(type) {
	case time.Time:
		timeTmp = input
		err = nil
	case string:
		timeTmp, err = timeFromString(input)
	case int:
		timeTmp, err = timeFromInt(input)
	case float64:
		timeTmp, err = timeFromFloat(input)
	default:
		err = errors.New("wrong type")
	}

	if err != nil {
		return err
	}

	t.Time = timeTmp
	return nil
}

// MarshalJSON serializes GraphQL Time to JSON.
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func timeFromString(val interface{}) (time.Time, error) {
	return time.Parse(time.RFC3339, val.(string))
}

func timeFromInt(val interface{}) (time.Time, error) {
	return time.Unix(int64(val.(int)), 0), nil
}

func timeFromFloat(val interface{}) (time.Time, error) {
	return time.Unix(int64(val.(int)), 0), nil
}

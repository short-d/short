package scalar

import (
	"encoding/json"
	"time"
	"tinyURL/app/graphql/parser"
	"tinyURL/fw"

	"github.com/pkg/errors"
)

var _ fw.Scalar = &Time{}

type Time struct {
	time.Time
}

func (Time) ImplementsGraphQLType(name string) bool {
	return name == "Time"
}

func (t *Time) UnmarshalGraphQL(input interface{}) error {
	var timeTmp time.Time
	var err error

	switch input := input.(type) {
	case time.Time:
		timeTmp = input
		err = nil
	case string:
		timeTmp, err = parser.NewStringTime().FromVal(input)
	case int:
		timeTmp, err = parser.NewIntTime().FromVal(input)
	case float64:
		timeTmp, err = parser.NewFloatTime().FromVal(input)
	default:
		err = errors.New("wrong type")
	}

	if err != nil {
		return err
	}

	t.Time = timeTmp
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

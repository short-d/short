package parser

import "time"

type Time interface {
	FromVal(val interface{}) (time.Time, error)
}

type StringTime struct {
}

func (StringTime) FromVal(val interface{}) (time.Time, error) {
	return time.Parse(time.RFC3339, val.(string))
}

func NewStringTime() Time {
	return StringTime{}
}

type IntTime struct {
}

func (IntTime) FromVal(val interface{}) (time.Time, error) {
	return time.Unix(int64(val.(int)), 0), nil
}

func NewIntTime() Time {
	return IntTime{}
}

type FloatTime struct {
}

func (FloatTime) FromVal(val interface{}) (time.Time, error) {
	return time.Unix(int64(val.(int)), 0), nil
}

func NewFloatTime() Time {
	return FloatTime{}
}

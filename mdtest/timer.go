package mdtest

import (
	"time"
)

type Timer struct {
	currentTime time.Time
}

func (t *Timer) SetCurrentTime(currentTime time.Time)  {
	t.currentTime = currentTime
}

func (t Timer) Now() time.Time {
	return t.currentTime
}

func NewFakeTimer(currentTime time.Time) Timer {
	return Timer{
		currentTime: currentTime,
	}
}

package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
)

type ChangeLog struct {
	changeLog []*Change
}

func (c ChangeLog) Changes() []*Change {
	return c.changeLog
}

func (c ChangeLog) LastViewedAt() *scalar.Time {
	currentTime := time.Now()
	return &scalar.Time{Time: currentTime}
}

func newChangeLog(changelog []entity.Change) *ChangeLog {
	var changes []*Change
	for _, v := range changelog {
		changes = append(changes, newChange(v))
	}

	return &ChangeLog{changeLog: changes}
}

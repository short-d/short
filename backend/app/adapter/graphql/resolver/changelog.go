package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
)

type ChangeLog struct {
	changeLog    []Change
	lastViewedAt time.Time
}

func (c ChangeLog) Changes() []Change {
	return c.changeLog
}

func (c ChangeLog) LastViewedAt() *scalar.Time {
	return &scalar.Time{Time: c.lastViewedAt}
}

func newChangeLog(changelog []entity.Change, lastViewedAt time.Time) ChangeLog {
	var changes []Change
	for _, v := range changelog {
		changes = append(changes, newChange(v))
	}

	return ChangeLog{changeLog: changes, lastViewedAt: lastViewedAt}
}

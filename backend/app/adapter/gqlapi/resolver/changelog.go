package resolver

import (
	"time"

	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

// ChangeLog retrieves full change log and the time when the user viewed it.
type ChangeLog struct {
	changeLog    []Change
	lastViewedAt *time.Time
}

// Changes retrieves full change log
func (c ChangeLog) Changes() []Change {
	return c.changeLog
}

// LastViewedAt retrieves the time the user viewed the change log.
func (c ChangeLog) LastViewedAt() *scalar.Time {
	if c.lastViewedAt == nil {
		return nil
	}
	return &scalar.Time{Time: *c.lastViewedAt}
}

func newChangeLog(changeLog []entity.Change, lastViewedAt *time.Time) ChangeLog {
	var changes []Change
	for _, v := range changeLog {
		changes = append(changes, newChange(v))
	}

	return ChangeLog{changeLog: changes, lastViewedAt: lastViewedAt}
}

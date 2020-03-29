package resolver

import (
	"time"

	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
)

// ChangeLog represents ChangeLog entity and user Last Viewed ChangeLog time
type ChangeLog struct {
	changeLog    []Change
	lastViewedAt time.Time
}

// Changes retrieves full ChangeLog
func (c ChangeLog) Changes() []Change {
	return c.changeLog
}

// LastViewedAt retrieves lastViewedAt for given user
func (c ChangeLog) LastViewedAt() *scalar.Time {
	return &scalar.Time{Time: c.lastViewedAt}
}

func newChangeLog(changeLog []entity.Change, lastViewedAt time.Time) ChangeLog {
	var changes []Change
	for _, v := range changeLog {
		changes = append(changes, newChange(v))
	}

	return ChangeLog{changeLog: changes, lastViewedAt: lastViewedAt}
}

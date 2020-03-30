package resolver

import (
	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
)

// Change represents single change in ChangeLog
type Change struct {
	change entity.Change
}

// ID retrieves ID of Change entity
func (c Change) ID() string {
	return c.change.ID
}

// Title retrieves Title of Change entity
func (c Change) Title() string {
	return c.change.Title
}

// SummaryMarkdown retrieves SummaryMarkdown of Change entity
func (c Change) SummaryMarkdown() *string {
	return c.change.SummaryMarkdown
}

// ReleasedAt retrieves ReleasedAt of Change entity
func (c Change) ReleasedAt() *scalar.Time {
	if c.change.ReleasedAt == nil {
		return nil
	}
	return &scalar.Time{Time: *c.change.ReleasedAt}
}

func newChange(change entity.Change) Change {
	return Change{change: change}
}

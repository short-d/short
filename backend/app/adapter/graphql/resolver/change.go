package resolver

import (
	"github.com/short-d/short/app/adapter/graphql/scalar"
	"github.com/short-d/short/app/entity"
)

type Change struct {
	change entity.Change
}

func (c Change) ID() string {
	return c.change.ID
}

func (c Change) Title() *string {
	return &c.change.Title
}

func (c Change) SummaryMarkdown() *string {
	return &c.change.SummaryMarkdown
}

func (c Change) ReleasedAt() scalar.Time {
	return scalar.Time{Time: *c.change.ReleasedAt}
}

func newChange(change entity.Change) Change {
	return Change{change: change}
}

package resolver

import (
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
)

// Change retrieves requested fields of a Change.
type Change struct {
	change entity.Change
}

// ID retrieves the ID of Change entity.
func (c Change) ID() string {
	return c.change.ID
}

// Title retrieves the title of Change entity.
func (c Change) Title() string {
	return c.change.Title
}

// SummaryMarkdown retrieves the summary of Change entity in Markdown format.
func (c Change) SummaryMarkdown() *string {
	return c.change.SummaryMarkdown
}

// ReleasedAt retrieves ReleasedAt of Change entity
func (c Change) ReleasedAt() scalar.Time {
	return scalar.Time{Time: c.change.ReleasedAt}
}

func newChange(change entity.Change) Change {
	return Change{change: change}
}

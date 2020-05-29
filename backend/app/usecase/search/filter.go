package search

// put those in entity package?
type Resource uint

const (
	ShortLink Resource = iota
	User
)

type Order uint

const (
	CreatedTimeASC Order = iota
)

type Filter struct {
	maxResults int
	resources  []Resource
	orders     []Order
}

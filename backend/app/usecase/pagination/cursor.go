package pagination

type Cursor struct {
	Current map[string]string
	Prev    map[string]string
	Next    map[string]string
	Size    int
}

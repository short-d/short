package filesystem

// FileSystem enables API consumers to persist a consecutive piece of data on
// the block storage.
type FileSystem interface {
	ReadFile(filepath string) ([]byte, error)
}

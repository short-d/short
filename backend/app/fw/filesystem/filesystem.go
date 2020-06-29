package filesystem

type FileSystem interface {
	ReadFile(filepath string) ([]byte, error)
}

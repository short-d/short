package service

// Key represents unique identifier
type Key string

// KeyFetcher fetches keys in batch
type KeyFetcher interface {
	FetchKeys(maxCount int) ([]Key, error)
}

package service

import "github.com/short-d/kgs/app/entity"

// KeyFetcher fetches keys in batch
type KeyFetcher interface {
	FetchKeys(maxCount int) ([]entity.Key, error)
}

package service

import "github.com/byliuyang/kgs/app/entity"

// KeyFetcher fetches keys in batch
type KeyFetcher interface {
	FetchKeys(maxCount int) ([]entity.Key, error)
}

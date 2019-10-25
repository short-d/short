package service

import "github.com/byliuyang/kgs/app/entity"

type KeyGen interface {
	FetchKeys(maxCount int) ([]entity.Key, error)
}

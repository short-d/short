package service

import "short/app/entity"

type Profile interface {
	GetUserProfile(accessToken string) (entity.UserProfile, error)
}

package service

import "short/app/entity"

var _ Profile = (*ProfileFake)(nil)

type ProfileFake struct {
	userProfile entity.UserProfile
}

func (p ProfileFake) GetUserProfile(accessToken string) (entity.UserProfile, error) {
	return p.userProfile, nil
}

func NewProfileFake(userProfile entity.UserProfile) ProfileFake {
	return ProfileFake{
		userProfile: userProfile,
	}
}

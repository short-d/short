package service

import "short/app/entity"

type SSOAccount interface {
	GetSingleSignOnUser(accessToken string) (entity.SSOUser, error)
}

package order

import (
	"github.com/short-d/short/backend/app/entity"
)

var _ Order = (*CreatedTime)(nil)

type CreatedTime struct {
}

func (c CreatedTime) ArrangeShortLinks(shortLinks []entity.ShortLink) []entity.ShortLink {
	panic("implement me")
}

func (c CreatedTime) ArrangeUsers(users []entity.User) []entity.User {
	panic("implement me")
}

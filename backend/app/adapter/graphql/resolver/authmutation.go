package resolver

import (
	"errors"
	"short/app/entity"
	"short/app/usecase/url"
)

type AuthMutation struct {
	user       *entity.User
	urlCreator url.Creator
}

type CreateURLArgs struct {
	CaptchaResponse string
	URL             URLInput
	AuthToken       string
}

func (a AuthMutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	if a.user == nil {
		return nil, errors.New("unauthorized request")
	}

	customAlias := args.URL.CustomAlias
	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	newURL, err := a.urlCreator.CreateURL(u, customAlias, *a.user)
	if err == nil {
		return &URL{url: newURL}, nil
	}

	switch err.(type) {
	case url.ErrAliasExist:
		return nil, ErrURLAliasExist(*customAlias)
	case url.ErrInvalidLongLink:
		return nil, ErrInvalidLongLink(u.OriginalURL)
	case url.ErrInvalidCustomAlias:
		return nil, ErrInvalidCustomAlias(*customAlias)
	default:
		return nil, ErrUnknown{}
	}
}

func newAuthMutation(user *entity.User, urlCreator url.Creator) AuthMutation {
	return AuthMutation{
		user:       user,
		urlCreator: urlCreator,
	}
}

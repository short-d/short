package resolver

import (
	"errors"
	"short/app/entity"
	"short/app/usecase/url"
	"time"
)

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	user       *entity.User
	urlCreator url.Creator
}

// URLInput represents possible URL attributes
type URLInput struct {
	OriginalURL string
	CustomAlias *string
	ExpireAt    *time.Time
}

// CreateURLArgs represents the possible parameters for CreateURL endpoint
type CreateURLArgs struct {
	URL      URLInput
	IsPublic bool
}

// CreateURL creates mapping between an alias and a long link mapping for a
// given user
func (a AuthMutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	if a.user == nil {
		return nil, errors.New("unauthorized request")
	}

	customAlias := args.URL.CustomAlias
	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	isPublic := args.IsPublic

	newURL, err := a.urlCreator.CreateURL(u, customAlias, *a.user, isPublic)
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

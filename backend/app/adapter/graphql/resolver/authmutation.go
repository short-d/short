package resolver

import (
	"errors"
	"time"

	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/url"
)

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	user             *entity.User
	changeLogCreator changelog.Creator
	urlCreator       url.Creator
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

type CreateChangeArgs struct {
	change ChangeInput
}

type ChangeInput struct {
	title           *string
	summaryMarkdown *string
}

// CreateURL creates mapping between an alias and a long link for a given user
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

func (a AuthMutation) CreateChange(args *CreateChangeArgs) (Change, error) {
	change, err := a.changeLogCreator.CreateChange("1234", *args.change.title, *args.change.summaryMarkdown)
	return *newChange(change), err
}

func newAuthMutation(user *entity.User, changeLogCreator changelog.Creator, urlCreator url.Creator) AuthMutation {
	return AuthMutation{
		user:             user,
		changeLogCreator: changeLogCreator,
		urlCreator:       urlCreator,
	}
}

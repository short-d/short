package resolver

import (
	"time"

	"github.com/short-d/short/app/usecase/authenticator"

	"github.com/short-d/short/app/entity"

	"github.com/short-d/short/app/usecase/changelog"
	"github.com/short-d/short/app/usecase/url"
)

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	authToken     *string
	authenticator authenticator.Authenticator
	changeLog     changelog.ChangeLog
	urlCreator    url.Creator
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

// CreateChangeArgs represents the possible parameters for CreateChange endpoint
type CreateChangeArgs struct {
	Change ChangeInput
}

// ChangeInput represents possible properties for Change
type ChangeInput struct {
	Title           string
	SummaryMarkdown *string
}

// CreateURL creates mapping between an alias and a long link for a given user
func (a AuthMutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	customAlias := args.URL.CustomAlias
	u := entity.URL{
		OriginalURL: args.URL.OriginalURL,
		ExpireAt:    args.URL.ExpireAt,
	}

	isPublic := args.IsPublic

	newURL, err := a.urlCreator.CreateURL(u, customAlias, user, isPublic)
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

// CreateChange creates a Change in the change log
func (a AuthMutation) CreateChange(args *CreateChangeArgs) (Change, error) {
	change, err := a.changeLog.CreateChange(args.Change.Title, args.Change.SummaryMarkdown)
	return newChange(change), err
}

func newAuthMutation(
	authToken *string,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	urlCreator url.Creator,
) AuthMutation {
	return AuthMutation{
		authToken:     authToken,
		authenticator: authenticator,
		changeLog:     changeLog,
		urlCreator:    urlCreator,
	}
}

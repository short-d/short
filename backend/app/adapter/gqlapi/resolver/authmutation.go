package resolver

import (
	"errors"
	"time"

	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/url"
)

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	authToken     *string
	authenticator authenticator.Authenticator
	changeLog     changelog.ChangeLog
	urlCreator    url.Creator
	urlUpdater    url.Updater
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
	case url.ErrMaliciousLongLink:
		return nil, ErrMaliciousContent(u.OriginalURL)
	default:
		return nil, ErrUnknown{}
	}
}

type UpdateURLArgs struct {
	OldAlias string
	Url      URLUpdateInput
	IsPublic bool
}

type URLUpdateInput struct {
	OriginalURL *string
	CustomAlias *string
	ExpireAt    *time.Time
}

func (u URLUpdateInput) isValid() bool {
	return u != URLUpdateInput{}
}

func (a AuthMutation) UpdateURL(args *UpdateURLArgs) (*URL, error) {
	originalURL := args.Url.OriginalURL
	customAlias := args.Url.CustomAlias
	expireAt := args.Url.ExpireAt

	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	if !args.Url.isValid() {
		return nil, errors.New("Empty Update")
	}

	update := &entity.URL{
		ExpireAt: expireAt,
	}

	if originalURL != nil {
		update.OriginalURL = *originalURL
	}

	if customAlias != nil {
		update.Alias = *customAlias
	}

	newURL, err := a.urlUpdater.UpdateURL(args.OldAlias, *update, user)
	if err != nil {
		return nil, err
	}

	return &URL{url: newURL}, nil
}

// CreateChange creates a Change in the change log
func (a AuthMutation) CreateChange(args *CreateChangeArgs) (Change, error) {
	change, err := a.changeLog.CreateChange(args.Change.Title, args.Change.SummaryMarkdown)
	return newChange(change), err
}

// ViewChangeLog records the time when the user viewed the change log
func (a AuthMutation) ViewChangeLog() (scalar.Time, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return scalar.Time{}, ErrInvalidAuthToken{}
	}

	lastViewedAt, err := a.changeLog.ViewChangeLog(user)
	return scalar.Time{Time: lastViewedAt}, err
}

func newAuthMutation(
	authToken *string,
	authenticator authenticator.Authenticator,
	changeLog changelog.ChangeLog,
	urlCreator url.Creator,
	urlUpdater url.Updater,
) AuthMutation {
	return AuthMutation{
		authToken:     authToken,
		authenticator: authenticator,
		changeLog:     changeLog,
		urlCreator:    urlCreator,
		urlUpdater:    urlUpdater,
	}
}

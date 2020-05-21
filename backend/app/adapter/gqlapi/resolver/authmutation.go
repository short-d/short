package resolver

import (
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
	OriginalURL *string
	CustomAlias *string
	ExpireAt    *time.Time
}

// isEmpty checks if the input contains only nil pointers
func (u *URLInput) isEmpty() bool {
	return *u == URLInput{}
}

// originalURL returns the URLInput OrignalURL and an empty string if the
// pointer references a nil value.
func (u *URLInput) originalURL() string {
	if u.OriginalURL == nil {
		return ""
	}
	return *u.OriginalURL
}

// customAlias returns the URLInput CustomAlias and an empty string if the
// pointer references a nil value.
func (u *URLInput) customAlias() string {
	if u.CustomAlias == nil {
		return ""
	}
	return *u.CustomAlias
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

	originalURL := args.URL.originalURL()
	customAlias := args.URL.CustomAlias
	u := entity.URL{
		OriginalURL: originalURL,
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

// UpdateURLArgs represents the possible parameters for updateURL endpoint
type UpdateURLArgs struct {
	OldAlias string
	URL      URLInput
}

func (u *URLInput) createUpdate() (*entity.URL, error) {
	if u.isEmpty() {
		return nil, ErrEmptyUpdate{}
	}

	return &entity.URL{
		Alias:       u.customAlias(),
		OriginalURL: u.originalURL(),
		ExpireAt:    u.ExpireAt,
	}, nil

}

// UpdateURL updates a short link mapping that belongs to a user
func (a AuthMutation) UpdateURL(args *UpdateURLArgs) (*URL, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	update, err := args.URL.createUpdate()
	if err != nil {
		return nil, err
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

package resolver

import (
	"time"

	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/authenticator"
	"github.com/short-d/short/backend/app/usecase/changelog"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

// AuthMutation represents GraphQL mutation resolver that acts differently based
// on the identify of the user
type AuthMutation struct {
	authToken        *string
	authenticator    authenticator.Authenticator
	changeLog        changelog.ChangeLog
	shortLinkCreator shortlink.Creator
	shortLinkUpdater shortlink.Updater
}

// URLInput represents possible ShortLink attributes
type URLInput struct {
	OriginalURL *string
	CustomAlias *string
	ExpireAt    *time.Time
}

func (u *URLInput) isEmpty() bool {
	return *u == URLInput{}
}

func (u *URLInput) originalURL() string {
	if u.OriginalURL == nil {
		return ""
	}
	return *u.OriginalURL
}

func (u *URLInput) customAlias() string {
	if u.CustomAlias == nil {
		return ""
	}
	return *u.CustomAlias
}

func (u *URLInput) createUpdate() *entity.ShortLink {
	if u.isEmpty() {
		return nil
	}

	return &entity.ShortLink{
		Alias:    u.customAlias(),
		LongLink: u.originalURL(),
		ExpireAt: u.ExpireAt,
	}

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

// CreateShortLink creates mapping between an alias and a long link for a given user
func (a AuthMutation) CreateURL(args *CreateURLArgs) (*URL, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	originalURL := args.URL.originalURL()
	customAlias := args.URL.CustomAlias
	u := entity.ShortLink{
		LongLink: originalURL,
		ExpireAt: args.URL.ExpireAt,
	}

	isPublic := args.IsPublic

	newShortLink, err := a.shortLinkCreator.CreateShortLink(u, customAlias, user, isPublic)
	if err == nil {
		return &URL{url: newShortLink}, nil
	}

	switch err.(type) {
	case shortlink.ErrAliasExist:
		return nil, ErrAliasExist(*customAlias)
	case shortlink.ErrInvalidLongLink:
		return nil, ErrInvalidLongLink(u.LongLink)
	case shortlink.ErrInvalidCustomAlias:
		return nil, ErrInvalidCustomAlias(*customAlias)
	case shortlink.ErrMaliciousLongLink:
		return nil, ErrMaliciousContent(u.LongLink)
	default:
		return nil, ErrUnknown{}
	}
}

// UpdateURLArgs represents the possible parameters for updateURL endpoint
type UpdateURLArgs struct {
	OldAlias string
	URL      URLInput
}

// UpdateURL updates the relationship between the short link and the user
func (a AuthMutation) UpdateURL(args *UpdateURLArgs) (*URL, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	update := args.URL.createUpdate()
	if update == nil {
		return nil, nil
	}

	newURL, err := a.shortLinkUpdater.UpdateURL(args.OldAlias, *update, user)
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
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
) AuthMutation {
	return AuthMutation{
		authToken:        authToken,
		authenticator:    authenticator,
		changeLog:        changeLog,
		shortLinkCreator: shortLinkCreator,
		shortLinkUpdater: shortLinkUpdater,
	}
}

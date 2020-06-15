package resolver

import (
	"fmt"
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
}

// ShortLinkInput represents possible ShortLink attributes
type ShortLinkInput struct {
	LongLink    string
	CustomAlias *string
	ExpireAt    *time.Time
}

// CreateShortLinkArgs represents the possible parameters for CreateShortLink endpoint
type CreateShortLinkArgs struct {
	ShortLink ShortLinkInput
	IsPublic  bool
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

// DeleteChangeArgs represents the possible parameters for DeleteChange endpoint
type DeleteChangeArgs struct {
	ID string
}

// UpdateChangeArgs represents the possible parameters for UpdateChange endpoint.
type UpdateChangeArgs struct {
	ID     string
	Change ChangeInput
}

// CreateShortLink creates mapping between an alias and a long link for a given user
func (a AuthMutation) CreateShortLink(args *CreateShortLinkArgs) (*ShortLink, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	customAlias := args.ShortLink.CustomAlias
	u := entity.ShortLink{
		LongLink: args.ShortLink.LongLink,
		ExpireAt: args.ShortLink.ExpireAt,
	}

	isPublic := args.IsPublic

	newShortLink, err := a.shortLinkCreator.CreateShortLink(u, customAlias, user, isPublic)
	if err == nil {
		return &ShortLink{shortLink: newShortLink}, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case shortlink.ErrAliasExist:
		return nil, ErrAliasExist(*customAlias)
	case shortlink.ErrInvalidLongLink:
		return nil, ErrInvalidLongLink{u.LongLink, string(err.(shortlink.ErrInvalidLongLink).Violation)}
	case shortlink.ErrInvalidCustomAlias:
		return nil, ErrInvalidCustomAlias{*customAlias, string(err.(shortlink.ErrInvalidCustomAlias).Violation)}
	case shortlink.ErrMaliciousLongLink:
		return nil, ErrMaliciousContent(u.LongLink)
	default:
		return nil, ErrUnknown{}
	}
}

// CreateChange creates a Change in the change log
func (a AuthMutation) CreateChange(args *CreateChangeArgs) (*Change, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	change, err := a.changeLog.CreateChange(args.Change.Title, args.Change.SummaryMarkdown, user)
	if err == nil {
		change := newChange(change)
		return &change, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case changelog.ErrUnauthorizedAction:
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to create a change", user.ID))
	default:
		return nil, ErrUnknown{}
	}
}

// DeleteChange removes a Change with given ID from change log
func (a AuthMutation) DeleteChange(args *DeleteChangeArgs) (*string, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	err = a.changeLog.DeleteChange(args.ID, user)
	if err == nil {
		return &args.ID, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case changelog.ErrUnauthorizedAction:
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to delete the change %s", user.ID, args.ID))
	default:
		return nil, ErrUnknown{}
	}
}

// UpdateChange updates a Change with given ID in change log.
func (a AuthMutation) UpdateChange(args *UpdateChangeArgs) (*Change, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	change, err := a.changeLog.UpdateChange(
		args.ID,
		args.Change.Title,
		args.Change.SummaryMarkdown,
		user,
	)
	if err == nil {
		change := newChange(change)
		return &change, nil
	}

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case changelog.ErrUnauthorizedAction:
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to update the change %s", user.ID, args.ID))
	default:
		return nil, ErrUnknown{}
	}
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
) AuthMutation {
	return AuthMutation{
		authToken:        authToken,
		authenticator:    authenticator,
		changeLog:        changeLog,
		shortLinkCreator: shortLinkCreator,
	}
}

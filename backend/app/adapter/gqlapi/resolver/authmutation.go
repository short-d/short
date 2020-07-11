package resolver

import (
	"errors"
	"fmt"

	"github.com/short-d/short/backend/app/adapter/gqlapi/args"
	"github.com/short-d/short/backend/app/adapter/gqlapi/scalar"
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

// CreateShortLinkArgs represents the possible parameters for CreateShortLink endpoint
type CreateShortLinkArgs struct {
	ShortLink args.ShortLinkInput
	IsPublic  bool
}

// CreateShortLink creates mapping between an alias and a long link for a given user
func (a AuthMutation) CreateShortLink(args *CreateShortLinkArgs) (*ShortLink, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	shortLink := args.ShortLink.CreateShortLinkInput()
	isPublic := args.IsPublic

	newShortLink, err := a.shortLinkCreator.CreateShortLink(shortLink, user, isPublic)
	if err == nil {
		return &ShortLink{shortLink: newShortLink}, nil
	}

	var (
		ae shortlink.ErrAliasExist
		l  shortlink.ErrInvalidLongLink
		c  shortlink.ErrInvalidCustomAlias
		m  shortlink.ErrMaliciousLongLink
	)
	if errors.As(err, &ae) {
		return nil, ErrAliasExist(shortLink.CustomAlias)
	}
	if errors.As(err, &l) {
		return nil, ErrInvalidLongLink{shortLink.LongLink, string(l.Violation)}
	}
	if errors.As(err, &c) {
		return nil, ErrInvalidCustomAlias{shortLink.CustomAlias, string(c.Violation)}
	}
	if errors.As(err, &m) {
		return nil, ErrMaliciousContent(shortLink.LongLink)
	}
	return nil, ErrUnknown{}
}

// UpdateShortLinkArgs represents the possible parameters for updateShortLink endpoint
type UpdateShortLinkArgs struct {
	OldAlias  string
	ShortLink args.ShortLinkInput
}

// UpdateShortLink updates the relationship between the short link and the user
func (a AuthMutation) UpdateShortLink(args *UpdateShortLinkArgs) (*ShortLink, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	update := args.ShortLink.CreateShortLinkInput()

	newShortLink, err := a.shortLinkUpdater.UpdateShortLink(args.OldAlias, update, user)
	if err == nil {
		return &ShortLink{shortLink: newShortLink}, nil
	}

	var (
		ae shortlink.ErrAliasExist
		l  shortlink.ErrInvalidLongLink
		c  shortlink.ErrInvalidCustomAlias
		m  shortlink.ErrMaliciousLongLink
		nf shortlink.ErrShortLinkNotFound
	)
	if errors.As(err, &ae) {
		return nil, ErrAliasExist(update.CustomAlias)
	}
	if errors.As(err, &l) {
		return nil, ErrInvalidLongLink{update.LongLink, string(l.Violation)}
	}
	if errors.As(err, &c) {
		return nil, ErrInvalidCustomAlias{update.CustomAlias, string(c.Violation)}
	}
	if errors.As(err, &m) {
		return nil, ErrMaliciousContent(update.LongLink)
	}
	if errors.As(err, &nf) {
		return nil, ErrShortLinkNotFound(update.CustomAlias)
	}
	return nil, ErrUnknown{}
}

// ChangeInput represents possible properties for Change
type ChangeInput struct {
	Title           string
	SummaryMarkdown *string
}

// CreateChangeArgs represents the possible parameters for CreateChange endpoint
type CreateChangeArgs struct {
	Change ChangeInput
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

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to create a change", user.ID))
	}
	return nil, ErrUnknown{}
}

// DeleteChangeArgs represents the possible parameters for DeleteChange endpoint
type DeleteChangeArgs struct {
	ID string
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

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to delete the change %s", user.ID, args.ID))
	}
	return nil, ErrUnknown{}
}

// UpdateChangeArgs represents the possible parameters for UpdateChange endpoint.
type UpdateChangeArgs struct {
	ID     string
	Change ChangeInput
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

	var (
		u changelog.ErrUnauthorizedAction
	)
	if errors.As(err, &u) {
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to update the change %s", user.ID, args.ID))
	}
	return nil, ErrUnknown{}
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

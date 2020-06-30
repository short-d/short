package resolver

import (
	"fmt"
	"github.com/short-d/short/backend/app/usecase/emotic"
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
	apiKey           *string
	authenticator    authenticator.Authenticator
	cloudApiAuth     authenticator.CloudAPI
	changeLog        changelog.ChangeLog
	shortLinkCreator shortlink.Creator
	shortLinkUpdater shortlink.Updater
	feedback         emotic.Feedback
}

type GenerateApiKeyArgs struct {
	AppID string
}

// ViewChangeLog records the time when the user viewed the change log
func (a AuthMutation) GenerateApiKey(args GenerateApiKeyArgs) (string, error) {
	//user, err := viewer(a.authToken, a.authenticator)
	//if err != nil {
	//	return "", ErrInvalidAuthToken{}
	//}

	return a.cloudApiAuth.GenerateApiKey(args.AppID)
}

// ShortLinkInput represents possible ShortLink attributes
type ShortLinkInput struct {
	LongLink    *string
	CustomAlias *string
	ExpireAt    *time.Time
}

// TODO(#840): remove this business logic and move it to use cases
func (s *ShortLinkInput) isEmpty() bool {
	return *s == ShortLinkInput{}
}

// TODO(#840): remove this business logic and move it to use cases
func (s *ShortLinkInput) longLink() string {
	if s.LongLink == nil {
		return ""
	}
	return *s.LongLink
}

// TODO(#840): remove this business logic and move it to use cases
func (s *ShortLinkInput) customAlias() string {
	if s.CustomAlias == nil {
		return ""
	}
	return *s.CustomAlias
}

// TODO(#840): remove this business logic and move it to use cases
func (s *ShortLinkInput) createUpdate() *entity.ShortLink {
	if s.isEmpty() {
		return nil
	}

	return &entity.ShortLink{
		Alias:    s.customAlias(),
		LongLink: s.longLink(),
		ExpireAt: s.ExpireAt,
	}

}

// CreateShortLinkArgs represents the possible parameters for CreateShortLink endpoint
type CreateShortLinkArgs struct {
	ShortLink ShortLinkInput
	IsPublic  bool
}

// CreateShortLink creates mapping between an alias and a long link for a given user
func (a AuthMutation) CreateShortLink(args *CreateShortLinkArgs) (*ShortLink, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	longLink := args.ShortLink.longLink()
	customAlias := args.ShortLink.CustomAlias
	u := entity.ShortLink{
		LongLink: longLink,
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

// UpdateShortLinkArgs represents the possible parameters for updateShortLink endpoint
type UpdateShortLinkArgs struct {
	OldAlias  string
	ShortLink ShortLinkInput
}

// UpdateShortLink updates the relationship between the short link and the user
func (a AuthMutation) UpdateShortLink(args *UpdateShortLinkArgs) (*ShortLink, error) {
	user, err := viewer(a.authToken, a.authenticator)
	if err != nil {
		return nil, ErrInvalidAuthToken{}
	}

	update := args.ShortLink.createUpdate()
	if update == nil {
		return nil, nil
	}

	newShortLink, err := a.shortLinkUpdater.UpdateShortLink(args.OldAlias, *update, user)
	if err != nil {
		return nil, err
	}

	return &ShortLink{shortLink: newShortLink}, nil
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

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case changelog.ErrUnauthorizedAction:
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to create a change", user.ID))
	default:
		return nil, ErrUnknown{}
	}
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

	// TODO(issue#823): refactor error type checking
	switch err.(type) {
	case changelog.ErrUnauthorizedAction:
		return nil, ErrUnauthorizedAction(fmt.Sprintf("user %s is not allowed to delete the change %s", user.ID, args.ID))
	default:
		return nil, ErrUnknown{}
	}
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

type FeedbackInput struct {
	CustomerRating *float64
	Comment        *string
	CustomerEmail  *string
}

type ReceiveFeedbackArgs struct {
	Feedback FeedbackInput
}

func (a AuthMutation) ReceiveFeedback(args ReceiveFeedbackArgs) (Feedback, error) {
	app, err := app(a.apiKey, a.cloudApiAuth)
	if err != nil {
		return Feedback{}, ErrInvalidAuthToken{}
	}

	var rating *int
	if args.Feedback.CustomerRating != nil {
		num := int(*args.Feedback.CustomerRating)
		rating = &num
	}
	feedbackInput := entity.FeedbackInput{
		AppID:          &app.ID,
		CustomerRating: rating,
		Comment:        args.Feedback.Comment,
		CustomerEmail:  args.Feedback.CustomerEmail,
	}
	fb, err := a.feedback.ReceiveFeedback(feedbackInput)
	if err != nil {
		return Feedback{}, err
	}
	return Feedback{
		feedback: fb,
	}, nil
}

func newAuthMutation(
	authToken *string,
	apiKey *string,
	authenticator authenticator.Authenticator,
	cloudApiAuth authenticator.CloudAPI,
	changeLog changelog.ChangeLog,
	shortLinkCreator shortlink.Creator,
	shortLinkUpdater shortlink.Updater,
	feedback emotic.Feedback,
) AuthMutation {
	return AuthMutation{
		authToken:        authToken,
		apiKey:           apiKey,
		authenticator:    authenticator,
		cloudApiAuth:     cloudApiAuth,
		changeLog:        changeLog,
		shortLinkCreator: shortLinkCreator,
		shortLinkUpdater: shortLinkUpdater,
		feedback:         feedback,
	}
}

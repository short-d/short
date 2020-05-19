package feature

import (
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/usecase/instrumentation"
)

const (
	ChangeLog string = "change-log"
	FacebookSignIn string = "facebook-sign-in"
	GithubSignIn string = "github-sign-in"
	GoogleSignIn string = "google-sign-in"
	SearchBar string = "search-bar"
	UserShortLinksSection string = "user-short-links-section"
	PreferenceToggles string = "preference-toggles"
	IncludeAdminPanel string = "include-admin-panel"
)

type PermissionChecker func(user entity.User) (bool, error)

// DecisionMaker determines whether a feature should be turned on or off under
// certain conditions.
type DecisionMaker interface {
	IsFeatureEnable(featureID string, user *entity.User) bool
}

// DecisionMakerFactory creates feature decision maker.
type DecisionMakerFactory interface {
	NewDecision(instrumentation instrumentation.Instrumentation) DecisionMaker
}

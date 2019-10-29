package facebook

import (
	"short/app/entity"

	fb "github.com/huandu/facebook"
)

// API accesses user's profile information through FB Graph API.
type API struct {
}

// GetUserProfile retrieves user's email and name from Facebook.
func (g API) GetUserProfile(accessToken string) (entity.UserProfile, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	var fbResponse response

	res, err := fb.Get("/me", fb.Params{
		"fields":       "name,email",
		"access_token": accessToken,
	})

	if err != nil {
		return entity.UserProfile{}, err
	}

	res.Decode(&fbResponse)

	return entity.UserProfile{
		Email: fbResponse.Email,
		Name:  fbResponse.Name,
	}, nil
}

// NewAPI initializes Github API access service.
func NewAPI() API {
	return API{}
}

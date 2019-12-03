package google

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"short/app/entity"
	"short/app/usecase/service"
	"strings"
)

var _ service.SSOAccount = (*Account)(nil)

// Account accesses user's account data through FB Graph API.
type Account struct {
}

// GetSingleSignOnUser retrieves user's email and name from Facebook API.
func (g Account) GetSingleSignOnUser(accessToken string) (entity.SSOUser, error) {
	type response struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	form := url.Values{}

	req, err := http.NewRequest("GET", "https://www.googleapis.com/drive/v2/files", strings.NewReader(form.Encode()))
	req.PostForm = form
	req.Header.Add("Authorization", "Bearer " + accessToken)
	if err != nil {
		return entity.SSOUser{}, err
	}

	hc := http.Client{}
	resp, err := hc.Do(req)
	if err != nil {
		return entity.SSOUser{}, err
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return entity.SSOUser{}, err
	}
	var result response
	if err := json.Unmarshal(data, &result); err != nil {
		return entity.SSOUser{}, err
	}

	return entity.SSOUser{
		Email: result.Email,
		Name:  result.Name,
	}, nil
}

// NewAccount initializes Facebook account API client.
func NewAccount() Account {
	return Account{}
}

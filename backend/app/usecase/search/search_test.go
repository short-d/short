package search

import (
	"testing"

	"github.com/short-d/app/mdtest"
	"github.com/short-d/short/app/entity"
	"github.com/short-d/short/app/usecase/repository"
)

func TestSearch_SearchForURLs(t *testing.T) {
	user_1 := entity.User{
		ID:    "1",
		Name:  "user_1",
		Email: "user_1@gmail.com",
	}
	user_2 := entity.User{
		ID:    "2",
		Name:  "user_2",
		Email: "user_2@gmail.com",
	}
	url_1 := entity.URL{
		Alias:       "baidu",
		OriginalURL: "http://www.baidu.com",
		CreatedBy:   &user_1,
	}
	url_2 := entity.URL{
		Alias:       "gg",
		OriginalURL: "http://www.google.com",
		CreatedBy:   &user_1,
	}
	urlMap := make(map[string]entity.URL)
	urlMap["baidu"] = url_1
	urlMap["gg"] = url_2
	urlRepo := repository.NewURLFake(urlMap)
	userUrlRepo := repository.NewUserURLRepoFake([]entity.User{user_1, user_1}, []entity.URL{url_1, url_2})

	testCases := []struct {
		name         string
		user         entity.User
		expectedURLs []entity.URL
		expectedErr  error
	}{
		{
			name: "success",
			user: user_1,
			expectedURLs: []entity.URL{
				url_1,
				url_2,
			},
			expectedErr: nil,
		},
		{
			name:         "no such user",
			user:         user_2,
			expectedURLs: []entity.URL{},
			expectedErr:  nil,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			searchAPI := NewSearch(&urlRepo, &userUrlRepo)

			urls, err := searchAPI.SearchForURLs(testCase.user)
			mdtest.Equal(t, testCase.expectedURLs, urls)
			mdtest.Equal(t, testCase.expectedErr, err)
		})
	}
}

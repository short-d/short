// +build !integration all

package shortlink

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestMetaTagPersist_GetOpenGraphTags(t *testing.T) {
	t.Parallel()

	title1 := "title1"
	description1 := "description1"
	imageURL1 := "imageURL1"

	title2 := "title2"
	description2 := "description2"
	imageURL2 := "imageURL2"

	title3 := "title3"
	description3 := "description3"
	imageURL3 := "imageURL3"

	defaultTitle := defaultTitle
	defaultDesc := defaultDesc
	defaultImageURL := defaultImageURL

	testCases := []struct {
		name                  string
		shortLinks            shortLinks
		alias                 string
		expHasErr             bool
		expectedOpenGraphTags metatag.OpenGraph
	}{
		{
			name: "Open Graph tags provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title3,
						Description: &description3,
						ImageURL:    &imageURL3,
					},
				},
			},
			alias: "54321",
			expectedOpenGraphTags: metatag.OpenGraph{
				Title:       &title2,
				Description: &description2,
				ImageURL:    &imageURL2,
			},
		},
		{

			name: "Open Graph tags not provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					TwitterTags: metatag.Twitter{
						Title:       &title3,
						Description: &description3,
						ImageURL:    &imageURL3,
					},
				},
			},
			alias: "54321",
			expectedOpenGraphTags: metatag.OpenGraph{
				Title:       &defaultTitle,
				Description: &defaultDesc,
				ImageURL:    &defaultImageURL,
			},
		},
		{
			name:                  "alias does not exist",
			shortLinks:            shortLinks{},
			alias:                 "654321",
			expHasErr:             true,
			expectedOpenGraphTags: metatag.OpenGraph{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks, nil)
			metaTag := NewMetaTagPersist(&shortLinkRepo)

			ogTags, err := metaTag.GetOpenGraphTags(testCase.alias)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedOpenGraphTags, ogTags)
		})
	}
}

func TestMetaTagPersist_GetTwitterTags(t *testing.T) {
	t.Parallel()

	title1 := "title1"
	description1 := "description1"
	imageURL1 := "imageURL1"

	title2 := "title2"
	description2 := "description2"
	imageURL2 := "imageURL2"

	title3 := "title3"
	description3 := "description3"
	imageURL3 := "imageURL3"

	defaultTitle := defaultTitle
	defaultDesc := defaultDesc
	defaultImageURL := defaultImageURL

	testCases := []struct {
		name                string
		shortLinks          shortLinks
		alias               string
		expHasErr           bool
		expectedTwitterTags metatag.Twitter
	}{
		{
			name: "Twitter tags provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title3,
						Description: &description3,
						ImageURL:    &imageURL3,
					},
				},
			},
			alias: "54321",
			expectedTwitterTags: metatag.Twitter{
				Title:       &title3,
				Description: &description3,
				ImageURL:    &imageURL3,
			},
		},
		{
			name: "Twitter tags not provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title1,
						Description: &description1,
						ImageURL:    &imageURL1,
					},
					TwitterTags: metatag.Twitter{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       &title2,
						Description: &description2,
						ImageURL:    &imageURL2,
					},
				},
			},
			alias: "54321",
			expectedTwitterTags: metatag.Twitter{
				Title:       &defaultTitle,
				Description: &defaultDesc,
				ImageURL:    &defaultImageURL,
			},
		},
		{
			name:                "alias does not exist",
			shortLinks:          shortLinks{},
			alias:               "654321",
			expHasErr:           true,
			expectedTwitterTags: metatag.Twitter{},
		},
	}

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			shortLinkRepo := repository.NewShortLinkFake(testCase.shortLinks, nil)
			metaTag := NewMetaTagPersist(&shortLinkRepo)

			twitterTags, err := metaTag.GetTwitterTags(testCase.alias)
			if testCase.expHasErr {
				assert.NotEqual(t, nil, err)
				return
			}

			assert.Equal(t, nil, err)
			assert.Equal(t, testCase.expectedTwitterTags, twitterTags)
		})
	}
}

// +build !integration all

package shortlink

import (
	"testing"

	"github.com/short-d/app/fw/assert"
	"github.com/short-d/short/backend/app/entity"
	"github.com/short-d/short/backend/app/entity/metatag"
	"github.com/short-d/short/backend/app/fw/ptr"
	"github.com/short-d/short/backend/app/usecase/repository"
)

func TestMetaTagPersist_GetOpenGraphTags(t *testing.T) {
	t.Parallel()

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
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("imageURL1"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title3"),
						Description: ptr.String("description3"),
						ImageURL:    ptr.String("imageURL3"),
					},
				},
			},
			alias: "54321",
			expectedOpenGraphTags: metatag.OpenGraph{
				Title:       ptr.String("title2"),
				Description: ptr.String("description2"),
				ImageURL:    ptr.String("imageURL2"),
			},
		},
		{

			name: "Open Graph tags not provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("imageURL1"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title3"),
						Description: ptr.String("description3"),
						ImageURL:    ptr.String("imageURL3"),
					},
				},
			},
			alias: "54321",
			expectedOpenGraphTags: metatag.OpenGraph{
				Title:       ptr.String(defaultTitle),
				Description: ptr.String(defaultDesc),
				ImageURL:    ptr.String(defaultImageURL),
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

			shortLinkRepo := repository.NewShortLinkFake(nil, testCase.shortLinks)
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
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("imageURL1"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title3"),
						Description: ptr.String("description3"),
						ImageURL:    ptr.String("imageURL3"),
					},
				},
			},
			alias: "54321",
			expectedTwitterTags: metatag.Twitter{
				Title:       ptr.String("title3"),
				Description: ptr.String("description3"),
				ImageURL:    ptr.String("imageURL3"),
			},
		},
		{
			name: "Twitter tags not provided",
			shortLinks: shortLinks{
				"12345": entity.ShortLink{
					Alias:    "12345",
					LongLink: "www.google.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title1"),
						Description: ptr.String("description1"),
						ImageURL:    ptr.String("imageURL1"),
					},
					TwitterTags: metatag.Twitter{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
				},
				"54321": entity.ShortLink{
					Alias:    "54321",
					LongLink: "www.facebook.com",
					OpenGraphTags: metatag.OpenGraph{
						Title:       ptr.String("title2"),
						Description: ptr.String("description2"),
						ImageURL:    ptr.String("imageURL2"),
					},
				},
			},
			alias: "54321",
			expectedTwitterTags: metatag.Twitter{
				Title:       ptr.String(defaultTitle),
				Description: ptr.String(defaultDesc),
				ImageURL:    ptr.String(defaultImageURL),
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

			shortLinkRepo := repository.NewShortLinkFake(nil, testCase.shortLinks)
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

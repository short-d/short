package ssr

import (
	"io/ioutil"
	"path/filepath"

	"github.com/short-d/short/frontend/serve/entity"
)

// RedirectPage renders redirect page on the server side.
type RedirectPage struct {
	pageRootDir string
}

// Render renders redirect page with Open Graph & Twitter meta tags.
func (r RedirectPage) Render(openGraphTags entity.OpenGraphTags, twitterTags entity.TwitterTags) (string, error) {
	ssrVars := map[string]string{
		"OPEN_GRAPH_TITLE":       openGraphTags.Title,
		"OPEN_GRAPH_DESCRIPTION": openGraphTags.Description,
		"OPEN_GRAPH_IMAGE":       openGraphTags.ImageURL,
		"TWITTER_TITLE":          twitterTags.Title,
		"TWITTER_DESCRIPTION":    twitterTags.Description,
		"TWITTER_IMAGE":          twitterTags.ImageURL,
	}
	buf, err := ioutil.ReadFile(filepath.Join(r.pageRootDir, "index.html"))
	if err != nil {
		return "", err
	}

	page := string(buf)
	return renderPage(ssrVars, page), nil
}

// NewRedirectPage initializes RedirectPage
func NewRedirectPage(rootDir string) RedirectPage {
	return RedirectPage{pageRootDir: rootDir}
}

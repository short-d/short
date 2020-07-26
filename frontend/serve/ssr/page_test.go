package ssr

import (
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestRenderPage(t *testing.T) {
	testPage := `
jkas'x'n'j'kkj
{{SSR_OG_TITLE}}
nnm'zxc;nkkas
{{SSR_OG_DESCRIPTION}}
pqwepc['xzmkm
{{SSR_OG_IMAGE}}
llasdm21k.asda2
{{SSR_TWITTER_TITLE}}
023elksdlmcs;
{{SSR_TWITTER_DESCRIPTION}}
lzxckmzcxmkm23
{{SSR_TWITTER_IMAGE}}
[;sd-p32er2l,csdc
`
	ssrVars := map[string]string{
		"OG_TITLE":            "Title",
		"OG_DESCRIPTION":      "Description",
		"OG_IMAGE":            "ImageURL",
		"TWITTER_TITLE":       "Title",
		"TWITTER_DESCRIPTION": "Description",
		"TWITTER_IMAGE":       "ImageURL",
	}
	pageRetrieved := renderPage(ssrVars, testPage)
	expectedPage := `
jkas'x'n'j'kkj
Title
nnm'zxc;nkkas
Description
pqwepc['xzmkm
ImageURL
llasdm21k.asda2
Title
023elksdlmcs;
Description
lzxckmzcxmkm23
ImageURL
[;sd-p32er2l,csdc
`
	assert.Equal(t, pageRetrieved, expectedPage)
}

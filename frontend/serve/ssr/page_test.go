package ssr

import (
	"testing"

	"github.com/short-d/app/fw/assert"
)

func TestRenderPage(t *testing.T) {
	testPage := `
jkas'x'n'j'kkj
{{SSR_OPEN_GRAPH_TITLE}}
nnm'zxc;nkkas
{{SSR_OPEN_GRAPH_DESCRIPTION}}
pqwepc['xzmkm
{{SSR_OPEN_GRAPH_IMAGE}}
llasdm21k.asda2
{{SSR_TWITTER_TITLE}}
023elksdlmcs;
{{SSR_TWITTER_DESCRIPTION}}
lzxckmzcxmkm23
{{SSR_TWITTER_IMAGE}}
[;sd-p32er2l,csdc
`
	ssrVars := map[string]string{
		"OPEN_GRAPH_TITLE":       "OpenGraphTitle",
		"OPEN_GRAPH_DESCRIPTION": "OpenGraphDescription",
		"OPEN_GRAPH_IMAGE":       "OpenGraphImageURL",
		"TWITTER_TITLE":          "TwitterTitle",
		"TWITTER_DESCRIPTION":    "TwitterDescription",
		"TWITTER_IMAGE":          "TwitterImageURL",
	}
	pageGot := renderPage(ssrVars, testPage)
	expectedPage := `
jkas'x'n'j'kkj
OpenGraphTitle
nnm'zxc;nkkas
OpenGraphDescription
pqwepc['xzmkm
OpenGraphImageURL
llasdm21k.asda2
TwitterTitle
023elksdlmcs;
TwitterDescription
lzxckmzcxmkm23
TwitterImageURL
[;sd-p32er2l,csdc
`
	assert.Equal(t, pageGot, expectedPage)
}

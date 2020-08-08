package ssr

import (
	"fmt"
	"strings"
)

func renderPage(vars map[string]string, page string) string {
	for key, val := range vars {
		target := fmt.Sprintf("{{SSR_%s}}", key)
		page = strings.ReplaceAll(page, target, val)
	}
	return page
}

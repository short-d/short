package mdrouter

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type URIMatcher struct {
	pathFormat string
	pattern    *regexp.Regexp
	paramNames []string
}

var pathSep = "/"

func (m URIMatcher) IsMatch(path string) (bool, Params) {
	matches := m.pattern.FindStringSubmatch(path)

	if len(matches) < 1 {
		return false, Params{}
	}

	matches = matches[1:]

	if len(matches) != len(m.paramNames) {
		return false, Params{}
	}

	params := Params{}

	for idx, param := range matches {
		paramName := m.paramNames[idx]
		params[paramName] = param
	}
	return true, params
}

func (m URIMatcher) Params() []string {
	return m.paramNames
}

var paramPattern = regexp.MustCompile("^:([^/]+)$")

func extractParamName(paths []string) []string {
	var paramNames = make([]string, 0)

	for _, path := range paths {
		if isParam(path) {
			paramNames = append(paramNames, path[1:])
		}
	}

	return paramNames
}

func getUriPattern(pathFormat string, paths []string) *regexp.Regexp {
	var newPaths []string

	for _, path := range paths {
		if isParam(path) {
			newPaths = append(newPaths, "([^/]+)")
			continue
		}
		newPaths = append(newPaths, path)
	}

	patternText := strings.Join(newPaths, pathSep)
	patternText = fmt.Sprintf(pathFormat, patternText)
	return regexp.MustCompile(patternText)
}

func isParam(text string) bool {
	return paramPattern.MatchString(text)
}

func newURIMatcher(pathFormat string, uriTemplate string) (URIMatcher, error) {
	if len(uriTemplate) < 1 {
		return URIMatcher{}, errors.New("uri is empty")
	}

	if !strings.HasPrefix(uriTemplate, pathSep) {
		return URIMatcher{}, errors.New("uri has to start with /")
	}

	paths := strings.Split(uriTemplate, pathSep)

	paramNames := extractParamName(paths)
	uriPattern := getUriPattern(pathFormat, paths)

	return URIMatcher{
		pattern:    uriPattern,
		paramNames: paramNames,
	}, nil
}

func NewURIPrefixMatcher(uriTemplate string) (URIMatcher, error) {
	return newURIMatcher("^%s.*$", uriTemplate)
}

func NewURIExactMatcher(uriTemplate string) (URIMatcher, error) {
	return newURIMatcher("^%s/?$", uriTemplate)
}

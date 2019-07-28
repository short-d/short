package modern

import (
	"fmt"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

type Router struct {
}

type UriMatcher struct {
	pattern *regexp.Regexp
	paramNames  []string
}

type Params map[string]string

var pathSep = "/"

func (m UriMatcher) IsMatch(path string) (bool, Params) {
	matches := m.pattern.FindStringSubmatch(path)

	if len(matches) < 1 {
		return false, Params{}
	}

	matches = matches[1:]

	if len(matches) != len(m.paramNames) {
		return false, Params {}
	}

	params := Params{}

	for idx, param := range matches {
		paramName := m.paramNames[idx]
		params[paramName] = param
	}
	return true, params
}

func (m UriMatcher) Params() []string {
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

func getUriPattern(paths []string) *regexp.Regexp {
	var newPaths []string

	for _, path := range paths {
		if isParam(path) {
			newPaths = append(newPaths, "([^/]+)")
			continue
		}
		newPaths = append(newPaths, path)
	}

	patternText := strings.Join(newPaths, pathSep)
	patternText = fmt.Sprintf("^%s/?$", patternText)
	return regexp.MustCompile(patternText)
}

func isParam(text string) bool {
	return paramPattern.MatchString(text)
}


func NewUriMatcher(uriTemplate string) (UriMatcher, error) {
	if len(uriTemplate) < 1 {
		return UriMatcher{}, errors.New("uri is empty")
	}

	if !strings.HasPrefix(uriTemplate, pathSep) {
		return UriMatcher{}, errors.New("uri has to start with /")
	}

	paths := strings.Split(uriTemplate, pathSep)

	paramNames := extractParamName(paths)
	uriPattern := getUriPattern(paths)

	return UriMatcher{
		pattern: uriPattern,
		paramNames:  paramNames,
	}, nil
}

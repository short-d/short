package mdrouter

import (
	"net/http"
)

type HttpHandler struct {
	routes []route
}

type Handle func(w http.ResponseWriter, r *http.Request, params Params)

type route struct {
	method      string
	pathMatcher UriMatcher
	queryParams []string
	handle      Handle
}

func (r HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if route.method != req.Method {
			continue
		}

		ok, params := route.pathMatcher.IsMatch(req.URL.Path)
		if !ok {
			continue
		}
		values := req.URL.Query()
		for name, value := range values {
			if len(value) < 1 {
				continue
			}
			params[name] = value[0]
		}
		route.handle(res, req, params)
		return
	}

	http.Redirect(res, req, "/404", http.StatusNotFound)
}

func (r *HttpHandler) AddRoute(method string, isPrefix bool, path string, handle Handle) error {

	var matcher UriMatcher
	var err error

	if isPrefix {
		matcher, err = NewUriPrefixMatcher(path)
	} else {
		matcher, err = NewUriExactMatcher(path)
	}

	if err != nil {
		return err
	}

	r.routes = append(r.routes, route{
		method:      method,
		pathMatcher: matcher,
		handle:      handle,
	})
	return nil
}

func NewHttpHandler() HttpHandler {
	return HttpHandler{}
}

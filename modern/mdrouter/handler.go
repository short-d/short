package mdrouter

import (
	"net/http"
)

type HTTPHandler struct {
	routes []route
}

type Handle func(w http.ResponseWriter, r *http.Request, params Params)

type route struct {
	method      string
	pathMatcher URIMatcher
	queryParams []string
	handle      Handle
}

func (r HTTPHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
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

func (r *HTTPHandler) AddRoute(method string, isPrefix bool, path string, handle Handle) error {

	var matcher URIMatcher
	var err error

	if isPrefix {
		matcher, err = NewURIPrefixMatcher(path)
	} else {
		matcher, err = NewURIExactMatcher(path)
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

func NewHTTPHandler() HTTPHandler {
	return HTTPHandler{}
}

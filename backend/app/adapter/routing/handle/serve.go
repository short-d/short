package handle

import (
	"net/http"

	"github.com/short-d/app/fw/router"
)

// ServeDir reads files in the given directory and makes them accessible on the web.
func ServeDir(dir string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		fs := http.FileServer(http.Dir(dir))
		fs.ServeHTTP(w, r)
	}
}

// ServeFile reads a given file and makes it accessible on the web.
func ServeFile(fileName string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		http.ServeFile(w, r, fileName)
	}
}

package handle

import (
	"net/http"

	"github.com/short-d/app/fw/router"
)

// File serves files from root directory.
func File(rootDir string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		fs := http.FileServer(http.Dir(rootDir))
		fs.ServeHTTP(w, r)
	}
}

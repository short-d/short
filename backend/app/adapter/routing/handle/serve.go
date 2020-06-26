package handle

import (
	"net/http"

	"github.com/short-d/app/fw/router"
)

func ServeDir(dir string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		fs := http.FileServer(http.Dir(dir))
		fs.ServeHTTP(w, r)
	}
}

func ServeFile(fileName string) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		http.ServeFile(w, r, fileName)
	}
}

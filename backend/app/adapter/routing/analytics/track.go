package analytics

import (
	"net/http"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/short/backend/app/adapter/request"
)

// TrackHandle records event happened in the API caller.
func TrackHandle(instrumentationFactory request.InstrumentationFactory) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		i := instrumentationFactory.NewHTTP(r)

		event := params["event"]
		i.Track(event)

		w.WriteHeader(http.StatusNoContent)
	}
}

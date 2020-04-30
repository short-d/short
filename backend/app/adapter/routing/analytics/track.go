package analytics

import (
	"net/http"

	"github.com/short-d/app/fw"
	"github.com/short-d/short/app/adapter/request"
)

// TrackHandle records event happened in the API caller.
func TrackHandle(instrumentationFactory request.InstrumentationFactory) fw.Handle {
	return func(w http.ResponseWriter, r *http.Request, params fw.Params) {
		i := instrumentationFactory.NewHTTP(r)

		event := params["event"]
		i.Track(event)

		w.WriteHeader(http.StatusNoContent)
	}
}

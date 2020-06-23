package handle

import (
	"net/http"
	"net/url"

	"github.com/short-d/app/fw/router"
	"github.com/short-d/app/fw/timer"
	"github.com/short-d/short/backend/app/adapter/request"
	"github.com/short-d/short/backend/app/usecase/shortlink"
)

// LongLink translates alias to the original long link.
func LongLink(
	instrumentationFactory request.InstrumentationFactory,
	shortLinkRetriever shortlink.Retriever,
	timer timer.Timer,
	webFrontendURL url.URL,
) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, params router.Params) {
		alias := params["alias"]

		i := instrumentationFactory.NewHTTP(r)
		i.RedirectingAliasToLongLink(alias)

		now := timer.Now()
		s, err := shortLinkRetriever.GetShortLink(alias, &now)
		if err != nil {
			i.LongLinkRetrievalFailed(err)
			serve404(w, r, webFrontendURL)
			return
		}
		i.LongLinkRetrievalSucceed()

		longLink := s.LongLink
		http.Redirect(w, r, longLink, http.StatusSeeOther)
		i.RedirectedAliasToLongLink(s)
	}
}

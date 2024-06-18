package pagerouter

import (
	"log/slog"
	"net/http"

	"github.com/awryme/dnsproxy/pkg/ui"
	"github.com/awryme/slogf"
	"github.com/go-chi/chi/v5"
)

func HandlePage[T any](mux chi.Router, logf slogf.Logf, routes []Route, page Page[T], params T) {
	mux.Get(page.Addr, func(w http.ResponseWriter, r *http.Request) {
		logf("handling page", slog.String("page-addr", page.Addr))
		ui.RenderComponent(logf, w, MainPage(routes, page, params))
	})
}

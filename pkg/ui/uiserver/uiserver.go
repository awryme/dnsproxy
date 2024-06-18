package uiserver

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/pages"
	"github.com/awryme/dnsproxy/pkg/ui/staticfiles"
	"github.com/awryme/slogf"
	"github.com/go-chi/chi/v5"
	"github.com/maragudk/gomponents"
	slogchi "github.com/samber/slog-chi"
)

type Params struct {
	Addr            string
	Port            int
	RewritesStorage *rewrites.CacheStorage
	LogAccess       bool
	SettingsInfo    map[string]string
}

func Start(ctx context.Context, logHandler slog.Handler, params Params) error {
	handler := logHandler.WithAttrs([]slog.Attr{slog.String("component", "uiserver")})
	logf := slogf.New(logHandler)

	mux := chi.NewMux()
	if params.LogAccess {
		useChiLogger(mux, handler)
	}

	staticfiles.Handle(mux)

	pages.Handle(logf, mux, params.RewritesStorage, params.SettingsInfo)

	mux.Post("/add-addr", HandlerAddAddr(logf, params.RewritesStorage))
	mux.Delete("/delete-addr/{name}", HandlerDeleteAddr(logf, params.RewritesStorage))

	mux.Post("/add-rewrite", HandlerAddRewrite(logf, params.RewritesStorage))
	mux.Delete("/delete-rewrite/{domain}/{type}", HandlerDeleteRewrite(logf, params.RewritesStorage))
	mux.Post("/add-rewrite-self", HandlerAddRewriteSelf(logf, params.RewritesStorage))
	mux.Delete("/delete-rewrite-self/{domain}/{type}", HandlerDeleteRewriteSelf(logf, params.RewritesStorage))

	listenAddr := fmt.Sprintf("%s:%d", params.Addr, params.Port)
	server := http.Server{
		Addr:    listenAddr,
		Handler: mux,
		// todo: use logger
	}
	logf("starting UI server",
		slog.String("addr", params.Addr),
		slog.Int("port", params.Port),
		slog.Bool("log_access", params.LogAccess),
	)
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("run UI server failed on addr %s: %w", listenAddr, err)
	}
	return nil
}

func renderComponent(logf slogf.Logf, w http.ResponseWriter, comp gomponents.Node) {
	err := comp.Render(w)
	if err != nil {
		logf("failed to render component", slogf.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func useChiLogger(mux chi.Router, handler slog.Handler) {
	logger := slog.New(handler)

	filters := []slogchi.Filter{
		slogchi.IgnorePathPrefix("/favicon.ico"),
		slogchi.IgnorePathPrefix(staticfiles.PathStatic),
	}

	cfg := slogchi.Config{
		DefaultLevel:     slog.LevelInfo,
		ClientErrorLevel: slog.LevelWarn,
		ServerErrorLevel: slog.LevelError,

		Filters: filters,
	}

	mux.Use(slogchi.NewWithConfig(logger, cfg))
}

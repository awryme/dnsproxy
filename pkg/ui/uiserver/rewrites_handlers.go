package uiserver

import (
	"log/slog"
	"net/http"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui"
	"github.com/awryme/dnsproxy/pkg/ui/components/rewritecomponents"
	"github.com/awryme/slogf"
	"github.com/go-chi/chi/v5"
)

func HandlerAddRewrite(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logf("failed to parse form for new addr", slogf.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		domain := r.PostForm.Get("domain")
		matchType := r.PostForm.Get("type")
		addr := r.PostForm.Get("addr")

		err = rewritesStorage.AddRewrite(domain, matchType, addr)
		if err != nil {
			logf("failed to add new addr to storage", slogf.Error(err),
				slog.String("domain", domain), slog.String("type", matchType), slog.String("addr", addr))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addrs := rewritesStorage.GetAddrs()
		rewrites := rewritesStorage.GetRewrites()
		ui.RenderComponent(logf, w, rewritecomponents.PanelAll(addrs, rewrites))
	}
}

func HandlerDeleteRewrite(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")
		matchType := chi.URLParam(r, "type")
		err := rewritesStorage.DeleteRewrite(domain, matchType)
		if err != nil {
			logf("failed to delete addr from storage", slogf.Error(err), slog.String("domain", domain), slog.String("type", matchType))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addrs := rewritesStorage.GetAddrs()
		rewrites := rewritesStorage.GetRewrites()
		ui.RenderComponent(logf, w, rewritecomponents.PanelAll(addrs, rewrites))
	}
}

func HandlerAddRewriteSelf(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logf("failed to parse form for new self addr", slogf.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		domain := r.PostForm.Get("domain")
		matchType := r.PostForm.Get("type")

		err = rewritesStorage.AddRewrite(domain, matchType, rewrites.SelfAddrName)
		if err != nil {
			logf("failed to add new self addr to storage", slogf.Error(err),
				slog.String("domain", domain), slog.String("type", matchType), slog.String("addr", rewrites.SelfAddrName))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rewrites := rewritesStorage.GetRewrites()

		ui.RenderComponent(logf, w, rewritecomponents.PanelSelf(rewrites))
	}
}

func HandlerDeleteRewriteSelf(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")
		matchType := chi.URLParam(r, "type")
		err := rewritesStorage.DeleteRewrite(domain, matchType)
		if err != nil {
			logf("failed to delete addr from storage", slogf.Error(err), slog.String("domain", domain), slog.String("type", matchType))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		rewrites := rewritesStorage.GetRewrites()

		ui.RenderComponent(logf, w, rewritecomponents.PanelSelf(rewrites))
	}
}

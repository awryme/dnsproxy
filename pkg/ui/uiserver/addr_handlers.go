package uiserver

import (
	"log/slog"
	"net/http"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/addrcomponents"
	"github.com/awryme/slogf"
	"github.com/go-chi/chi/v5"
)

func HandlerAddAddr(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			logf("failed to parse form for new addr", slogf.Error(err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		name := r.PostForm.Get("name")
		ip := r.PostForm.Get("ip")
		err = rewritesStorage.AddAddr(name, ip)
		if err != nil {
			logf("failed to add new addr to storage", slogf.Error(err), slog.String("name", name), slog.String("ip", ip))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addrs := rewritesStorage.GetAddrs()
		renderComponent(logf, w, addrcomponents.Panel(addrs))
		renderComponent(logf, w, addrcomponents.Selector(addrs))
	}
}

func HandlerDeleteAddr(logf slogf.Logf, rewritesStorage *rewrites.CacheStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		err := rewritesStorage.DeleteAddr(name)
		if err != nil {
			logf("failed to delete addr from storage", slogf.Error(err), slog.String("name", name))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		addrs := rewritesStorage.GetAddrs()
		renderComponent(logf, w, addrcomponents.Panel(addrs))
		renderComponent(logf, w, addrcomponents.Selector(addrs))
	}
}

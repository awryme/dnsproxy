package pages

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/pagerouter"
	"github.com/awryme/dnsproxy/pkg/ui/components/rewritecomponents"
	"github.com/awryme/slogf"
	"github.com/go-chi/chi/v5"
)

var PageRewritesSelf = pagerouter.MakePage("/", "Self Rewrites", rewritecomponents.PageSelf)
var PageRewritesAll = pagerouter.MakePage("/rewrites", "All Rewrites", rewritecomponents.PageAll)
var PageSettings = pagerouter.MakePage("/settings", "Settings", pagerouter.SettingsPage)

var pageTabs = []pagerouter.Route{
	pagerouter.MakeRoute(PageRewritesSelf),
	pagerouter.MakeRoute(PageRewritesAll),
	pagerouter.MakeRoute(PageSettings),
}

func Handle(logf slogf.Logf, mux chi.Router, rewritesStorage *rewrites.CacheStorage, settingsInfo map[string]string) {
	pagerouter.HandlePage(logf, mux, pageTabs, PageRewritesSelf, rewritesStorage)
	pagerouter.HandlePage(logf, mux, pageTabs, PageRewritesAll, rewritesStorage)
	pagerouter.HandlePage(logf, mux, pageTabs, PageSettings, settingsInfo)
}

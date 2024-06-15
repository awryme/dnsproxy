package components

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/rewritecomponents"
	"github.com/maragudk/gomponents"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

type Page[T any] struct {
	Component func(params T) gomponents.Node
	TabIdx    int
	Tab       func() gomplements.Element
}

var PageRewritesSelf = Page[*rewrites.CacheStorage]{
	Component: func(storage *rewrites.CacheStorage) gomponents.Node {
		return rewritecomponents.PageSelf(storage)
	},
	TabIdx: 0,
	Tab: func() gomplements.Element {
		return bulma.TabAHref("/", "Self Rewrites")
	},
}

var PageRewritesAll = Page[*rewrites.CacheStorage]{
	Component: func(storage *rewrites.CacheStorage) gomponents.Node {
		return rewritecomponents.PageAll(storage)
	},
	TabIdx: 1,
	Tab: func() gomplements.Element {
		return bulma.TabAHref("/rewrites", "All Rewrites")
	},
}

var PageSettings = Page[map[string]string]{
	Component: func(info map[string]string) gomponents.Node {
		return SettingsPage(info)
	},
	TabIdx: 2,
	Tab: func() gomplements.Element {
		return bulma.TabAHref("/settings", "Settings")
	},
}

func Tabbar(tabidx int) gomplements.Element {
	tabElemtents := []gomplements.Element{
		PageRewritesSelf.Tab(),
		PageRewritesAll.Tab(),
		PageSettings.Tab(),
	}
	tabElemtents[tabidx].With(bulma.Active)

	tabs := bulma.Tabs(
		bulma.Centered,
		bulma.Medium,
	)

	for _, elem := range tabElemtents {
		tabs.With(elem)
	}
	return tabs
}

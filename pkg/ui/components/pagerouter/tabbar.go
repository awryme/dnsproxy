package pagerouter

import (
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Tabbar[T any](routes []Route, page Page[T]) gomplements.Element {
	tabElements, idxMapping := makeTabs(routes)
	idx := idxMapping[page.Addr]
	tabElements[idx].With(bulma.Active)

	tabs := bulma.Tabs(
		bulma.Centered,
		bulma.Medium,
	)

	for _, elem := range tabElements {
		tabs.With(elem)
	}
	return tabs
}

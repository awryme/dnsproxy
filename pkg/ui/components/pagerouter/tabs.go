package pagerouter

import "github.com/willoma/gomplements"

type Route struct {
	addr string
	tab  func() gomplements.Element
}

func MakeRoute[T any](p Page[T]) Route {
	return Route{
		p.Addr,
		p.Tab,
	}
}

func makeTabs(routes []Route) ([]gomplements.Element, map[string]int) {
	tabs := make([]gomplements.Element, 0, len(routes))
	idxMapping := make(map[string]int, len(routes))

	for i, tabInfo := range routes {
		tabs = append(tabs, tabInfo.tab())
		idxMapping[tabInfo.addr] = i
	}
	return tabs, idxMapping
}

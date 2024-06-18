package pagerouter

import (
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

type Page[T any] struct {
	Component func(params T) gomplements.Element
	Addr      string
	Info      string
}

func (p Page[T]) Tab() gomplements.Element {
	return bulma.TabAHref(p.Addr, p.Info)
}

func MakePage[T any](addr string, info string, comp func(params T) gomplements.Element) Page[T] {
	return Page[T]{
		Addr:      addr,
		Info:      info,
		Component: comp,
	}
}

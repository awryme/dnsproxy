package hlist

import (
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

type flag int32

const (
	Dynamic flag = 1 << iota
	Static

	PinLast
	PinFirst
)

func List(children ...any) gomplements.Element {
	columns := bulma.Columns()

	isFirstCol := true
	setFlags := map[flag]bool{}

	for childIdx, c := range children {
		switch f := c.(type) {
		case flag:
			setFlags[f] = true
			continue
		}
		isLastCol := childIdx == len(children)-1
		col := bulma.Column()

		if isFirstCol && setFlags[PinFirst] {
			col.With(bulma.Narrow)
			isFirstCol = false
		}
		if isLastCol && setFlags[PinLast] {
			col.With(bulma.Narrow)
		}
		if !setFlags[PinFirst] && !setFlags[PinLast] {
			col.With(bulma.Narrow)
		}
		col.With(c)
		columns.With(col)
	}

	if !setFlags[Dynamic] {
		columns.With(bulma.Mobile)
	}
	return columns
}

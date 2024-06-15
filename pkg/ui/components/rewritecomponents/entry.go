package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/hlist"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Entry(entry rewrites.Entry, isSelf bool) gomplements.Element {
	listType := hlist.Dynamic
	if isSelf {
		listType = hlist.Static
	}
	return bulma.PanelBlock(
		bulma.Container(
			hlist.List(
				hlist.PinLast,
				hlist.List(
					hlist.PinLast, listType,
					RewriteInfo(entry, isSelf),
					RewriteTags(entry, isSelf),
				),
				ButtonDelete(entry, isSelf),
			),
		),
	)
}

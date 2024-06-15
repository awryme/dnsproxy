package addrcomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Entry(addr rewrites.Addr) gomplements.Element {
	var deleteButton gomplements.Element
	if addr.Name != rewrites.SelfAddrName {
		deleteButton = gomplements.Div(
			bulma.PulledRight,
			ButtonDelete(addr),
		)
	}
	return bulma.PanelBlock(
		bulma.Container(
			gomplements.Div(
				bulma.PulledLeft,
				bulma.MarginLeft(2),
				addr.Name,
			),
			gomplements.Div(
				bulma.PulledLeft,
				bulma.MarginLeft(2),
				"->",
			),
			gomplements.Div(
				bulma.PulledLeft,
				bulma.MarginLeft(2),
				addr.IP.String(),
			),
			deleteButton,
		),
	)
}

package addrcomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/ids"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Panel(addrs []rewrites.Addr) gomplements.Element {
	panel := bulma.Panel(
		ids.AddrPanel.ID(),
		bulma.PanelHeading("Addrs"),
	)

	for _, addr := range addrs {
		panel.With(Entry(addr))
	}
	panel.With(bulma.PanelBlock(
		FormAdd(),
	))

	return panel
}

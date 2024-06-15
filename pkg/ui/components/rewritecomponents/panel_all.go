package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/ids"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PanelAll(addrs []rewrites.Addr, rewrites rewrites.Rewrites) gomplements.Element {
	allRewrites := rewrites.SortedByDomain()

	panel := bulma.Panel(
		ids.RewritesPanel.ID(),
		bulma.PanelHeading("Rewrites"),
	)

	for _, entry := range allRewrites {
		// panel.With(RewritesEntry(entry, false))
		panel.With(Entry(entry, false))
	}
	panel.With(FormAdd(addrs, false))

	return panel
}

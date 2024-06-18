package rewritecomponents

import (
	"slices"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PanelSelf(entrySet rewrites.EntrySet) gomplements.Element {
	allRewrites := entrySet.SortedByDomain()
	selfRewrites := slices.Clone(allRewrites)
	selfRewrites = slices.DeleteFunc(selfRewrites, func(entry rewrites.Entry) bool {
		return entry.AddrName != rewrites.SelfAddrName
	})

	panel := bulma.Panel(
		IDRewritesPanel,
		bulma.PanelHeading("Self Rewrites"),
	)

	for _, entry := range selfRewrites {
		panel.With(Entry(entry, true))
	}
	panel.With(FormAdd(nil, true))

	return panel
}

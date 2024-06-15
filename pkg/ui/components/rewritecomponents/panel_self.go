package rewritecomponents

import (
	"slices"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/ids"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PanelSelf(rewritesStorage *rewrites.CacheStorage) gomplements.Element {
	allRewrites := rewritesStorage.GetRewrites().SortedByDomain()
	selfRewrites := slices.Clone(allRewrites)
	selfRewrites = slices.DeleteFunc(selfRewrites, func(entry rewrites.Entry) bool {
		return entry.AddrName != rewrites.SelfAddrName
	})

	panel := bulma.Panel(
		ids.RewritesPanel.ID(),
		bulma.PanelHeading("Self Rewrites"),
	)

	for _, entry := range selfRewrites {
		panel.With(Entry(entry, true))
	}
	panel.With(FormAdd(nil, true))

	return panel
}

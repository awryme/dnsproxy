package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/addrcomponents"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PageAll(rewritesStorage *rewrites.CacheStorage) gomplements.Element {
	addrs := rewritesStorage.GetAddrs()
	rewrites := rewritesStorage.GetRewrites()

	addrsPanel := addrcomponents.Panel(addrs)
	rewritesPanel := PanelAll(addrs, rewrites)
	return bulma.Container(
		addrsPanel,
		rewritesPanel,
	)
}

package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PageSelf(rewritesStorage *rewrites.CacheStorage) gomplements.Element {
	rewrites := rewritesStorage.GetRewrites()
	rewritesPanel := PanelSelf(rewrites)
	return bulma.Container(
		rewritesPanel,
	)
}

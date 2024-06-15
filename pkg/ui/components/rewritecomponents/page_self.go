package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func PageSelf(rewritesStorage *rewrites.CacheStorage) gomplements.Element {
	rewritesPanel := PanelSelf(rewritesStorage)
	return bulma.Container(
		rewritesPanel,
	)
}

package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/hlist"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func RewriteTags(entry rewrites.Entry, isSelf bool) gomplements.Element {
	if isSelf {
		return hlist.List(
			bulma.Tag(
				bulma.Info,
				string(entry.Type),
			),
		)
	}
	return hlist.List(
		bulma.Tag(
			bulma.Primary,
			entry.AddrName,
		),
		bulma.Tag(
			bulma.Info,
			string(entry.Type),
		),
	)
}

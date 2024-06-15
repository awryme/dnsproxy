package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/hlist"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func RewriteInfo(entry rewrites.Entry, isSelf bool) gomplements.Element {
	domainInfo := gomplements.Span()
	if entry.Type == rewrites.MatchTypeSuffix {
		domainInfo.With(
			gomplements.Span(
				bulma.TextGrey,
				"*.",
			),
		)
	}
	domainInfo.With(entry.Domain)
	if isSelf {
		return hlist.List(domainInfo)
	}

	return hlist.List(
		domainInfo,
		"->",
		entry.IP,
	)
}

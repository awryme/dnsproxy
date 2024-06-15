package rewritecomponents

import (
	"fmt"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func TypeSelector() gomplements.Element {
	return bulma.Select(
		bulma.OnSelect(gomplements.Name("type")),
		bulma.Option(rewrites.MatchTypeSuffix.String(), fmt.Sprintf("match: %s", rewrites.MatchTypeSuffix.String())),
		bulma.Option(rewrites.MatchTypeStrict.String(), fmt.Sprintf("match: %s", rewrites.MatchTypeStrict.String())),
	)
}

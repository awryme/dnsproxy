package rewritecomponents

import (
	"fmt"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/ids"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func ButtonDelete(entry rewrites.Entry, isSelf bool) gomplements.Element {
	deleteUrl := fmt.Sprintf("/delete-rewrite-self/%s/%s", entry.Domain, entry.Type)
	if !isSelf {
		deleteUrl = fmt.Sprintf("/delete-rewrite/%s/%s", entry.Domain, entry.Type)
	}

	return bulma.Delete(
		htmx.Delete(deleteUrl),
		htmx.Target(ids.RewritesPanel.Query()),
	)
}

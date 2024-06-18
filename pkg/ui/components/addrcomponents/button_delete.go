package addrcomponents

import (
	"fmt"

	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/queryselector"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func ButtonDelete(addr rewrites.Addr) gomplements.Element {
	return bulma.Delete(
		htmx.Delete(fmt.Sprintf("/delete-addr/%s", addr.Name)),
		htmx.Target(queryselector.ID(IDAddrPanel)),
		htmx.SelectOOB(queryselector.ID(IDAddrSelector)),
	)
}

package addrcomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Selector(addrs []rewrites.Addr) gomplements.Element {
	addrSelector := bulma.Select(
		IDAddrSelector,
		bulma.OnSelect(gomplements.Name("addr")),
	)
	for _, addr := range addrs {
		addrSelector.With(bulma.Option(addr.Name, addr.Name))
	}
	return addrSelector
}

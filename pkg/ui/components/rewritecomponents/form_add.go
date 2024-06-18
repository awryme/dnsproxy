package rewritecomponents

import (
	"github.com/awryme/dnsproxy/pkg/rewrites"
	"github.com/awryme/dnsproxy/pkg/ui/components/addrcomponents"
	"github.com/awryme/dnsproxy/pkg/ui/queryselector"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func FormAdd(addrs []rewrites.Addr, isSelfOnly bool) gomplements.Element {
	var addrSelector gomplements.Element
	if !isSelfOnly {
		addrSelector = addrcomponents.Selector(addrs)
	}

	var postRewriteUrl = "/add-rewrite-self"
	if !isSelfOnly {
		postRewriteUrl = "/add-rewrite"
	}

	return bulma.PanelBlock(
		gomplements.Form(
			htmx.Post(postRewriteUrl),
			htmx.Target(queryselector.ID(IDRewritesPanel)),
			bulma.Field(
				bulma.GroupedMultiline,
				bulma.Control(
					bulma.InputText(
						gomplements.Name("domain"),
						gomplements.Placeholder("domain"),
					),
				),
				bulma.Control(
					TypeSelector(),
				),
				bulma.Control(addrSelector),
				bulma.Control(
					bulma.ButtonSubmit(
						bulma.LinkBold,
						"Add",
					),
				),
			),
		),
	)
}

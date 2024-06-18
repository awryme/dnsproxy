package addrcomponents

import (
	"github.com/awryme/dnsproxy/pkg/ui/queryselector"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func FormAdd() gomplements.Element {
	return gomplements.Form(
		htmx.Post("/add-addr"),
		htmx.Target(queryselector.ID(IDAddrPanel)),
		htmx.SelectOOB(queryselector.ID(IDAddrSelector)),
		bulma.Field(
			bulma.GroupedMultiline,
			bulma.Control(
				bulma.InputText(
					gomplements.Name("name"),
					gomplements.Placeholder("name"),
				),
			),
			bulma.Control(
				bulma.InputText(
					gomplements.Name("ip"),
					gomplements.Placeholder("ip"),
				),
			),
			bulma.Control(
				bulma.ButtonSubmit(
					bulma.LinkBold,
					"Add",
				),
			),
		),
	)
}

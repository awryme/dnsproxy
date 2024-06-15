package addrcomponents

import (
	"github.com/awryme/dnsproxy/pkg/ui/components/ids"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func FormAdd() gomplements.Element {
	return gomplements.Form(
		htmx.Post("/add-addr"),
		htmx.Target(ids.AddrPanel.Query()),
		htmx.SelectOOB(ids.AddrSelector.Query()),
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

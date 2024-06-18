package pagerouter

import (
	"github.com/maragudk/gomponents"
	htmx "github.com/maragudk/gomponents-htmx"
	"github.com/maragudk/gomponents/html"
	"github.com/willoma/bulma-gomponents"
)

func MainPage[T any](routes []Route, page Page[T], params T) gomponents.Node {
	return html.Doctype(
		html.HTML(
			html.Class("theme-dark"),
			html.Lang("en"),
			html.Title("dnsproxy"),
			html.Head(
				html.Meta(html.Charset("utf-8")),
				html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1")),
				html.Link(html.Rel("stylesheet"), html.Href("/static/bulma.min.css")),
				html.Link(html.Rel("stylesheet"), html.Href("/static/styles.css")),
				html.Script(html.Src("/static/htmx.min.js")),
			),
			html.Body(
				htmx.Boost("true"),
				bulma.Container(
					bulma.MaxDesktop,
					Header(),
					Tabbar(routes, page),
					page.Component(params),
				),
			),
		),
	)
}

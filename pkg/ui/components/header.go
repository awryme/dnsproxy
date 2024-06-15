package components

import (
	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func Header() gomplements.Element {
	return bulma.Notification(
		bulma.Primary,
		bulma.Title("DnsProxy"),
	)
}

package components

import (
	"slices"

	"github.com/willoma/bulma-gomponents"
	"github.com/willoma/gomplements"
)

func SettingsPage(settings map[string]string) gomplements.Element {
	grid := bulma.FixedGrid(
		bulma.Cell(
			bulma.TextInfo,
			bulma.PulledRight,
			"Name",
		),
		bulma.Cell(
			bulma.TextInfo,
			bulma.PulledLeft,
			"Value",
		),
	)
	walkMapSorted(settings, func(name, value string) {
		grid.With(
			bulma.Cell(
				bulma.PulledRight,
				name,
			),
			bulma.Cell(
				bulma.PulledLeft,
				value,
			),
		)
	})
	return grid
}

func walkMapSorted(settings map[string]string, callback func(name string, value string)) {
	names := make([]string, 0, len(settings))
	for name := range settings {
		names = append(names, name)
	}
	slices.Sort(names)
	for _, name := range names {
		callback(name, settings[name])
	}
}

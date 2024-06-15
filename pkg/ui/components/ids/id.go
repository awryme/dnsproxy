package ids

import "github.com/willoma/gomplements"

type ID string

func (id ID) ID() gomplements.ID {
	return gomplements.ID(id)
}

func (id ID) Query() string {
	return "#" + string(id)
}

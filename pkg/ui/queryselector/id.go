package queryselector

import "github.com/willoma/gomplements"

func ID(id gomplements.ID) string {
	return "#" + string(id)
}

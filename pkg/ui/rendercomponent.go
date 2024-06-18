package ui

import (
	"net/http"

	"github.com/awryme/slogf"
	"github.com/maragudk/gomponents"
)

func RenderComponent(logf slogf.Logf, w http.ResponseWriter, comp gomponents.Node) {
	err := comp.Render(w)
	if err != nil {
		logf("failed to render component", slogf.Error(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

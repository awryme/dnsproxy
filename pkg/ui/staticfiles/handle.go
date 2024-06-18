package staticfiles

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const PathStatic = "/static/"

func Handle(mux chi.Router) {
	mux.Get(PathStatic+"*", func(w http.ResponseWriter, r *http.Request) {
		fileName := chi.URLParam(r, "*")
		filePath := fmt.Sprintf("files/%s", fileName)
		http.ServeFileFS(w, r, embeddedFiles, filePath)
	})
}

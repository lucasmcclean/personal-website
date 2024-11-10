package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/lucasmcclean/personal-website/pages"
)

func addRoutes(mux *http.ServeMux) {
	staticFS := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFS))

	mux.Handle("/{$}", templ.Handler(pages.Index()))
}

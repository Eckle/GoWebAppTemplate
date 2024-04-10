package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pbentes/80_20/src/templates/components"
	"github.com/pbentes/80_20/src/templates/pages"
)

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	fs := http.FileServer(http.Dir("assets"))
	router.Handle("GET /assets/", http.StripPrefix("/assets/", fs))

	router.HandleFunc("GET /", templ.Handler(components.Layout(pages.Index)).ServeHTTP)
	router.HandleFunc("GET /index", templ.Handler(pages.Index()).ServeHTTP)
	router.HandleFunc("GET /content", templ.Handler(pages.Content()).ServeHTTP)
	return router
}

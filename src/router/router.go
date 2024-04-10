package router

import (
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/pbentes/80_20/src/templates/components"
	"github.com/pbentes/80_20/src/templates/pages"
)

var dir string = "assets"
var fs http.Handler = http.FileServer(http.Dir(dir))

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /", ServeFiles(templ.Handler(components.Layout(pages.Index))))

	router.HandleFunc("GET /index", templ.Handler(pages.Index()).ServeHTTP)
	router.HandleFunc("GET /content", templ.Handler(pages.Content()).ServeHTTP)
	return router
}

func ServeFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := dir + r.URL.Path
		_, err := os.Stat(path)
		if os.IsNotExist(err) || r.URL.Path == "/" {
			next.ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

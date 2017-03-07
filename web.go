package main

import (
	"crypto/subtle"
	"net/http"
	"os"

	"github.com/flosch/pongo2"
	"github.com/gorilla/mux"
)

func setup_routes() {
	username := os.Getenv("HTTP_USERNAME")
	password := os.Getenv("HTTP_PASSWORD")

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", BasicAuth(homepage, username, password))
	r.HandleFunc("/preview/{name}", BasicAuth(preview, username, password))
	r.HandleFunc("/status", BasicAuth(status, username, password))

	fs := http.FileServer(http.Dir("static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)

}

func BasicAuth(handler http.HandlerFunc, username string, password string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(username)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(password)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter auth details"`)
			w.WriteHeader(401)
			w.Write([]byte("Unauthorised.\n"))
			return
		}
		handler(w, r)
	}
}

var tplIndex = pongo2.Must(pongo2.FromFile("templates/index.html"))
var tplPreview = pongo2.Must(pongo2.FromFile("templates/preview.html"))
var tplStatus = pongo2.Must(pongo2.FromFile("templates/status.html"))
var tpl404 = pongo2.Must(pongo2.FromFile("templates/404.html"))

func notFound(w http.ResponseWriter, r *http.Request) {
	tpl404.ExecuteWriter(pongo2.Context{}, w)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	tplIndex.ExecuteWriter(pongo2.Context{"adapters": config.Adapters}, w)
}

func status(w http.ResponseWriter, r *http.Request) {
	tplStatus.ExecuteWriter(pongo2.Context{}, w)
}

func preview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adapter := findAdapter(vars["name"])

	// Get most recent file in adapter.Arguments["folder"]
	// Provide to template for display

	headers := []string{"Header1", "Header2", "Header3"}
	rows := [][]string{}
	rows = append(rows, []string{"1", "2", "3"})
	rows = append(rows, []string{"4", "5", "6"})
	rows = append(rows, []string{"7", "8", "9"})

	tplPreview.ExecuteWriter(pongo2.Context{
		"adapter": adapter, "headers": headers, "rows": rows}, w)
}

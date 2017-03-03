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
	r.HandleFunc("/preview", BasicAuth(preview, username, password))
	r.HandleFunc("/status", BasicAuth(status, username, password))
	r.HandleFunc("/add", BasicAuth(add_adapter, username, password))
	r.HandleFunc("/add/directory", BasicAuth(add_directory_adapter, username, password))

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
var tplAddAdapter = pongo2.Must(pongo2.FromFile("templates/adapter_add.html"))
var tplAddAdapterDirectory = pongo2.Must(pongo2.FromFile("templates/adapter_add_directory.html"))
var tpl404 = pongo2.Must(pongo2.FromFile("templates/404.html"))

func notFound(w http.ResponseWriter, r *http.Request) {
	tpl404.ExecuteWriter(pongo2.Context{}, w)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	tplIndex.ExecuteWriter(pongo2.Context{}, w)
}

func add_adapter(w http.ResponseWriter, r *http.Request) {
	tplAddAdapter.ExecuteWriter(pongo2.Context{}, w)
}

func add_directory_adapter(w http.ResponseWriter, r *http.Request) {
	errors := make(map[string]string)
	data := make(map[string]string)

	if r.Method == "POST" {
		r.ParseForm()
		data["title"] = r.Form.Get("title")
		data["dataset"] = r.Form.Get("dataset")
		data["folder"] = r.Form.Get("folder")

		if data["title"] == "" {
			errors["title"] = "Title is required"
		}
		if data["dataset"] == "" {
			errors["dataset"] = "Dataset name is required"
		}
		if data["folder"] == "" {
			errors["folder"] = "File folder is required"
		}
	}

	tplAddAdapterDirectory.ExecuteWriter(pongo2.Context{"errors": errors, "data": data}, w)
}

func status(w http.ResponseWriter, r *http.Request) {
	tplStatus.ExecuteWriter(pongo2.Context{}, w)
}

func preview(w http.ResponseWriter, r *http.Request) {
	tplPreview.ExecuteWriter(pongo2.Context{}, w)
}

package main

import (
	"crypto/subtle"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	username := os.Getenv("HTTP_USERNAME")
	password := os.Getenv("HTTP_PASSWORD")

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/status", BasicAuth(status, username, password))
	http.HandleFunc("/preview", BasicAuth(preview, username, password))
	http.HandleFunc("/", BasicAuth(homepage, username, password))

	log.Println("Listening...")
	http.ListenAndServe(":"+port, nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	tmpl := get_template("index.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func status(w http.ResponseWriter, r *http.Request) {
	tmpl := get_template("status.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func preview(w http.ResponseWriter, r *http.Request) {
	tmpl := get_template("preview.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func get_template(name string) *template.Template {
	lp := filepath.Join("templates", "layout.html")
	fp := filepath.Join("templates", name)

	tmpl, _ := template.ParseFiles(lp, fp)
	return tmpl
}

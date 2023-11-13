package main

import (
	"html/template"
	"net/http"
)

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./dist"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template.Must(template.ParseFiles("./index.html")).Execute(w, nil)
	})

	http.ListenAndServe(":8000", nil)
}

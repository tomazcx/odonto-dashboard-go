package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

func main() {

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.Handle("/login", templ.Handler(views.Login()))

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", nil)
}

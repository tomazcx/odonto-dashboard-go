package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tomazcx/odonto-dashboard-go/views"
	"github.com/tomazcx/odonto-dashboard-go/views/layouts"
)

func main() {

	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	http.Handle("/login", templ.Handler(views.Login()))
	http.Handle("/", templ.Handler(layouts.DashboardLayout()))

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", nil)
}

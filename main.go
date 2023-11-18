package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

func main() {

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	http.Handle("/login", templ.Handler(views.Login()))
	http.Handle("/", templ.Handler(views.Home()))
	http.Handle("/patients", templ.Handler(views.Clients()))
	http.Handle("/newPatient", templ.Handler(views.NewClient()))

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", nil)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/routes"
)

func main() {

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	routes.AuthRoutes()
	routes.DashboardRoutes()

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", nil)
}

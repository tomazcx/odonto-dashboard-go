package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/tomazcx/odonto-dashboard-go/application/routes"
	"github.com/tomazcx/odonto-dashboard-go/application/utils/authutils"
)

func main() {

	//Load the env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	authutils.Initialize()

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	routes.AuthRoutes()
	routes.DashboardRoutes()

	fmt.Println("Server running at port 8000!")
	http.ListenAndServe(":8000", nil)
}

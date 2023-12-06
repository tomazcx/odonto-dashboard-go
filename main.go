package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomazcx/odonto-dashboard-go/application/routes"
	"github.com/tomazcx/odonto-dashboard-go/application/utils/authutils"
	"github.com/tomazcx/odonto-dashboard-go/infra/db"
)

func main() {

	//Load the env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db.ConnectToDatabase()
	db := db.GetConn()
	defer db.Close()

	authutils.Initialize()

	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	routes.DefineRoutes()

	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "8000"
	}

	fmt.Println("Server running at port " + port + "!")
	http.ListenAndServe(":"+port, nil)
}

package routes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
)

func AuthRoutes() {
	http.HandleFunc("/login", controllers.AuthController{}.Index)

	http.HandleFunc("/logout", controllers.AuthController{}.Logout)
	http.HandleFunc("/login/submit", controllers.AuthController{}.Login)
}

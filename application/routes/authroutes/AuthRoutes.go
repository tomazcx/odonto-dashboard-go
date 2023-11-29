package authroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
)

func AuthRoutes() {

	authController := controllers.AuthController{}

	http.HandleFunc("/login", authController.Index)

	http.HandleFunc("/logout", authController.Logout)
	http.HandleFunc("/login/submit", authController.Login)
}

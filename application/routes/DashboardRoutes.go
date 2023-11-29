package routes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func DashboardRoutes() {
	http.HandleFunc("/", middlewares.UseAuth(controllers.HomeController{}.Index))
}

package routes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
)

func DashboardRoutes() {
	http.HandleFunc("/", controllers.HomeController{}.Index)
	http.HandleFunc("/patients", controllers.ClientsController{}.Index)
	http.HandleFunc("/newPatient", controllers.NewClientController{}.Index)
}

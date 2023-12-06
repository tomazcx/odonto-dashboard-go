package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func AppointmentsRoutes() {
	appointmentsController := controllers.AppointmentsController{}

	http.HandleFunc("/appointments", middlewares.UseAuth(appointmentsController.Index))
	http.HandleFunc("/appointment/create", middlewares.UseAuth(appointmentsController.Create))
	http.HandleFunc("/appointment/delete", middlewares.UseAuth(appointmentsController.Delete))
	http.HandleFunc("/appointment/getPage", middlewares.UseAuth(appointmentsController.ChangePage))
}

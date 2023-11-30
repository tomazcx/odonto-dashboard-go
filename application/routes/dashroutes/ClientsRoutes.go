package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func ClientsRoutes() {
	clientsController := controllers.ClientsController{}

	http.HandleFunc("/patients", middlewares.UseAuth(clientsController.Index))
	http.HandleFunc("/patients/getPatients", middlewares.UseAuth(clientsController.GetClients))
	http.HandleFunc("/patients/getPage", middlewares.UseAuth(clientsController.ChangePage))
}

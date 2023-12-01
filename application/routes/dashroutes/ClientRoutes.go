package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func ClientRoutes() {
	clientsController := controllers.ClientController{}

	http.HandleFunc("/patient", middlewares.UseAuth(clientsController.Index))
}

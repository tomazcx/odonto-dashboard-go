package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func ClientRoutes() {
	clientController := controllers.ClientController{}

	http.HandleFunc("/patient", middlewares.UseAuth(clientController.Index))
	http.HandleFunc("/patient/delete", middlewares.UseAuth(clientController.DeleteClient))
}

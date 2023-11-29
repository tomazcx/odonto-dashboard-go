package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func NewClientRoutes() {
	newClientController := controllers.NewClientController{}

	http.HandleFunc("/newPatient", middlewares.UseAuth(newClientController.Index))
	http.HandleFunc("/newPatient/handleCreatePatient", middlewares.UseAuth(newClientController.CreateClient))

}

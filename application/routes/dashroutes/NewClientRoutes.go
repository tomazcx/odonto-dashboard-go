package dashroutes

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/controllers"
	"github.com/tomazcx/odonto-dashboard-go/application/middlewares"
)

func NewClientRoutes() {
	http.HandleFunc("/newPatient", middlewares.UseAuth(controllers.NewClientController{}.Index))
	http.HandleFunc("/newPatient/handleCreatePatient", middlewares.UseAuth(controllers.NewClientController{}.CreateClient))

}

package routes

import (
	"github.com/tomazcx/odonto-dashboard-go/application/routes/authroutes"
	"github.com/tomazcx/odonto-dashboard-go/application/routes/dashroutes"
)

func DefineRoutes() {
	authroutes.AuthRoutes()
	dashroutes.ClientsRoutes()
	dashroutes.ClientRoutes()
	dashroutes.NewClientRoutes()
}

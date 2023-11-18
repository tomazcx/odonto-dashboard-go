package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/views"
)

type ClientsController struct{}

func (a ClientsController) Index(w http.ResponseWriter, r *http.Request) {
	views.Clients().Render(r.Context(), w)
}

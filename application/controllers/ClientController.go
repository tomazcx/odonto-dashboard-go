package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/views"
)

type ClientController struct {
}

func (c ClientController) Index(w http.ResponseWriter, r *http.Request) {
	views.Client().Render(r.Context(), w)
}

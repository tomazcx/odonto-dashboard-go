package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/views"
)

type NewClientController struct{}

func (a NewClientController) Index(w http.ResponseWriter, r *http.Request) {
	views.NewClient().Render(r.Context(), w)
}

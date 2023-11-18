package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/views"
)

type HomeController struct{}

func (a HomeController) Index(w http.ResponseWriter, r *http.Request) {
	views.Home().Render(r.Context(), w)
}

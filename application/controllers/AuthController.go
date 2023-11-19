package controllers

import (
	"log"
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/application/utils/authutils"
	"github.com/tomazcx/odonto-dashboard-go/components"
	"github.com/tomazcx/odonto-dashboard-go/data/services/auth"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

type AuthController struct{}

func (a AuthController) Index(w http.ResponseWriter, r *http.Request) {
	views.Login().Render(r.Context(), w)
}

func (a AuthController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		w.Header().Set("HX-Retarget", "#formError")
		components.ErrorMessage("Erro ao processar o formulário. Contate o administrador.").Render(r.Context(), w)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	err := auth.LoginService{}.Execute(login, password)

	if err != nil {
		w.Header().Set("HX-Retarget", "#formError")
		components.ErrorMessage(err.Error()).Render(r.Context(), w)
		return
	}

	err = authutils.SetUserSession(w, r)

	if err != nil {
		log.Printf("Error logging the user: %v", err)
		w.Header().Set("HX-Retarget", "#formError")
		components.ErrorMessage("Error logging the user").Render(r.Context(), w)
		return
	}
	w.Header().Set("HX-Redirect", "/")
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	err := authutils.ExpireSession(w, r)

	if err != nil {
		log.Fatalf("Error logging the use out: %v", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

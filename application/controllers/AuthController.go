package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/components"
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

	err := r.ParseForm()

	if err != nil {
		w.Header().Set("HX-Retarget", "#formError")
		components.ErrorMessage("Erro ao processar o formulário. Contate o administrador.").Render(r.Context(), w)
		return
	}

	//	login := r.FormValue("login")
	//	password := r.FormValue("password")

	w.Header().Set("HX-Retarget", "#formError")
	components.ErrorMessage("Credenciais inválidas.").Render(r.Context(), w)
	return
}

func (a AuthController) Logout(w http.ResponseWriter, r *http.Request) {

}

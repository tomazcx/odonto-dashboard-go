package controllers

import (
	"net/http"

	"github.com/tomazcx/odonto-dashboard-go/data/services/client"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

type ClientController struct {
}

func (c ClientController) Index(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	GetClientService := client.GetClientService{}
	client, err := GetClientService.Execute(id)

	if err != nil {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	views.Client(client).Render(r.Context(), w)
}

func (c ClientController) DeleteClient(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	DeleteClientService := client.DeleteClientService{}
	err := DeleteClientService.Execute(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("HX-Redirect", "/patients")
}

package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomazcx/odonto-dashboard-go/components"
	"github.com/tomazcx/odonto-dashboard-go/data/dto"
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

func (c ClientController) UpdateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		components.ErrorMessage("Erro ao processar o formulário. Contate o administrador.").Render(r.Context(), w)
		return
	}

	name := r.FormValue("name")
	age := r.FormValue("age")
	telephone := r.FormValue("telephone")
	city := r.FormValue("city")
	address := r.FormValue("address")
	district := r.FormValue("district")
	budget := r.FormValue("budget")
	budgetDescription := r.FormValue("budgetDescription")
	anamnese := r.FormValue("anamnese")

	ageInt, _ := strconv.Atoi(age)

	id := r.URL.Query().Get("id")

	if len(id) == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	updateClientDto := dto.UpdateClientDTO{Name: name, Age: ageInt, Telephone: telephone, City: city, Address: address, District: district, Budget: budget, BudgetDescription: budgetDescription, Anamnese: anamnese}
	UpdateClientService := client.UpdateClientService{}

	err := UpdateClientService.Execute(r.Context(), id, updateClientDto)

	if err != nil {
		components.ErrorMessage(fmt.Sprintf("Erro ao salvar os dados do paciente: %v", err)).Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Trigger", "patientUpdated")
	components.ErrorMessage("").Render(r.Context(), w)

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

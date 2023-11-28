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

type NewClientController struct{}

func (a NewClientController) Index(w http.ResponseWriter, r *http.Request) {
	views.NewClient().Render(r.Context(), w)
}

func (a NewClientController) CreateClient(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	createClientDto := dto.CreateClientDTO{Name: name, Age: ageInt, Telephone: telephone, City: city, Address: address, District: district, Budget: budget, BudgetDescription: budgetDescription, Anamnese: anamnese}
	err := client.CreateClientService{}.Execute(createClientDto)

	if err != nil {
		components.ErrorMessage(fmt.Sprintf("Erro ao cadastrar o paciente: %v", err)).Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Trigger", "patientCreated")
	components.ErrorMessage("").Render(r.Context(), w)
}

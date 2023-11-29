package controllers

import (
	"net/http"
	"strings"

	"github.com/tomazcx/odonto-dashboard-go/application/utils/tableutils"
	"github.com/tomazcx/odonto-dashboard-go/components"
	"github.com/tomazcx/odonto-dashboard-go/data/services/client"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

const ENTITIES_PER_PAGE = 15

type FiltersState struct {
	Name    string
	OrderBy string
	Page    int
}

var filtersState FiltersState

type ClientsController struct{}

func (c ClientsController) Index(w http.ResponseWriter, r *http.Request) {
	clients, totalCount, err := client.GetClientsService{}.Execute(0, "", true, "")

	if err != nil {
		http.Error(w, "Erro ao buscar os pacientes na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}

	numOfPages := tableutils.CalculatePagination(totalCount, ENTITIES_PER_PAGE)

	views.Clients(views.ClientsViewProps{Clients: clients, TotalCount: totalCount, NumberOfPages: numOfPages}).Render(r.Context(), w)
}

func (c ClientsController) GetClients(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()

	if err != nil {
		http.Error(w, "Erro ao processar o formulário.", http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	orderBy := r.FormValue("orderBy")

	var field string
	var ascendingSort bool

	if len(orderBy) > 0 {
		parts := strings.Split(orderBy, "/")

		if len(parts) != 2 {
			http.Error(w, "Erro ao processar o formulário.", http.StatusInternalServerError)
			return
		}

		field = parts[0]

		if parts[1] == "descending" {
			ascendingSort = false
		} else {
			ascendingSort = true
		}
	}

	clients, totalCount, err := client.GetClientsService{}.Execute(0, field, ascendingSort, name)

	if err != nil {
		http.Error(w, "Erro ao buscar os pacientes na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}

	numOfPages := tableutils.CalculatePagination(totalCount, ENTITIES_PER_PAGE)
	components.ClientsTable(clients, totalCount, numOfPages, 1).Render(r.Context(), w)
}

func (c ClientsController) ChangePage(w http.ResponseWriter, r *http.Request) {

}

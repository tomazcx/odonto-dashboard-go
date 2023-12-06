package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomazcx/odonto-dashboard-go/application/utils/tableutils"
	"github.com/tomazcx/odonto-dashboard-go/components"
	"github.com/tomazcx/odonto-dashboard-go/data/services/client"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

const ENTITIES_PER_PAGE = 15
const PAGE_RANGE_SIZE = 3

type FiltersState struct {
	name          string
	orderBy       string
	totalEntities int
}

var filtersState FiltersState

type ClientsController struct{}

func (c ClientsController) Index(w http.ResponseWriter, r *http.Request) {
	clients, totalCount, err := client.GetClientsService{}.Execute(0, "", true, "")

	if err != nil {
		http.Error(w, "Erro ao buscar os pacientes na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}
	filtersState.totalEntities = totalCount

	numOfPages := tableutils.CalculatePagination(totalCount, ENTITIES_PER_PAGE)
	startPage := 1
	endPage := PAGE_RANGE_SIZE + startPage - 1

	if endPage > numOfPages {
		endPage = numOfPages
	}

	views.Clients(views.ClientsViewProps{Clients: clients, TotalCount: totalCount, StartPage: startPage, EndPage: endPage}).Render(r.Context(), w)
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
	err = tableutils.GetFieldAndAscending(orderBy, &field, &ascendingSort)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro: %v", err), http.StatusUnprocessableEntity)
		return
	}

	clients, totalCount, err := client.GetClientsService{}.Execute(0, field, ascendingSort, name)

	if err != nil {
		http.Error(w, "Erro ao buscar os pacientes na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}

	filtersState.name = name
	filtersState.orderBy = orderBy
	filtersState.totalEntities = totalCount

	numOfPages := tableutils.CalculatePagination(totalCount, ENTITIES_PER_PAGE)
	startPage := 1
	endPage := PAGE_RANGE_SIZE + startPage - 1

	if endPage > numOfPages {
		endPage = numOfPages
	}

	components.ClientsTable(clients, totalCount, startPage, endPage, 1).Render(r.Context(), w)
}

func (c ClientsController) ChangePage(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	page := r.URL.Query().Get("page")
	pageInt, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Erro: Página inválida.", http.StatusUnprocessableEntity)
		return
	}

	name := filtersState.name
	orderBy := filtersState.orderBy
	totalEntities := filtersState.totalEntities

	var field string
	var ascendingSort bool
	err = tableutils.GetFieldAndAscending(orderBy, &field, &ascendingSort)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro: %v", err), http.StatusUnprocessableEntity)
		return
	}

	totalPages := tableutils.CalculatePagination(totalEntities, ENTITIES_PER_PAGE)

	if pageInt > totalPages || pageInt == 0 {
		http.Error(w, "Erro: Página inválida.", http.StatusUnprocessableEntity)
		return
	}

	var skip, selectedPage int

	if pageInt == -1 {
		skip = (totalPages - 1) * ENTITIES_PER_PAGE
		selectedPage = totalPages
	} else {
		skip = (pageInt - 1) * ENTITIES_PER_PAGE
		selectedPage = pageInt
	}

	clients, totalCount, err := client.GetClientsService{}.Execute(skip, field, ascendingSort, name)

	if err != nil {
		http.Error(w, "Erro ao buscar os pacientes na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}

	numOfPages := tableutils.CalculatePagination(totalCount, ENTITIES_PER_PAGE)
	startPage := selectedPage - PAGE_RANGE_SIZE/2

	if startPage < 1 {
		startPage = 1
	}

	endPage := PAGE_RANGE_SIZE + startPage - 1

	if endPage > numOfPages {
		endPage = numOfPages
	}

	components.ClientsTable(clients, totalCount, startPage, endPage, selectedPage).Render(r.Context(), w)
}

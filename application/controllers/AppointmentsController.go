package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tomazcx/odonto-dashboard-go/application/utils/tableutils"
	"github.com/tomazcx/odonto-dashboard-go/components"
	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/services/appointment"
	"github.com/tomazcx/odonto-dashboard-go/views"
)

const APPOINTMENTS_ENTITIES_PER_PAGE = 15
const APPOINTMENTS_PAGE_RANGE_SIZE = 3

type AppointmentsState struct {
	patiendID     string
	totalEntities int
}

var appointmentsState AppointmentsState

type AppointmentsController struct{}

func (a AppointmentsController) Index(w http.ResponseWriter, r *http.Request) {

	appointmentsState.patiendID = r.URL.Query().Get("patientId")

	if len(appointmentsState.patiendID) == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	skip := 0
	getAppointmentsService := appointment.GetAppointmentsService{}
	totalCount, appointments, err := getAppointmentsService.Execute(appointmentsState.patiendID, skip)

	if err != nil {
		http.Error(w, fmt.Sprintf("Não foi possível buscar as consultas do paciente: %v", err), http.StatusBadRequest)
		return
	}

	numOfPages := tableutils.CalculatePagination(totalCount, APPOINTMENTS_ENTITIES_PER_PAGE)
	startPage := 1
	endPage := APPOINTMENTS_PAGE_RANGE_SIZE + startPage - 1

	if endPage > numOfPages {
		endPage = numOfPages
	}

	appointmentsState.totalEntities = totalCount
	views.Appointments(views.AppointmentsViewProps{Appointments: appointments, ClientID: appointmentsState.patiendID, TotalCount: totalCount, StartPage: startPage, EndPage: endPage}).Render(r.Context(), w)
}

func (a AppointmentsController) ChangePage(w http.ResponseWriter, r *http.Request) {
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

	totalEntities := appointmentsState.totalEntities
	patiendID := appointmentsState.patiendID

	totalPages := tableutils.CalculatePagination(totalEntities, APPOINTMENTS_ENTITIES_PER_PAGE)

	if pageInt > totalPages || pageInt == 0 {
		http.Error(w, "Erro: Página inválida.", http.StatusUnprocessableEntity)
		return
	}

	var skip, selectedPage int

	if pageInt == -1 {
		skip = (totalPages - 1) * APPOINTMENTS_ENTITIES_PER_PAGE
		selectedPage = totalPages
	} else {
		skip = (pageInt - 1) * APPOINTMENTS_ENTITIES_PER_PAGE
		selectedPage = pageInt
	}

	getAppointmentsService := appointment.GetAppointmentsService{}
	totalCount, appointments, err := getAppointmentsService.Execute(patiendID, skip)

	if err != nil {
		http.Error(w, "Erro ao buscar as consultas na base de dados. Contate o administrador do sistema.", http.StatusInternalServerError)
		return
	}

	numOfPages := tableutils.CalculatePagination(totalCount, APPOINTMENTS_ENTITIES_PER_PAGE)
	startPage := selectedPage - APPOINTMENTS_PAGE_RANGE_SIZE/2

	if startPage < 1 {
		startPage = 1
	}

	endPage := APPOINTMENTS_PAGE_RANGE_SIZE + startPage - 1

	if endPage > numOfPages {
		endPage = numOfPages
	}

	components.AppointmentsTable(appointments, totalCount, startPage, endPage, selectedPage).Render(r.Context(), w)
}

func (a AppointmentsController) Create(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if len(appointmentsState.patiendID) == 0 {
		http.Error(w, "Paciente não encontrado", http.StatusNotFound)
		return
	}

	err := r.ParseForm()

	if err != nil {
		components.ErrorMessage("Erro ao processar o formulário. Contate o admnistrador do sistema.").Render(r.Context(), w)
		return
	}

	date := r.FormValue("date")
	teeth := r.FormValue("teeth")
	proccedure := r.FormValue("proccedure")

	createAppointmentsService := appointment.CreateAppointmentService{}
	createAppointmentDTO := dto.CreateAppointmentDTO{Date: date, Teeth: teeth, Proccedure: proccedure}
	err = createAppointmentsService.Execute(appointmentsState.patiendID, createAppointmentDTO)

	if err != nil {
		components.ErrorMessage(fmt.Sprintf("Erro ao registar a consulta: %v", err)).Render(r.Context(), w)
		return
	}

	w.Header().Set("HX-Trigger", "appointmentCreated, reloadTable")
	components.ErrorMessage("").Render(r.Context(), w)
}

func (a AppointmentsController) Delete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	appointmentId := r.URL.Query().Get("id")

	if len(appointmentId) == 0 {
		http.Error(w, "Consulta não encontrada", http.StatusNotFound)
		return
	}

	deleteAppointmentService := appointment.DeleteAppointmentService{}
	err := deleteAppointmentService.Execute(appointmentId)

	if err != nil {
		http.Error(w, fmt.Sprintf("Erro ao deletar a consulta: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Set("HX-Trigger", "reloadTable")
}

package appointment

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type GetAppointmentsService struct{}

func (s GetAppointmentsService) Execute(clientId string, skip int) (int, []entities.Appointment, error) {
	appointmentRepository := repositories.NewAppointmentRepository()
	clientRepository := repositories.NewClientRepository()

	if clientExists, _ := clientRepository.Exists(clientId); !clientExists {
		return 0, []entities.Appointment{}, errors.New("Cliente não encontrado.")
	}

	totalEntities, appointments, err := appointmentRepository.FindByClientId(clientId, skip)

	if err != nil {
		return 0, []entities.Appointment{}, errors.New("Erro ao buscar as consultas na base de dados.")
	}

	return totalEntities, appointments, nil
}

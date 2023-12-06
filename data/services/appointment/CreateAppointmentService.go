package appointment

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type CreateAppointmentService struct{}

func (a CreateAppointmentService) Execute(clientId string, dto dto.CreateAppointmentDTO) error {
	appointmentRepository := repositories.NewAppointmentRepository()
	clientRepository := repositories.NewClientRepository()

	if clientExists, _ := clientRepository.Exists(clientId); !clientExists {
		return errors.New("Cliente não encontrado.")
	}

	err := appointmentRepository.Create(clientId, dto)

	return err
}

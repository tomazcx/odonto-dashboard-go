package appointment

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type DeleteAppointmentService struct{}

func (a DeleteAppointmentService) Execute(id string) error {

	appointmentRepository := repositories.NewAppointmentRepository()

	if appointmentExists := appointmentRepository.Exists(id); !appointmentExists {
		return errors.New("Consulta não encontrada.")
	}

	err := appointmentRepository.Delete(id)

	return err
}

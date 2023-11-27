package client

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
	"github.com/tomazcx/odonto-dashboard-go/infra/db"
)

type CreateClientService struct{}

func (c CreateClientService) Execute(data dto.CreateClientDTO) error {
	clientRepository := repositories.ClientRepository{Db: db.GetConn()}
	err := clientRepository.Create(data)

	if err != nil {
		return errors.New("Error registering the client. Contact the system administrator.")
	}

	return nil
}

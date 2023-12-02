package client

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type DeleteClientService struct{}

func (c DeleteClientService) Execute(id string) error {
	clientRepository := repositories.NewClientRepository()

	clientExists, _ := clientRepository.Exists(id)

	if !clientExists {
		return errors.New("Cliente não encontrado")
	}

	err := clientRepository.Delete(id)

	return err
}

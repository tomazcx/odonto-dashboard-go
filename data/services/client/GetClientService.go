package client

import (
	"errors"
	"log"

	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type GetClientService struct{}

func (c GetClientService) Execute(id string) (entities.Client, error) {
	clientRepository := repositories.NewClientRepository()

	clientExists, _ := clientRepository.Exists(id)

	if !clientExists {
		return entities.Client{}, errors.New("Cliente não encontrado")
	}

	clientData, err := clientRepository.GetClient(id)

	if err != nil {
		log.Println("Error getting the client data: %v\n", err)
		return entities.Client{}, err
	}

	return clientData, nil
}

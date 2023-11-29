package client

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type GetClientsService struct{}

func (c GetClientsService) Execute(skip int, orderByField string, ascendingSort bool, name string) ([]entities.Client, int, error) {
	clientRepository := repositories.NewClientRepository()

	clients, totalEntities, err := clientRepository.GetMany(skip, orderByField, ascendingSort, name)

	if err != nil {
		return []entities.Client{}, 0, errors.New("Error fetching the clients table")
	}

	return clients, totalEntities, nil
}

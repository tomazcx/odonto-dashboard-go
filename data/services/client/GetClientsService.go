package client

import (
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/entities"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type GetClientService struct{}

func (c GetClientService) Execute(skip int, orderByField string, ascendingSort bool, name string) ([]entities.Client, int, error) {
	clientRepository := repositories.NewClientRepository()

	pageSize := 15
	clients, err := clientRepository.GetMany(pageSize, skip, orderByField, ascendingSort, name)

	if err != nil {
		return []entities.Client{}, 0, errors.New("Error fetching the clients table")
	}

	count, err := clientRepository.GetCount()

	if err != nil {
		return []entities.Client{}, 0, errors.New("Error fetching the clients count")
	}

	return clients, count, nil
}

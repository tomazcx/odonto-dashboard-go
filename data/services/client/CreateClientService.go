package client

import (
	"context"
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type CreateClientService struct{}

func (c CreateClientService) Execute(ctx context.Context, data dto.CreateClientDTO) error {
	clientRepository := repositories.NewClientRepository()
	err := clientRepository.Create(ctx, data)

	if err != nil {
		return errors.New("Error registering the client. Contact the system administrator.")
	}

	return nil
}

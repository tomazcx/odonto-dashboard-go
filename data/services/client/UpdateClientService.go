package client

import (
	"context"
	"errors"

	"github.com/tomazcx/odonto-dashboard-go/data/dto"
	"github.com/tomazcx/odonto-dashboard-go/data/repositories"
)

type UpdateClientService struct{}

func (c UpdateClientService) Execute(ctx context.Context, id string, dto dto.UpdateClientDTO) error {
	clientRepository := repositories.NewClientRepository()

	clientExists, _ := clientRepository.Exists(id)

	if !clientExists {
		return errors.New("Cliente não encontrado.")
	}

	err := clientRepository.Update(ctx, id, dto)

	return err
}

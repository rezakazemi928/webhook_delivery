package application

import (
	"context"

	"github.com/rezakazemi928/webhook_delivery/internal/domain"
	"github.com/rezakazemi928/webhook_delivery/internal/ports"
)

type EndpointService struct {
	Repository ports.EndpointRepository
}

func NewEndpointService(repo *ports.EndpointRepository) *EndpointService {
	return &EndpointService{Repository: *repo}
}

func (s *EndpointService) CreateEndpoint(ctx context.Context, endpoint domain.Endpoints) error {
	if err := s.Repository.Create(ctx, &endpoint); err != nil{
		return err
	}
	return nil
}
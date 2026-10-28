package ports

import (
	"context"

	"github.com/rezakazemi928/webhook_delivery/internal/domain"
)

type EndpointRepository interface {
	Create(ctx context.Context, endpoint *domain.Endpoints) error
}
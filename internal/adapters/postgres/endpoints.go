package postgres

import (
	"context"
	"fmt"

	"github.com/rezakazemi928/webhook_delivery/internal/domain"
	"github.com/rezakazemi928/webhook_delivery/internal/ports"
	"gorm.io/gorm"
)

type EndpointRepository struct {
	db *gorm.DB
}

func NewEndpointRepository(db *gorm.DB) *EndpointRepository {
	return &EndpointRepository{db: db}
}

var _ ports.EndpointRepository = (*EndpointRepository)(nil)

func (r *EndpointRepository) Create(ctx context.Context, endpoint *domain.Endpoints) error {
	if err := r.db.WithContext(ctx).Create(endpoint).Error; err != nil {
		return fmt.Errorf("create endpoint: %w", err)
	}

	return nil
}
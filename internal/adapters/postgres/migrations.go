package postgres

import (
	"github.com/rezakazemi928/webhook_delivery/internal/domain"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.Endpoints{})
}

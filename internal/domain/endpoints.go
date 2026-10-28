package domain

import "time"

type Endpoints struct {
	ID          int64     `gorm:"primaryKey;autoIncrement;column:id"`
	URL         string    `gorm:"type:text;not null;column:url"`
	Events      []string  `gorm:"type:jsonb;serializer:json;column:events"`
	Description *string   `gorm:"type:text;column:description"`
	CreatedAt   time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
